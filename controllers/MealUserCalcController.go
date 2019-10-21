package controllers

import (
	"github.com/astaxie/beego"
	"meal/models"
	"encoding/json"
	"meal/enums"
)

type MealUserCalcController struct {
	WeixinController
}

func (c *MealUserCalcController) Prepare() {

}

func (c *MealUserController) Secday() {
	var params models.MealUserCalcQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	beego.Info("params",params)
	if params.Tomorrow  {
		exist := models.CheckIsExists(params.MealDate,params.UserId)
		if !exist {
			err := models.UpdateUserCalc(&params)
			//定义返回的数据结构
			if err != nil {
				c.jsonResult(enums.JRCodeSucc,"ok",nil)
			} else {
				c.jsonResult(enums.JRCodeFailed,"更新失败："+err.Error(),nil)
			}
		} else {
			c.jsonResult(enums.JRCodeFailed,"一天只能点一次：",nil)
		}

	} else {
		c.jsonResult(enums.JRCodeSucc,"ok",nil)
	}

}
