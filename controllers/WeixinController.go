package controllers

import (
	"github.com/astaxie/beego"
	"meal/enums"
	"meal/models"
	"encoding/json"
	"github.com/medivhzhan/weapp"
	"meal/utils"
	"github.com/astaxie/beego/logs"
	"strings"
)

type WeixinController struct {
	beego.Controller
	//WeixinUser models.MealUser
}

func (c *WeixinController) Prepare() {
	c.checkWeixinLogin()
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

func (c *WeixinController) checkWeixinLogin() (bool,error) {
	token,openid := c.getAuthToken()
	if token == "" || openid == "" {
		logs.Error("getAuthToken err openid or token empty")
		return false,enums.ErrTokenOrOpenidNotExist
	}
	u := c.GetSession(openid).(models.MealUser)
	logs.Info("session",u)
	if u.Id != 0 && u.OpenId != "" {
		return true,nil
	}
	return true,nil
}


func (c *WeixinController) getAuthToken() (string,string) {
	tokenstr := c.Ctx.Input.Header("Authorization")
	return  strings.Split(tokenstr," ")[1],c.Ctx.Input.Header("openid")
}

func (c *WeixinController) MealList() {
	var req models.DailyMealQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	beego.Info("req",req)
	logs.Info("datetype",req.DateType)
	m := make(map[string]interface{})
	//当天
	if req.DateType == enums.MealToday {
		//获取当前日期时间戳
		date := utils.GetNow()
		req.Ddate = date
		list,count := models.DailyMealPageList(&req)
		m["total"] = count
		m["list"] = list
	} else if req.DateType == enums.MealWeek {
		datelist,err := utils.GetCurrentDays()
		if err != nil {
			c.jsonResult(enums.JRCodeFailed,"getdate err"+err.Error(),datelist)
		}
		logs.Info("datelist:",datelist)
		list := make([]*models.DailyMeal,0)
		total := int64(0)
		for _,v:= range datelist {
			req.Ddate = v
			list1,count := models.DailyMealPageList(&req)
			total += count
			list = append(list,list1...)
		}
		m["total"] = total
		m["list"] = list
	}
	logs.Info("m",m)
	c.jsonResult(enums.JRCodeSucc,"OK",m)
}

func (c *WeixinController) Secday() {
	var params models.MealUserCalcQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	beego.Info("params",params)
	if params.Tomorrow  {
		params.MealDate = utils.GetNow()
		exist := models.CheckIsExists(params.MealDate,params.UserId)
		if !exist {
			err := models.UpdateUserCalc(&params)
			//定义返回的数据结构
			if err == nil {
				c.jsonResult(enums.JRCodeSucc,"ok",nil)
			} else {
				c.jsonResult(enums.JRCodeFailed,"更新失败："+err.Error(),nil)
			}
		} else {
			c.jsonResult(enums.JRCodeFailed,"一天只能点一次：",nil)
		}

	} else {
		c.jsonResult(enums.JRCodeSucc,"谢谢参与",nil)
	}

}

func (c *WeixinController) OutList() {
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