package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// init 初始化
func init() {
	orm.RegisterModel(
		new(Course),
		new(BackendUser),
		new(Resource),
		new(Role),
		new(RoleResourceRel),
		new(RoleBackendUserRel),
		new(Meal),
		new(DailyMeal),
		new(MealUser),
		new(MealUserCalc),
		new(MealUserCalcHistory),
		new(MealType),
		new(MealUserOrder),
		new(MealAdvise))
}

// TableName 下面是统一的表名管理
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}

// BackendUserTBName 获取 BackendUser 对应的表名称
func BackendUserTBName() string {
	return TableName("backend_user")
}

// ResourceTBName 获取 Resource 对应的表名称
func ResourceTBName() string {
	return TableName("resource")
}

// RoleTBName 获取 Role 对应的表名称
func RoleTBName() string {
	return TableName("role")
}

// RoleResourceRelTBName 角色与资源多对多关系表
func RoleResourceRelTBName() string {
	return TableName("role_resource_rel")
}

// RoleBackendUserRelTBName 角色与用户多对多关系表
func RoleBackendUserRelTBName() string {
	return TableName("role_backenduser_rel")
}

func CourseTBName() string {
	return TableName("course")
}

func MealTBName() string {
	return TableName("meal")
}

func MealTypeTBName() string {
	return TableName("meal_type")
}

func DailyMealTBName() string {
	return TableName("day_meal")
}

func MealUserTBName() string {
	return TableName("user")
}

func MealUserAdviseTBName() string {
	return TableName("advise")
}

func MealAdviseTagTBName() string {
	return TableName("advise_tag")
}

func MealUserOrderTBName() string {
	return TableName("user_order")
}

func MealUserCalcOrderTBName() string {
	return TableName("user_meal_calc")
}

func MealUserCalcHistoryTBName() string {
	return TableName("user_calc_history")
}