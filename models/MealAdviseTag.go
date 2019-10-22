package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// BackendUserQueryParam 用于查询的类
type MealAdviseTagQueryParam struct {
	BaseQueryParam
	NameLike string
}

type MealAdviseTag struct {
	Id     int64
	TagName   string
	Time   int64
}

// TableName 设置MealUser表名
func (a *MealAdviseTag) TableName() string {
	return MealAdviseTagTBName()
}


// CoursePageList 获取分页数据
func MealAdviseTagPageList(params *MealAdviseTagQueryParam) ([]*MealAdviseTag, int64) {
	query := orm.NewOrm().QueryTable(MealAdviseTagTBName())
	data := make([]*MealAdviseTag, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("tag_name__istartswith", params.NameLike)
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// MealUserOne 根据id获取单条
func MealAdviseTagOne(id int64) (*MealAdviseTag, error) {
	o := orm.NewOrm()
	m := MealAdviseTag{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func AddMealAdviseTag(name string) (int64,error) {
	o := orm.NewOrm()
	m := &MealAdviseTag{TagName:name,Time:time.Now().Unix()}
	id,err := o.Insert(m)
	return id,err
}