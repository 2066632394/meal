package models

import (
	"github.com/astaxie/beego/orm"
)

// BackendUserQueryParam 用于查询的类
type MealUserCalcHistoryQueryParam struct {
	BaseQueryParam
	MealDate  int64
	Phone  int64
	UserId int64
}

type MealUserCalcHistory struct {
	Id     int64
	MealDate int64
	MealNums int32
}

// MealUserOne 根据id获取单条
func MealUserCalcHistoryOne(id int64) (*MealUserCalcHistory, error) {
	o := orm.NewOrm()
	m := MealUserCalcHistory{MealDate: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func CheckIsExists(date,userid int64) bool {
	o := orm.NewOrm()
	m := MealUserCalcHistory{MealDate:date,Id:userid}
	err := o.Read(&m)
	if err == orm.ErrNoRows {
		return false
	}
	return true
}
