package controllers

import (
	"encoding/json"
	"time"

	"meal/enums"
	"meal/models"
	"fmt"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"meal/utils"
	"github.com/astaxie/beego/logs"
)

//DailyMealController 菜单管理
type DailyMealController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *DailyMealController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "DataList", "UpdateSeq")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

//Index 角色管理首页
func (c *DailyMealController) Index() {
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "dailymeal/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "dailymeal/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("DailyMealController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("DailyMealController", "Delete")
}

// 获取所有菜谱
func (c *DailyMealController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params models.DailyMealQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	beego.Info("params",params)
	if params.Ddate == 0 {
		if params.DdateStr != "" {
			nowdate ,err := time.ParseInLocation("2006-01-02",params.DdateStr,time.Local)
			if err != nil {
				beego.Info("date err",err)
			}
			params.Ddate = nowdate.Unix()
		} else {
			//获取当前日期时间戳
			date := utils.GetNow()
			params.NameLike = utils.ToString(date)
			params.Ddate = date
		}

	} else {
		nowdate ,err := time.ParseInLocation("2006-01-02",utils.ToString(params.Ddate),time.Local)
		if err != nil {
			beego.Info("date err",err)
		}
		params.Ddate = nowdate.Unix()
	}

	//获取数据列表和总数
	data, total := models.DailyMealPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}

//Edit 添加、编辑菜谱界面
func (c *DailyMealController) Edit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	c.setTpl("dailymeal/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "dailymeal/edit_footerjs.html"
	c.LayoutSections["headcssjs"] = "dailymeal/edit_headcssjs.html"
}

//Save 添加、编辑页面 保存
func (c *DailyMealController) Save() {
	var err error
	m := models.DailyMeal{}
	//获取form里的值
	ddate := c.GetString("ddate")
	dtype ,_:= c.GetInt32("dtype")
	strs := c.GetString("ids")

	if ddate == "" {
		c.jsonResult(enums.JRCodeFailed, "添加失败，可能原因：请选择日期", ddate)
	}
	if dtype == -1 {
		c.jsonResult(enums.JRCodeFailed, "添加失败，可能原因：未选择类型", dtype)
	}
	if strs == "" {
		c.jsonResult(enums.JRCodeFailed, "添加失败，可能原因：请选择菜谱", strs)
	}
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	beego.Info("DailyMeal====",ddate,dtype,ids)
	date ,err:= time.ParseInLocation("2006-01-02",ddate,time.Local)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "添加失败，可能原因："+err.Error(), date)
	}
	beego.Info("date:",date,date.Unix())
	if m.Id == 0 {
		//m.Creator = &c.curUser

		if exist,succ, err := models.DailyMealBatchAdd(ids,date.Unix(),dtype); err == nil {
			c.jsonResult(enums.JRCodeSucc, "添加成功"+utils.ToString(succ)+"个"+"过滤"+utils.ToString(exist)+"个", strs)
		} else {
			c.jsonResult(enums.JRCodeFailed, "添加失败，可能原因："+err.Error(), m.Id)
		}

	} else {
		c.jsonResult(enums.JRCodeFailed, "每日菜谱只能添加删除",m.Id)
	}

}

//Delete 批量删除
func (c *DailyMealController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := models.DailyMealBatchDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

func (c *DailyMealController) UpdateSeq() {
	Id, _ := c.GetInt64("pk", 0)
	oM, err := models.DailyMealOne(Id)
	if err != nil || oM == nil {
		c.jsonResult(enums.JRCodeFailed, "选择的数据无效", 0)
	}
	value, _ := c.GetInt32("value", 0)
	oM.Seq = value
	o := orm.NewOrm()
	if _, err := o.Update(oM); err == nil {
		c.jsonResult(enums.JRCodeSucc, "修改成功", oM.Id)
	} else {
		c.jsonResult(enums.JRCodeFailed, "修改失败", oM.Id)
	}
}

func (c *DailyMealController) OutList() {
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "dailymeal/outlist_headcssjs.html"
	c.LayoutSections["footerjs"] = "dailymeal/outlist_footerjs.html"
}

func (c *DailyMealController) OutGrid() {
	var req models.DailyMealQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	date := utils.GetNow()
	req.Ddate = date
	req.Dtype = -1
	req.IsOut = true
	//从今日菜单中栓选出外卖数据
	list,total := models.DailyMealPageList(&req)
	var typereq models.MealTypeQueryParam

	typelist,_ := models.MealTypePageList(&typereq)
	rows := make([]*models.DailyMeal,0)
	mlist := make(map[int64]string,0)

	for _,v:=range typelist {
		for _,vv := range list {
			if v.Id == vv.Meal.MealType.Id &&  vv.Meal.IsOut == 0 {
				if _,ok := mlist[vv.Meal.Id];!ok {
					mlist[vv.Meal.Id] = vv.Meal.MealName
					rows = append(rows,vv)
				}
			}
		}
	}
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = rows
	c.Data["json"] = result
	c.ServeJSON()
}


func (c *DailyMealController) UpHot() {
	mm := enums.ReqHot{}
	//获取form里的值
	json.Unmarshal(c.Ctx.Input.RequestBody, &mm)
	logs.Info("uphot",mm)
	oM, err := models.DailyMealOne(mm.Id)
	if err != nil || oM == nil {
		c.jsonResult(enums.JRCodeFailed, "选择的菜单不存在", 0)
	}
	oM.IsHot = 2
	if mm.Utype {
		oM.IsHot = 1
	}

	o := orm.NewOrm()
	if _, err := o.Update(oM,"is_hot"); err == nil {
		c.jsonResult(enums.JRCodeSucc, "修改成功", oM.Id)
	} else {
		c.jsonResult(enums.JRCodeFailed, "修改失败", oM.Id)
	}
}
