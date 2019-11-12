package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// BackendUserQueryParam 用于查询的类
type MealTypeQueryParam struct {
	BaseQueryParam
	NameLike string
}

type MealType struct {
	Id     int64
	Name   string
	Time   int64
}

// TableName 设置MealUser表名
func (a *MealType) TableName() string {
	return MealTypeTBName()
}


// CoursePageList 获取分页数据
func MealTypePageList(params *MealTypeQueryParam) ([]*MealType, int64) {
	query := orm.NewOrm().QueryTable(MealTypeTBName())
	data := make([]*MealType, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	if params.NameLike != "" {
		query = query.Filter("name__istartswith", params.NameLike)
	}

	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// MealUserOne 根据id获取单条
func MealTypeOne(id int64) (*MealType, error) {
	o := orm.NewOrm()
	m := MealType{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func AddMealType(name string) (int64,error) {
	o := orm.NewOrm()
	m := &MealType{Name:name,Time:time.Now().Unix()}
	id,err := o.Insert(m)
	return id,err
}
// MealBatchDelete 批量删除
func MealTypeBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(MealTypeTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
