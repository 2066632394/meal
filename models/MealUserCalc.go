package models

import (
	"github.com/astaxie/beego/orm"
)

func (a *MealUserCalc) TableName() string {
	return MealUserCalcOrderTBName()
}
// BackendUserQueryParam 用于查询的类
type MealUserCalcQueryParam struct {
	BaseQueryParam
	MealDate  int64
	Tomorrow bool
	UserId int64
}

type MealUserCalc struct {
	MealDate int64 `orm:"pk"`
	MealNums int32
}


// MealUserOne 根据id获取单条
func MealUserCalcOne(id int64) (*MealUserCalc, error) {
	query := orm.NewOrm().QueryTable(MealUserCalcOrderTBName())
	data := make([]*MealUserCalc, 0)
	query.Filter("meal_date", id).Limit(1).All(&data)
	if len(data) == 0 {
		return nil, orm.ErrNoRows
	}
	return data[0], nil
}

func UpdateUserCalc(params *MealUserCalcQueryParam) error {
	o := orm.NewOrm()
	m := &MealUserCalc{MealDate:params.MealDate,MealNums:1}
	//err := o.Read(m)
	//if err != nil && err == orm.ErrNoRows {
	//	if _,err := o.Insert(m);err != nil {
	//		return nil
	//	} else {
	//		return err
	//	}
	//}
	var r orm.RawSeter
	r = o.Raw("INSERT INTO `rms_user_meal_calc` (`meal_date`, `meal_nums`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `meal_date`=?, `meal_nums`= `meal_nums` + ?",m.MealDate,m.MealNums,m.MealDate,m.MealNums)
	_,err := r.Exec()
	if err != nil {
		return err
	}
	return nil
	//_,err := o.InsertOrUpdate(m,"meal_nums=meal_nums+1")
	//return err
}