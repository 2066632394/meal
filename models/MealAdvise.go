package models

import (
	"github.com/astaxie/beego/orm"
)

// TableName 设置DailyMeal表名
func (a *MealAdvise) TableName() string {
	return MealUserAdviseTBName()
}
// BackendUserQueryParam 用于查询的类
type MealAdviseQueryParam struct {
	BaseQueryParam
	NameLike string
	UserId int64
	Advise string
}

type MealAdvise struct {
	Id     int64
	MealDate int64
	MealNums int32
}

// MealUserOne 根据id获取单条
func MealAdviseOne(id int64) (*MealAdvise, error) {
	o := orm.NewOrm()
	m := MealAdvise{MealDate: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
