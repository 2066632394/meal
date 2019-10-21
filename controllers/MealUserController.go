package controllers

import (
	"github.com/astaxie/beego"
	"meal/models"
	"encoding/json"
	"meal/enums"
)

type MealUserController struct {
	WeixinController
}

func (c *MealUserController) Prepare() {

}

func (c *MealUserController) OrderList() {
	var params models.MealUserOrderQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	beego.Info("params",params)
	if params.Phone == 0 {
		c.jsonResult(enums.JRCodeFailed, "用户手机号为空", params)
	}
	data,total := models.MealUserOrderPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
}
