package controllers

import (
	"github.com/astaxie/beego"
	"meal/models"
	"encoding/json"
	"meal/enums"
)

type MealUserController struct {
	BaseController
}


//Prepare 参考beego官方文档说明
func (c *MealUserController) Prepare() {
	//先执行
	c.BaseController.Prepare()
	//如果一个Controller的多数Action都需要权限控制，则将验证放到Prepare
	c.checkAuthor("DataGrid", "DataList", "UpdateSeq", "UploadImage")
	//如果一个Controller的所有Action都需要登录验证，则将验证放到Prepare
	//权限控制里会进行登录验证，因此这里不用再作登录验证
	//c.checkLogin()
}

//Index 角色管理首页
func (c *MealUserController) Index() {
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "mealuser/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "mealuser/index_footerjs.html"
	//页面里按钮权限控制
	//c.Data["canEdit"] = c.checkActionAuthor("MealUserController", "Edit")
	//c.Data["canDelete"] = c.checkActionAuthor("MealUserController", "Delete")
	//beego.Info("MealUserController,",c.Data["canEdit"],c.Data["canEdit"])
}

// 获取所有菜谱
func (c *MealUserController) DataGrid() {
	//直接反序化获取json格式的requestbody里的值
	var params models.MealUserQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.MealUserPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data

	c.Data["json"] = result
	c.ServeJSON()
}
func (c *MealUserController) OrderList() {
	var params models.MealUserOrderQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	beego.Info("params",params)
	data,total := models.MealUserOrderPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.jsonResult(enums.JRCodeSucc,"ok",result)
}

//func (c *MealUserController) AddOrder() {
//	var params models.MealUserOrderQueryParam
//	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
//	if len(params.Ids) == 0 {
//		c.jsonResult(enums.JRCodeFailed,"菜编号为空",nil)
//	}
//	o := orm.NewOrm()
//	mealIds := ""
//	for k,v := range params.Ids {
//		var dailymeal models.DailyMeal
//		dailymeal.Meal.Id = v
//		dailymeal.MealDate = utils.GetNow()
//
//		if err := o.Read(&dailymeal);err != nil {
//			c.jsonResult(enums.JRCodeFailed,"此菜单不在单日菜谱上",params.MealId)
//		}
//		if k == 0 {
//			mealIds = utils.ToString(v) + ","
//		} else if k == len(params.Ids)-1 {
//			mealIds = mealIds + utils.ToString(v)
//		} else {
//			mealIds = mealIds + ","
//		}
//	}
//
//	Loop:
//	code := utils.RandomString(6)
//	var req models.MealUserOrder
//	req.MealCode = code
//	req.MealDate = utils.GetNow()
//	if err := o.Read(&req); err != nil {
//		goto Loop
//	}
//	req.Time = time.Now().Unix()
//	req.MealIds = mealIds
//	req.Status = enums.OutCommit
//	if id,err := o.Insert(&req);err != nil && id ==0{
//		c.jsonResult(enums.JRCodeFailed,"点餐预订异常"+err.Error(),params)
//	} else {
//		c.jsonResult(enums.JRCodeSucc,"点餐预订成功",code)
//	}
//}