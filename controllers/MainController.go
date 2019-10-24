package controllers

import (
	"github.com/astaxie/beego"
	"meal/enums"
	"meal/models"
	"meal/utils"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) jsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}


func (c *MainController) OutList() {
	now := utils.GetNow()
	req := &models.DailyMealQueryParam{
		Dtype:enums.TakeOut,
		Ddate:now,
	}
	list,count := models.DailyMealPageList(req)
	m := make(map[string]interface{},0)
	m["count"] = count
	m["list"] = list
	c.jsonResult(enums.JRCodeSucc,"ok",m)
}
