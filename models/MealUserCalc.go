package models

import (
	"github.com/astaxie/beego/orm"
)

// BackendUserQueryParam 用于查询的类
type MealUserCalcQueryParam struct {
	BaseQueryParam
	MealDate  int64
	Tomorrow bool
	UserId int64
}

type MealUserCalc struct {
	Id     int64
	MealDate int64
	MealNums int32
}

// MealUserOne 根据id获取单条
func MealUserCalcOne(id int64) (*MealUserCalc, error) {
	o := orm.NewOrm()
	m := MealUserCalc{MealDate: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func UpdateUserCalc(params *MealUserCalcQueryParam) error {
	o := orm.NewOrm()
	m := &MealUserCalc{MealDate:params.MealDate}
	_,err := o.InsertOrUpdate(m,"MealDate=MealDate+1")
	return err
}