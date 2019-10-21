package models

import (
	"github.com/astaxie/beego/orm"
)

// TableName 设置MealUserOrder表名
func (a *MealUserOrder) TableName() string {
	return MealUserOrderTBName()
}

// MealUserOrderQueryParam 用于查询的类
type MealUserOrderQueryParam struct {
	BaseQueryParam
	NameLike string
	Phone int64
	UserId int64
}

// MealUserOrder 实体类
type MealUserOrder struct {
	Id              int64
	User *MealUser  `orm:rel(pk)`
	Type 			int32
	Meal *Meal      `orm:rel(pk)`
	MealCode        string
	Time            int64
}

// MealUserOrderPageList 获取分页数据
func MealUserOrderPageList(params *MealUserOrderQueryParam) ([]*MealUserOrder, int64) {
	query := orm.NewOrm().QueryTable(MealUserOrderTBName())
	data := make([]*MealUserOrder, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("name__istartswith", params.NameLike)
	if params.Phone >0 {
		query = query.Filter("phone", params.Phone)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// MealUserOrderOne 根据id获取单条
func MealUserOrderOne(id int64) (*MealUserOrder, error) {
	o := orm.NewOrm()
	m := MealUserOrder{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}


