package controllers

import (
	"github.com/astaxie/beego"
	"meal/enums"
	"meal/models"
	"encoding/json"
	"meal/utils"
	"github.com/astaxie/beego/logs"
	"time"
	"github.com/astaxie/beego/orm"
)

type WxapiController struct {
	beego.Controller
	//WeixinUser models.MealUser
}

func (c *WxapiController) Prepare() {
	c.checkWeixinLogin()
}

func (c *WxapiController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *WxapiController) jsonResult(code enums.JsonResultCode, msg string, obj interface{}) {
	r := &models.JsonResult{code, msg, obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

func (c *WxapiController) checkWeixinLogin() {
	token,openid := c.getAuthToken()
	if token == "" || openid == "" {
		logs.Error("getAuthToken err openid or token empty")
		c.jsonResult(enums.JRCodeFailed,"用户未授权"+enums.ErrNotAuthored.Error(),nil)
	}
	exists,err := models.WeixinUserCheckonly(openid)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed,"数据库异常"+err.Error(),nil)
	}
	if !exists {
		c.jsonResult(enums.JRCodeFailed,"用户未登陆"+err.Error(),nil)
	}
	//校验token
	var et utils.EasyToken
	_,err = et.ValidateToken(token)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed,"用户token校验失败"+err.Error(),nil)
	}
}


func (c *WxapiController) getAuthToken() (string,string) {
	tokenstr := c.Ctx.Input.Header("Authorization")
	return  tokenstr,c.Ctx.Input.Header("openid")
}

func (c *WxapiController) MealList() {
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

func (c *WxapiController) Secday() {
	var params models.MealUserCalcQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	beego.Info("params",params)

	params.MealDate = utils.GetNow()
	exist := models.CheckIsExists(params.MealDate,params.UserId)
	if !exist {
		if params.Tomorrow  {
			//堂食
			err := models.UpdateUserCalc(&params,false)
			//定义返回的数据结构
			if err == nil {
				c.jsonResult(enums.JRCodeSucc,"ok",nil)
			} else {
				c.jsonResult(enums.JRCodeFailed,"更新失败："+err.Error(),nil)
			}
		} else {
			//不堂食
			err := models.UpdateUserCalc(&params,true)
			//定义返回的数据结构
			if err == nil {
				c.jsonResult(enums.JRCodeSucc,"ok",nil)
			} else {
				c.jsonResult(enums.JRCodeFailed,"更新失败："+err.Error(),nil)
			}
		}
	} else {
		c.jsonResult(enums.JRCodeFailed,"一天只能点一次：",nil)
	}
}

func (c *WxapiController) OutList() {
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

func (c *WxapiController) AddOrder() {
	var params models.MealUserOrderQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	if len(params.Ids) == 0 {
		c.jsonResult(enums.JRCodeFailed,"菜编号为空",nil)
	}
	o := orm.NewOrm()
	mealIds := ""
	for k,v := range params.Ids {
		var dailymeal models.DailyMeal
		dailymeal.Meal.Id = v
		dailymeal.MealDate = utils.GetNow()

		if err := o.Read(&dailymeal);err != nil {
			c.jsonResult(enums.JRCodeFailed,"此菜单不在单日菜谱上",params.MealId)
		}
		if k == 0 {
			mealIds = utils.ToString(v) + ","
		} else if k == len(params.Ids)-1 {
			mealIds = mealIds + utils.ToString(v)
		} else {
			mealIds = mealIds + ","
		}
	}

Loop:
	code := utils.RandomString(6)
	var req models.MealUserOrder
	req.MealCode = code
	req.MealDate = utils.GetNow()
	if err := o.Read(&req); err != nil {
		goto Loop
	}
	req.Time = time.Now().Unix()
	req.MealIds = mealIds
	req.Status = enums.OutCommit
	if id,err := o.Insert(&req);err != nil && id ==0{
		c.jsonResult(enums.JRCodeFailed,"点餐预订异常"+err.Error(),params)
	} else {
		c.jsonResult(enums.JRCodeSucc,"点餐预订成功",code)
	}
}

func (c *WxapiController) OrderList() {
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

func (c *WxapiController) Advise() {
	var params models.MealAdviseQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	beego.Info("params",params)
	var advise models.MealAdvise
	advise.Time = time.Now().Unix()
	advise.Score = params.Score
	advise.Advise = params.Advise
	advise.User.Id = params.UserId
	isadd,id,err := models.MealAdviseAddOne(&advise)
	if err == nil && id >0 && isadd {
		c.jsonResult(enums.JRCodeSucc,"ok",nil)
	}else {
		c.jsonResult(enums.JRCodeSucc,"添加失败："+err.Error(),nil)
	}

}