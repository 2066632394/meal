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
	"strings"
	"github.com/astaxie/beego/validation"
	"os"
	"github.com/hunterhug/go_image"
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
	o := orm.NewOrm()
	var user models.MealUser
	user.OpenId = openid
	//exists,err := models.WeixinUserCheckonly(openid)
	if err := o.Read(&user,"openid");err != nil && err == orm.ErrNoRows {
		c.jsonResult(enums.JRCodeFailed,"用户未注册"+err.Error(),nil)
	} else if err != nil && err != orm.ErrNoRows {
		c.jsonResult(enums.JRCodeFailed,"用户异常"+err.Error(),nil)
	}
	//校验token
	var et utils.EasyToken
	_,err := et.ValidateToken(token)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed,"用户token校验失败"+err.Error(),nil)
	}
	c.OpenId = openid
	c.UserId = user.Id
	logs.Info("OPENID:",openid)
	logs.Info("userid:",user.Id)
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

	url := "http://"+beego.AppConfig.String("imgaddr")+":"+beego.AppConfig.String("httpport")
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
	logs.Info(" sec params",params)

	params.MealDate = utils.GetNow()+86400
	exist := models.CheckIsExists(params.MealDate,c.UserId)
	if !exist {
		o := orm.NewOrm()
		var history models.MealUserCalcHistory
		history.UserId = c.UserId
		history.MealDate = params.MealDate
		history.Time = time.Now().Unix()
		if params.Tomorrow  {
			//堂食
			err := models.UpdateUserCalc(&params,false)
			//定义返回的数据结构
			if err == nil {
				if id,err := o.Insert(&history);err != nil {
					c.jsonResult(enums.JRCodeFailed,"更新失败："+err.Error(),nil)
				} else {
					if id > 0 {
						c.jsonResult(enums.JRCodeSucc,"ok",nil)
					} else {
						c.jsonResult(enums.JRCodeFailed,"更新失败："+err.Error(),nil)
					}
				}
			} else {
				c.jsonResult(enums.JRCodeFailed,"更新失败："+err.Error(),nil)
			}
		} else {
			//不堂食
			err := models.UpdateUserCalc(&params,true)
			//定义返回的数据结构
			if err == nil {
				if id,err := o.Insert(&history);err != nil {
					c.jsonResult(enums.JRCodeFailed,"更新失败："+err.Error(),nil)
				} else {
					if id >0 {
						c.jsonResult(enums.JRCodeSucc,"ok",nil)
					}else {
						c.jsonResult(enums.JRCodeFailed,"更新失败："+err.Error(),nil)
					}
				}
			} else {
				c.jsonResult(enums.JRCodeFailed,"更新失败："+err.Error(),nil)
			}
		}
	} else {
		c.jsonResult(enums.JRCodeFailed,"一天只能点一次",nil)
	}
}

func (c *WxapiController) OutList() {
	var req models.DailyMealQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	date := utils.GetNow()
	req.Ddate = date
	req.Dtype = -1
	req.IsOut = true
	//从今日菜单中栓选出外卖数据
	list,_ := models.DailyMealPageList(&req)
	url := "http://"+beego.AppConfig.String("imgaddr")+":"+beego.AppConfig.String("httpport")
	var typereq models.MealTypeQueryParam
	todaylist := make([]interface{},0)
	typel := make([]*models.MealType,0)
	typelist,_ := models.MealTypePageList(&typereq)
	mlist := make(map[int64]string,0)
	hot := &models.MealType{Id:-1,Name:"今日推荐"}
    typel = append(typel,hot)
    td := &todaydata{}
    td.Id = -1
    td.Title = "今日推荐"
    tdRows := make([]unit1,0)
	for _,v := range list {
		var un unit1
		if v.Meal.IsOut == 0 && v.IsHot == 1 {
			un.Id = v.Meal.Id
			un.Name = v.Meal.MealName
			un.Url = url + v.Meal.MealImg
			un.Sold = v.Meal.Sold
			un.Price = v.Meal.Price
			tdRows = append(tdRows,un)
		}
	}
	td.List = tdRows
	todaylist = append(todaylist,td)
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
				if _,ok := mlist[vv.Meal.Id];!ok {
					mlist[vv.Meal.Id] = vv.Meal.MealName
					rows = append(rows,un)
				}
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
	logs.Info("order:",params)
	if len(params.Ids) == 0 {
		c.jsonResult(enums.JRCodeFailed,"菜编号为空",nil)
	}
	o := orm.NewOrm()
	user := models.MealUser{OpenId:c.OpenId}
	if err := o.Read(&user,"open_id");err != nil {
		c.jsonResult(enums.JRCodeFailed,"用户数据异常",nil)
	}
	params.UserId = user.Id
	response,result,code,err := models.AddOrder(&params)

	if err != nil{
		c.jsonResult(enums.JRCodeFailed,"点餐预订异常"+err.Error(),params)
	} else if err == nil && result && code != ""{
		c.jsonResult(enums.JRCodeSucc,"点餐预订成功",response)
	}
}

func (c *WxapiController) OrderList() {
	var params models.MealUserOrderQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	params.UserId = c.UserId
	params.Order = "desc"
	logs.Info("params",params)
	data,total := models.MealUserOrderPageList(&params)
	mealmap := models.MealAll()
	m := make([]*models.ResponseOrder,0)
	for _,v := range data {
		var order models.ResponseOrder
		orderdetail := make([]*models.OrderDetail,0)
		order.UserOrder = v
		meallist := strings.Split(v.MealIds,",")
		for _,vv := range meallist {
			var o models.OrderDetail
			orderd := strings.Split(vv,"-")
			if len(orderd) != 2 {
				logs.Warn("mealids ",orderd)
				continue
			}
			o.MealNums = utils.ToInt32(orderd[1])
			o.MealId = utils.ToInt64(orderd[0])
			if _,ok := mealmap[o.MealId];ok {
				o.MealName = mealmap[o.MealId].MealName
				o.MealAmount = utils.ToString(int64(o.MealNums)* utils.ToInt64(mealmap[o.MealId].Price))
			} else {
				o.MealName = "菜谱已删除"
				o.MealAmount = "0"
			}

			orderdetail = append(orderdetail,&o)
		}
		order.OrderDetail = orderdetail
		m = append(m,&order)
	}
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = m
	logs.Info("orderlist",m)
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
	params.Order = "desc"
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
	user.Name = utils.FilterEmoji(req.Name)
	user.NickName = utils.FilterEmoji(req.Name)
	user.Img = req.Img
	o := orm.NewOrm()
	nums,err := o.Update(&user,"name","nick_name","img")
	if err == nil  {
		c.jsonResult(enums.JRCodeSucc,"ok",nums)
	} else {
		c.jsonResult(enums.JRCodeFailed,"更新用户出错"+err.Error(),req)
	}
}

func (c *WxapiController) ImgList() {
	var params models.MealCarouselQueryParam
	list,count := models.MealCarousePageList(&params)
	url := "http://"+beego.AppConfig.String("imgaddr")+":"+beego.AppConfig.String("httpport")
	for _,v:= range list {
		v.Img = url+v.Img
	}
	m := make(map[string]interface{})
	m["list"] = list
	m["count"] = count
	c.jsonResult(enums.JRCodeSucc,"ok",m)
}


func (c *WxapiController) Maintain() {
	var params enums.ReqMaintain
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	logs.Info("params",params)
	valid := validation.Validation{}
	_,err := valid.Valid(&params)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed,"参数异常："+err.Error(),nil)
	}
	var maintain models.Maintain
	maintain.Time = time.Now().Unix()
	maintain.ContractName = params.Name
	maintain.ContractPhone = params.Phone
	maintain.DeviceType = params.Type
	maintain.Content = params.Desc
	maintain.Images = params.Images
	maintain.Ext = params.Ext
	maintain.User = &models.MealUser{Id:c.UserId}

	id,err := models.AddMaintain(&maintain)
	if err == nil && id >0 {
		c.jsonResult(enums.JRCodeSucc,"ok",nil)
	}else {
		c.jsonResult(enums.JRCodeFailed,"添加失败："+err.Error(),nil)
	}
}

func (c *WxapiController) UploadImage() {
	//这里type没有用，只是为了演示传值
	beego.Info("body",c.Input().Get("Img"))

	f, h, err := c.GetFile("fileImg")
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "上传失败", "")
	}
	defer f.Close()
	filepreix := "static/upload/"
	exist := utils.IsExist(filepreix)
	if exist {
		logs.Info("filepreix",filepreix)
	} else {
		logs.Error("file path not exist",filepreix)
		// 创建文件夹
		err := os.Mkdir(filepreix, os.ModePerm)
		if err != nil {
			logs.Error("create file err",err)
		} else {
			logs.Error("mkdir success!\n")
		}
	}
	filePath := filepreix + h.Filename
	c.SaveToFile("fileImg", filePath)
	realfilename, err := go_image.RealImageName(filePath)
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "上传失败"+err.Error(), "")
	}
	sp := strings.Split(realfilename,".")
	if len(sp) < 2 {
		c.jsonResult(enums.JRCodeFailed, "上传失败,图片格式异常", "")
	}
	newfile := filepreix+time.Unix(utils.GetNow(),0).Format("2006-01-02")+utils.RandomString(10)+"."+sp[len(sp)-1]
	// 保存位置在 static/upload, 没有文件夹要先创建
	err = go_image.ScaleF2F(filePath, newfile,600 )
	if err != nil {
		c.jsonResult(enums.JRCodeFailed, "上传失败"+err.Error(), "")
	}
	logs.Info("oldfile",filePath)
	logs.Info("newfile",newfile)
	c.jsonResult(enums.JRCodeSucc, "上传成功", "/"+newfile)

}