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
	UserId int64
	OpenId string
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

type returndata struct {
	Title string
	Breakfast []unit
	Luanch []unit
	Dinner []unit
	Takeout []unit
}
type todaydata struct {
	Id int64
	Title string
	List []unit1
}
type unit struct {
	Id int64
	Name string
	Url  string
}

type unit1 struct {
	Id int64
	Name string
	Url  string
	Sold int64
	Price string
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
	c.OpenId = openid
}


func (c *WxapiController) getAuthToken() (string,string) {
	tokenstr := c.Ctx.Input.Header("Authorization")
	return  tokenstr,c.Ctx.Input.Header("openid")
}

func (c *WxapiController) MealList() {
	var req models.DailyMealQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	logs.Info("req",req)
	logs.Info("datetype",req.DateType)
	m := make([]returndata,0)
	//当天

	url := "http://"+beego.AppConfig.String("httpaddr1")+":"+beego.AppConfig.String("httpport")
	if req.DateType == enums.MealToday {
		//获取当前日期时间戳
		date := utils.GetNow()
		req.Ddate = date
		list,_ := models.DailyMealPageList(&req)
		var rows returndata
		for _,v:= range list {
			var un unit
			if v.Type == enums.Breakfast {
				un.Id = v.Id
				un.Name = v.Meal.MealName
				un.Url = url + v.Meal.MealImg
				rows.Breakfast = append(rows.Breakfast, un)
			}
			if v.Type == enums.Lunch {
				un.Id = v.Id
				un.Name = v.Meal.MealName
				un.Url = url + v.Meal.MealImg
				rows.Luanch = append(rows.Luanch, un)
			}
			if v.Type == enums.Dinner {
				un.Id = v.Id
				un.Name = v.Meal.MealName
				un.Url = url + v.Meal.MealImg
				rows.Dinner = append(rows.Dinner, un)
			}
			if v.Type == enums.TakeOut {
				un.Id = v.Id
				un.Name = v.Meal.MealName
				un.Url = url + v.Meal.MealImg
				rows.Takeout = append(rows.Takeout, un)
			}

		}
		rows.Title = "当天"
		m = append(m, rows)
	} else if req.DateType == enums.MealWeek {
		datelist,err := utils.GetCurrentDays()
		if err != nil {
			c.jsonResult(enums.JRCodeFailed,"getdate err"+err.Error(),datelist)
		}
		logs.Info("datelist:",datelist)
		for kk,vv:= range datelist {
			req.Ddate = vv
			list1,_ := models.DailyMealPageList(&req)
			var rows returndata
			for _,v:= range list1 {
				var un unit
				if v.Type == enums.Breakfast {
					un.Id = v.Id
					un.Name = v.Meal.MealName
					un.Url = url + v.Meal.MealImg
					rows.Breakfast = append(rows.Breakfast,un)
				}
				if v.Type == enums.Lunch {
					un.Id = v.Id
					un.Name = v.Meal.MealName
					un.Url = url + v.Meal.MealImg
					rows.Luanch = append(rows.Luanch,un)
				}
				if v.Type == enums.Dinner {
					un.Id = v.Id
					un.Name = v.Meal.MealName
					un.Url = url + v.Meal.MealImg
					rows.Dinner = append(rows.Dinner,un)
				}
				if v.Type == enums.TakeOut {
					un.Id = v.Id
					un.Name = v.Meal.MealName
					un.Url = url + v.Meal.MealImg
					rows.Takeout = append(rows.Takeout,un)
				}
			}
			rows.Title = utils.GetDayName(kk+1)
			m = append(m,rows)
		}

	}

	logs.Info("m",m)
	c.jsonResult(enums.JRCodeSucc,"OK",m)
}

func (c *WxapiController) Secday() {
	var params models.MealUserCalcQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	logs.Info("params",params)

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
	var req models.DailyMealQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	date := utils.GetNow()
	req.Ddate = date
	req.Dtype = -1
	//从今日菜单中栓选出外卖数据
	list,_ := models.DailyMealPageList(&req)
	url := "http://"+beego.AppConfig.String("httpaddr1")+":"+beego.AppConfig.String("httpport")
	var typereq models.MealTypeQueryParam
	todaylist := make([]interface{},0)
	typelist,_ := models.MealTypePageList(&typereq)
	for _,v:=range typelist {
		today := &todaydata{}
		today.Id = v.Id
		today.Title = v.Name
		rows := make([]unit1,0)
		for _,vv := range list {
			var un unit1
			if v.Id == vv.Meal.MealType.Id &&  vv.Meal.IsOut == 0 {
				un.Id = vv.Meal.Id
				un.Name = vv.Meal.MealName
				un.Url = url + vv.Meal.MealImg
				un.Sold = vv.Meal.Sold
				un.Price = vv.Meal.Price
				rows = append(rows,un)
			}
		}
		today.List = rows
		todaylist = append(todaylist,today)
	}
	c.jsonResult(enums.JRCodeSucc,"OK",todaylist)
}

func (c *WxapiController) AddOrder() {
	var params models.MealUserOrderQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	if len(params.Ids) == 0 {
		c.jsonResult(enums.JRCodeFailed,"菜编号为空",nil)
	}
	o := orm.NewOrm()
	user := models.MealUser{OpenId:c.OpenId}
	if err := o.Read(&user,"open_id");err != nil {
		c.jsonResult(enums.JRCodeFailed,"用户数据异常",nil)
	}
	params.UserId = user.Id
	result,code,err := models.AddOrder(&params)

	if err != nil{
		c.jsonResult(enums.JRCodeFailed,"点餐预订异常"+err.Error(),params)
	} else if err == nil && result && code != ""{
		c.jsonResult(enums.JRCodeSucc,"点餐预订成功",code)
	}
}

func (c *WxapiController) OrderList() {
	var params models.MealUserOrderQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	logs.Info("params",params)
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
	logs.Info("params",params)
	var advise models.MealAdvise
	advise.Time = time.Now().Unix()
	advise.Score = params.Score
	advise.Advise = params.Advise
	advise.User = &models.MealUser{Id:c.UserId}
	isadd,id,err := models.MealAdviseAddOne(&advise)
	if err == nil && id >0 && isadd {
		c.jsonResult(enums.JRCodeSucc,"ok",nil)
	}else {
		c.jsonResult(enums.JRCodeFailed,"添加失败："+err.Error(),nil)
	}
}

func (c *WxapiController) AdviseList() {
	var params models.MealAdviseQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	logs.Info("params",params)
	list,count:= models.MealAdvisePageList(&params)
	m := make(map[string]interface{})
	m["list"] = list
	m["count"] = count
	c.jsonResult(enums.JRCodeSucc,"ok",m)
	
}

func (c *WxapiController) UserInfo() {
	var req  enums.ReqUpUserInfo
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if req.Openid == "" {
		c.jsonResult(enums.JRCodeFailed,"openid 为空",req)
	}
	var user models.MealUser
	user.Id = c.UserId
	user.OpenId = req.Openid
	user.Name = req.Name
	user.NickName = req.Name
	user.Img = req.Img
	o := orm.NewOrm()
	if nums,err := o.Update(&user,"name","nick_name","img");err == nil && nums == 1 {
		c.jsonResult(enums.JRCodeSucc,"ok",nil)
	} else {
		c.jsonResult(enums.JRCodeFailed,"更新用户出错"+err.Error(),req)
	}


}