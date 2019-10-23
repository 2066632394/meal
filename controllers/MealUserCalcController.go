package controllers

import (
	"meal/models"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

type MealUserCalcController struct {
	BaseController
}

//Prepare 参考beego官方文档说明
func (c *MealUserCalcController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "DataList")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

//Index 角色管理首页
func (c *MealUserCalcController) Index() {
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "mealusercalc/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "mealusercalc/index_footerjs.html"
	//页面里按钮权限控制
}

// 获取所有菜谱
func (c *MealUserCalcController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params models.MealUserCalcQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	logs.Info("params",params)
	data, total := models.MealUserCalcPageList(&params)

	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}






