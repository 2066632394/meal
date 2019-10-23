package controllers

import (
	"github.com/astaxie/beego"
	"meal/enums"
	"meal/models"
	"encoding/json"
	"github.com/medivhzhan/weapp"
	"meal/utils"
	"time"
)

type WeixinController struct {
	beego.Controller
	//WeixinUser models.MealUser
}

func (c *WeixinController) Prepare() {
	//c.checkWeixinLogin()
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
		appid := beego.AppConfig.String("weixin::appid")
		appsecret := beego.AppConfig.String("weixin::appsecret")
		if appid == "" || appsecret == "" {
			c.jsonResult(enums.JRCodeFailed,"appid 或者appsecret 为空",nil)
		}
		res, err := weapp.Login(appid, appsecret, req.Code)
		if err != nil {
			beego.Error("weixin login err",err)
			c.jsonResult(enums.JRCodeFailed,"微信登录异常"+err.Error(),nil)
		} else {
			beego.Info("weixin login succ",res)
			isadd,id ,err := models.WeixinUserCheck(res.OpenID)
			if err != nil {
				c.jsonResult(enums.JRCodeFailed,"用户校验异常"+err.Error(),nil)
			}
			var et utils.EasyToken
			et.Username = res.OpenID
			et.Expires = time.Now().Add(time.Hour * 480).Unix()
			token,err := et.GetToken()
			if err != nil {
				beego.Error("get token err",err.Error())
			}
			beego.Info("token:",token)
			//新添加用户
			m := make(map[string]interface{},0)
			var user models.MealUser
			user.OpenId = res.OpenID
			user.SessionKey = res.SessionKey
			user.Id = id
			user.AccessToken = token
			if id > 0 && isadd {
				beego.Info("weixin add new user",res.OpenID," id ",id)

				c.SetSession(user.OpenId,user)
			} else if id > 0 && !isadd {
				a := c.GetSession(res.OpenID)
				if a == nil {
					c.SetSession(res.OpenID,user)
				} else {
					aa,ok := a.(models.MealUser)
					if ok {
						aa.AccessToken = token
						aa.SessionKey = res.SessionKey
						c.SetSession(res.OpenID,aa)
					} else {
						c.SetSession(res.OpenID,user)
					}
				}

			}
			m["openid"] = res.OpenID
			m["accesstoken"] = token
			c.jsonResult(enums.JRCodeSucc,"用户登录成功",m)
		}

	}

}