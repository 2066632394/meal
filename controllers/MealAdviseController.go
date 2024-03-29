package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"meal/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"time"
	"strings"
	"strconv"
	"fmt"
	"meal/enums"
)

type MealAdviseController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *MealAdviseController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "DataList")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

//Index 角色管理首页
func (c *MealAdviseController) Index() {
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)

	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "mealadvise/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "mealadvise/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("MealAdviseController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("MealOrderController", "Delete")
	logs.Info("MealAdviseController,",c.Data["canDelete"])
}

// 获取所有菜谱
func (c *MealAdviseController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params models.MealAdviseQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.MealAdvisePageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	c.Data["json"] = result
	c.ServeJSON()
}

//Edit 添加、编辑菜谱界面
func (c *MealAdviseController) Edit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id, _ := c.GetInt64(":id", 0)
	m := models.MealAdvise{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	c.setTpl("mealadvise/edit.html", "shared/layout_page.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "mealadvise/edit_headcssjs.html"
	c.LayoutSections["footerjs"] = "mealadvise/edit_footerjs.html"

	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor("MealAdviseController.Index")
}

//Save 添加、编辑页面 保存
func (c *MealAdviseController) Save() {
	var err error
	m := models.MealAdvise{}
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
func (c *MealAdviseController) Delete() {
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

//回复意见界面
func (c *MealAdviseController) Reply() {
	if c.Ctx.Request.Method == "POST" {
		c.UpdateAdvise()
	}
	Id, _ := c.GetInt64(":id", 0)
	m := models.MealAdvise{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	c.setTpl("mealadvise/reply.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "mealadvise/reply_footerjs.html"
}

//Save 添加、编辑页面 保存
func (c *MealAdviseController) UpdateAdvise() {

	m := models.MealAdvise{}

	//获取form里的值
	if err := c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "提交表单数据失败，可能原因："+err.Error(), m.Id)
	}
	beego.Info("m====",m)
	o := orm.NewOrm()

	//m.Creator = &c.curUser
	m.Time = time.Now().Unix()

	if _, err := o.Update(&m,"reply"); err == nil {
		c.jsonResult(enums.JRCodeSucc, "回复成功", m.Id)
	} else {
		c.jsonResult(enums.JRCodeFailed, "回复失败，可能原因："+err.Error(), m.Id)
	}



}

