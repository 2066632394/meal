package routers

import (
	"meal/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/context"
)

func init() {


	var FilterToken = func(ctx *context.Context) {
		logs.Info("current router path is ", ctx.Request.RequestURI)
		if ctx.Request.RequestURI != "/weixin/wxlogin" && ctx.Input.Header("Authorization") == "" {
			logs.Error("without token, unauthorized !!")
			ctx.ResponseWriter.WriteHeader(401)
			ctx.ResponseWriter.Write([]byte("no permission"))
		}
		if ctx.Request.RequestURI != "/weixin/wxlogin" && ctx.Input.Header("Authorization") != "" {
			token := ctx.Input.Header("Authorization")
			logs.Info("curernttoken: ", token)
			openid := ctx.Input.Header("openid")
			logs.Info("openid",openid)
		}
	}

	beego.InsertFilter("/weixin/*",beego.BeforeRouter,FilterToken)

	//weixin

	beego.Router("/weixin/wxlogin", &controllers.WeixinController{}, "Post:WeixinLogin")
	//当天菜单、一周菜单、今日外卖
	beego.Router("/weixin/list", &controllers.WxapiController{}, "Post:MealList")

	beego.Router("/weixin/userinfo", &controllers.WxapiController{}, "Post:UserInfo")
	//次日用餐统计
	beego.Router("/weixin/secday", &controllers.WxapiController{}, "Post:Secday")
	//外卖用餐统计
	beego.Router("/weixin/outlist", &controllers.WxapiController{}, "Post:OutList")
	//外卖预定
	beego.Router("/weixin/addorder", &controllers.WxapiController{}, "Post:AddOrder")
	//外卖取餐列表
	beego.Router("/weixin/orderlist", &controllers.WxapiController{}, "Post:OrderList")
	//提交意见
	beego.Router("/weixin/advise", &controllers.WxapiController{}, "Post:Advise")
	//提交意见
	beego.Router("/weixin/adviselist", &controllers.WxapiController{}, "Post:AdviseList")
	//菜谱管理
	beego.Router("/meal/index", &controllers.MealController{}, "*:Index")
	beego.Router("/meal/datagrid", &controllers.MealController{}, "Get,Post:DataGrid")
	beego.Router("/meal/edit/?:id", &controllers.MealController{}, "Get,Post:Edit")
	beego.Router("/meal/delete", &controllers.MealController{}, "Post:Delete")
	beego.Router("/meal/updateseq", &controllers.MealController{}, "Post:UpdateSeq")
	beego.Router("/meal/uploadimage", &controllers.MealController{}, "Post:UploadImage")
	//每日菜谱
	beego.Router("/dailymeal/index", &controllers.DailyMealController{}, "*:Index")
	beego.Router("/dailymeal/datagrid", &controllers.DailyMealController{}, "Get,Post:DataGrid")
	beego.Router("/dailymeal/edit/?:id", &controllers.DailyMealController{}, "Get,Post:Edit")
	beego.Router("/dailymeal/delete", &controllers.DailyMealController{}, "Post:Delete")
	beego.Router("/dailymeal/updateseq", &controllers.DailyMealController{}, "Post:UpdateSeq")

	//菜单分类
	beego.Router("/mealtype/index", &controllers.MealTypeController{}, "*:Index")
	beego.Router("/mealtype/edit/?:id", &controllers.MealTypeController{}, "Get,Post:Edit")
	beego.Router("/mealtype/delete", &controllers.MealTypeController{}, "Post:Delete")
	beego.Router("/mealtype/datagrid", &controllers.MealTypeController{}, "Get,Post:DataGrid")

	//后台外卖显示
	beego.Router("/mealuserorder/index", &controllers.MealUserOrderController{}, "*:Index")
	beego.Router("/mealuserorder/delete", &controllers.MealUserOrderController{}, "Post:Delete")
	beego.Router("/mealuserorder/updatestatus", &controllers.MealUserOrderController{}, "Post:UpdateStatus")
	beego.Router("/mealuserorder/datagrid", &controllers.MealUserOrderController{}, "Get,Post:DataGrid")
	//后台意见显示
	beego.Router("/mealadvise/index", &controllers.MealAdviseController{}, "*:Index")
	beego.Router("/mealadvise/delete", &controllers.MealAdviseController{}, "Post:Delete")
	beego.Router("/mealadvise/datagrid", &controllers.MealAdviseController{}, "Get,Post:DataGrid")
	//后台次日统计显示
	beego.Router("/mealusercalchistory/index", &controllers.MealUserCalcController{}, "*:Index")
	beego.Router("/mealusercalchistory/datagrid", &controllers.MealUserCalcController{}, "Get,Post:DataGrid")

	//用户
	beego.Router("/mealuser/orderlist", &controllers.MealUserController{}, "Post:OrderList")
	beego.Router("/mealuser/index", &controllers.MealUserController{}, "*:Index")
	//beego.Router("/mealuser/edit/?:id", &controllers.MealUserController{}, "Get,Post:Edit")
	//beego.Router("/mealuser/delete", &controllers.MealUserController{}, "Post:Delete")
	beego.Router("/mealuser/datagrid", &controllers.MealUserController{}, "Get,Post:DataGrid")


	//次日用餐统计
	beego.Router("/mealusercalc/index", &controllers.MealUserCalcController{}, "*:Index")

	////课程路由
	//beego.Router("/course/index", &controllers.CourseController{}, "*:Index")
	//beego.Router("/course/datagrid", &controllers.CourseController{}, "Get,Post:DataGrid")
	//beego.Router("/course/edit/?:id", &controllers.CourseController{}, "Get,Post:Edit")
	//beego.Router("/course/delete", &controllers.CourseController{}, "Post:Delete")
	//beego.Router("/course/updateseq", &controllers.CourseController{}, "Post:UpdateSeq")
	//beego.Router("/course/uploadimage", &controllers.CourseController{}, "Post:UploadImage")

	//用户角色路由
	beego.Router("/role/index", &controllers.RoleController{}, "*:Index")
	beego.Router("/role/datagrid", &controllers.RoleController{}, "Get,Post:DataGrid")
	beego.Router("/role/edit/?:id", &controllers.RoleController{}, "Get,Post:Edit")
	beego.Router("/role/delete", &controllers.RoleController{}, "Post:Delete")
	beego.Router("/role/datalist", &controllers.RoleController{}, "Post:DataList")
	beego.Router("/role/allocate", &controllers.RoleController{}, "Post:Allocate")
	beego.Router("/role/updateseq", &controllers.RoleController{}, "Post:UpdateSeq")

	//资源路由
	beego.Router("/resource/index", &controllers.ResourceController{}, "*:Index")
	beego.Router("/resource/treegrid", &controllers.ResourceController{}, "POST:TreeGrid")
	beego.Router("/resource/edit/?:id", &controllers.ResourceController{}, "Get,Post:Edit")
	beego.Router("/resource/parent", &controllers.ResourceController{}, "Post:ParentTreeGrid")
	beego.Router("/resource/delete", &controllers.ResourceController{}, "Post:Delete")
	//快速修改顺序
	beego.Router("/resource/updateseq", &controllers.ResourceController{}, "Post:UpdateSeq")

	//通用选择面板
	beego.Router("/resource/select", &controllers.ResourceController{}, "Get:Select")
	//用户有权管理的菜单列表（包括区域）
	beego.Router("/resource/usermenutree", &controllers.ResourceController{}, "POST:UserMenuTree")
	beego.Router("/resource/checkurlfor", &controllers.ResourceController{}, "POST:CheckUrlFor")

	//后台用户路由
	beego.Router("/backenduser/index", &controllers.BackendUserController{}, "*:Index")
	beego.Router("/backenduser/datagrid", &controllers.BackendUserController{}, "POST:DataGrid")
	beego.Router("/backenduser/edit/?:id", &controllers.BackendUserController{}, "Get,Post:Edit")
	beego.Router("/backenduser/delete", &controllers.BackendUserController{}, "Post:Delete")
	//后台用户中心
	beego.Router("/usercenter/profile", &controllers.UserCenterController{}, "Get:Profile")
	beego.Router("/usercenter/basicinfosave", &controllers.UserCenterController{}, "Post:BasicInfoSave")
	beego.Router("/usercenter/uploadimage", &controllers.UserCenterController{}, "Post:UploadImage")
	beego.Router("/usercenter/passwordsave", &controllers.UserCenterController{}, "Post:PasswordSave")

	beego.Router("/home/index", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/login", &controllers.HomeController{}, "*:Login")
	beego.Router("/home/dologin", &controllers.HomeController{}, "Post:DoLogin")
	beego.Router("/home/logout", &controllers.HomeController{}, "*:Logout")
	beego.Router("/home/datareset", &controllers.HomeController{}, "Post:DataReset")

	beego.Router("/home/404", &controllers.HomeController{}, "*:Page404")
	beego.Router("/home/error/?:error", &controllers.HomeController{}, "*:Error")

	beego.Router("/", &controllers.HomeController{}, "*:Index")

	//本地测试  绕过token
	beego.Router("/test/outlist", &controllers.MainController{}, "Post:OutList")

}

