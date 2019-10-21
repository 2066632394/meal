package routers

import (
	"meal/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
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
			token = strings.Split(token, "")[1]
			logs.Info("curernttoken: ", token)
		}
	}

	beego.InsertFilter("/weixin/*",beego.BeforeRouter,FilterToken)

	//weixin

	beego.Router("/weixin/wxlogin", &controllers.WeixinController{}, "Post:WeixinLogin")
	//当天菜单、一周菜单、今日外卖
	beego.Router("/weixin/list", &controllers.WeixinController{}, "Post:MealList")
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



	//用户
	beego.Router("/mealuser/orderlist", &controllers.MealUserController{}, "Post:OrderList")
	//次日用餐统计
	beego.Router("/mealuser/secday", &controllers.MealUserCalcController{}, "Post:Secday")
	//提交意见
	beego.Router("/mealuser/advise", &controllers.MealAdviseController{}, "Post:Advise")
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

}

