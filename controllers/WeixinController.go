package controllers

import (
	"github.com/astaxie/beego"
	"meal/enums"
	"meal/models"
	"encoding/json"
	"github.com/medivhzhan/weapp"
	"meal/utils"
)

type WeixinController struct {
	beego.Controller
	//WeixinUser models.MealUser
}

func (c *WeixinController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *WeixinController) jsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

func (c *WeixinController) WeixinLogin() {
	var req enums.ReqLogin
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if req.Code == "" {
		c.jsonResult(enums.JRCodeFailed,"code 为空",nil)
	} else {
		appid := beego.AppConfig.String("appid")
		appsecret := beego.AppConfig.String("appsecret")
		if appid == "" || appsecret == "" {
			c.jsonResult(enums.JRCodeFailed,"appid 或者appsecret 为空",nil)
		}
		res, err := weapp.Login(appid, appsecret, req.Code)
		if err != nil {
			beego.Error("weixin login err",err)
			c.jsonResult(enums.JRCodeFailed,"微信登录异常"+err.Error(),nil)
		}
		if err := res.CommonError.GetResponseError(); err != nil {
			c.jsonResult(enums.JRCodeFailed,"微信接口登录异常"+err.Error(),nil)
		} else {
			beego.Info("weixin login succ",res)
			isadd,id ,err := models.WeixinUserCheck(res.OpenID)
			if err != nil {
				c.jsonResult(enums.JRCodeFailed,"用户校验异常"+err.Error(),nil)
			}
			var et utils.EasyToken
			et.Username = res.OpenID
			et.Expires = int64(3600*24)
			token,err := et.GetToken()
			if err != nil {
				beego.Error("get token err",err.Error())
			}
			beego.Info("token:",token)
			//新添加用户
			m := make(map[string]interface{},0)
			if id > 0 && isadd {
				beego.Info("weixin add new user",res.OpenID," id ",id)
				var user models.MealUser
				user.OpenId = res.OpenID
				user.SessionKey = res.SessionKey
				user.Id = id
				user.AccessToken = token
				c.SetSession(user.OpenId,user)
			} else if id == 0 && !isadd {
				a := c.GetSession(res.OpenID)
				aa := a.(models.MealUser)
				aa.AccessToken = token
				aa.SessionKey = res.SessionKey
				c.SetSession(res.OpenID,aa)
			}
			m["openid"] = res.OpenID
			m["accesstoken"] = token
			c.jsonResult(enums.JRCodeSucc,"用户登录成功",m)
		}

	}

}

func (c *WeixinController) checkWeixinLogin() {

}
