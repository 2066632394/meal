package controllers

import (
	"strings"

	"meal/enums"
	"meal/models"
	"meal/utils"
	"github.com/astaxie/beego/logs"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	//判断是否登录
	c.checkLogin()
	m := make(map[string]interface{},0)
	//获取次日用餐人数
	m["tomorrow_order_num"] = 0
	m["user_num"] = 0
	m["today_out_num"] = 0
	tomorrow := utils.GetNow()+86400
	calc,err := models.MealUserCalcOne(tomorrow)
	if err != nil {
		logs.Info("get tomorrow calc order err",err.Error())
	}
	if calc!= nil && calc.MealNums != 0 {
		m["tomorrow_order_num"] = calc.MealNums
	}
	//获取当日外卖单数
	var order models.MealUserOrderQueryParam
	order.MealDate = utils.GetNow()
	_,count := models.MealUserOrderPageList(&order)
	m["today_out_num"] = count
	//获取总用户数
	var user models.MealUserQueryParam
	user.Limit = 1
	user.Order = "desc"
	last,_ := models.MealUserPageList(&user)
	if len(last) > 0 {
		m["user_num"] = last[0].Id
	}
	logs.Info("calc",m)
	c.Data["json"] = m
	c.setTpl()

}
func (c *HomeController) Page404() {
	c.setTpl()
}
func (c *HomeController) Error() {
	c.Data["error"] = c.GetString(":error")
	c.setTpl("home/error.html", "shared/layout_pullbox.html")
}
func (c *HomeController) Login() {

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "home/login_headcssjs.html"
	c.LayoutSections["footerjs"] = "home/login_footerjs.html"
	c.setTpl("home/login.html", "shared/layout_base.html")
}
func (c *HomeController) DoLogin() {
	username := strings.TrimSpace(c.GetString("UserName"))
	userpwd := strings.TrimSpace(c.GetString("UserPwd"))
	if len(username) == 0 || len(userpwd) == 0 {
		c.jsonResult(enums.JRCodeFailed, "用户名和密码不正确", "")
	}
	userpwd = utils.String2md5(userpwd)
	user, err := models.BackendUserOneByUserName(username, userpwd)
	if user != nil && err == nil {
		if user.Status == enums.Disabled {
			c.jsonResult(enums.JRCodeFailed, "用户被禁用，请联系管理员", "")
		}
		//保存用户信息到session
		c.setBackendUser2Session(user.Id)
		//获取用户信息
		c.jsonResult(enums.JRCodeSucc, "登录成功", "")
	} else {
		c.jsonResult(enums.JRCodeFailed, "用户名或者密码错误", "")
	}
}
func (c *HomeController) Logout() {
	user := models.BackendUser{}
	c.SetSession("backenduser", user)
	c.pageLogin()
}
func (c *HomeController) DataReset() {
	if ok, err := models.DataReset(); ok {
		c.jsonResult(enums.JRCodeSucc, "初始化成功", "")
	} else {
		c.jsonResult(enums.JRCodeFailed, "初始化失败,可能原因:"+err.Error(), "")
	}

}
