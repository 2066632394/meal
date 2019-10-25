package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"time"
	"strings"
	"strconv"
	"fmt"
	"meal/models"
	"meal/enums"
	"github.com/astaxie/beego/logs"
	"meal/utils"
)

type MealUserOrderController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *MealUserOrderController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "DataList")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

//Index 角色管理首页
func (c *MealUserOrderController) Index() {
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "mealuserorder/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "mealuserorder/index_footerjs.html"
	//页面里按钮权限控制
	//c.Data["canEdit"] = c.checkActionAuthor("MealUserOrderController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("MealOrderController", "Delete")
	logs.Info("MealUserOrderController,",c.Data["canDelete"])
}

// 获取所有菜谱
func (c *MealUserOrderController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var req models.MealQueryParam
	meallist,_ := models.MealPageList(&req)
	maplist := make(map[int64]string,0)
	for _,v:= range meallist {
		maplist[v.Id] = v.MealName
	}
	var params models.MealUserOrderQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)

	//获取数据列表和总数
	data, total := models.MealUserOrderPageList(&params)
	for _,v := range data {
		ids := strings.Split(v.MealIds,",")
		temp := ""
		for _,vv:= range ids {
			if vv != "" {
				mm := strings.Split(vv,"-")
				if len(mm) != 2 {
					continue
				}
				i := utils.ToInt64(mm[0])
				num := utils.ToString(mm[1])
				temp += maplist[i]+ num +"份  "
			}
		}
		v.MealIds = temp
	}
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	c.Data["json"] = result
	c.ServeJSON()
}

//Edit 添加、编辑菜谱界面
func (c *MealUserOrderController) Edit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt64(":id", 0)
	m := models.MealUserOrder{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	c.setTpl("mealuserorder/edit.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "mealuserorder/edit_headcssjs.html"
	c.LayoutSections["footerjs"] = "mealuserorder/edit_footerjs.html"

	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor("MealUserOrderController.Index")
}

//Save 添加、编辑页面 保存
func (c *MealUserOrderController) Save() {
	var err error
	m := models.MealUserOrder{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "提交表单数据失败，可能原因："+err.Error(), m.Id)
	}
	beego.Info("meal order====",m)
	o := orm.NewOrm()
	if m.Id == 0 {
		//m.Creator = &c.curUser
		m.Time = time.Now().Unix()

		if _, err = o.Insert(&m); err == nil {
			c.jsonResult(enums.JRCodeSucc, "添加成功", m.Id)
		} else {
			c.jsonResult(enums.JRCodeFailed, "添加失败，可能原因："+err.Error(), m.Id)
		}

	} else {
		c.jsonResult(enums.JRCodeFailed, "订单编辑权限不足", m.Id)
	}

}

//Delete 批量删除
func (c *MealUserOrderController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := models.MealBatchDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

func (c *MealUserOrderController) UpdateStatus() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := models.MealUserOrderBatchUpdate(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}