package controllers

import (
	"github.com/astaxie/beego"
	"meal/models"
	"encoding/json"
)

type MealAdviseController struct {
	WeixinController
}

func (c *MealAdviseController) Prepare() {

}

func (c *MealAdviseController) Advise() {
	var params models.MealAdviseQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	beego.Info("params",params)
	//if params.UserId != 0 {
	//	exist := models.CheckIsExists(params.MealDate,params.UserId)
	//	if !exist {
	//		err := models.UpdateUserCalc(&params)
	//		//定义返回的数据结构
	//		if err != nil {
	//			c.jsonResult(enums.JRCodeSucc,"ok",nil)
	//		} else {
	//			c.jsonResult(enums.JRCodeFailed,"更新失败："+err.Error(),nil)
	//		}
	//	} else {
	//		c.jsonResult(enums.JRCodeFailed,"一天只能点一次：",nil)
	//	}
	//
	//} else {
	//	c.jsonResult(enums.JRCodeSucc,"ok",nil)
	//}

}
