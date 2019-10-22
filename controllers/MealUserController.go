package controllers

import (
	"github.com/astaxie/beego"
	"meal/models"
	"encoding/json"
	"meal/enums"
	"meal/utils"
	"github.com/astaxie/beego/orm"
	"time"
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
	data,total := models.MealUserOrderPageList(&params)
	//定义返回的数据结构
	result := make(map[string]interface{})
	result["total"] = total
	result["rows"] = data
	c.jsonResult(enums.JRCodeSucc,"ok",result)
}

func (c *MealUserController) AddOrder() {
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