package models

import (
	"github.com/astaxie/beego/orm"
)

// TableName 设置Meal表名
func (a *Carouse) TableName() string {
	return MealCarouseTBName()
}

// MealQueryParam 用于搜索的类
type MealCarouseQueryParam struct {
	BaseQueryParam
	NameLike string
}

// Carouse 实体类
type Carouse struct {
	Id        int64
	Img     string
	Type    int32
	Time int64
}

// MealPageList 获取分页数据
func MealCarousePageList(params *MealCarouseQueryParam) ([]*Carouse, int64) {
	query := orm.NewOrm().QueryTable(MealCarouseTBName())
	data := make([]*Carouse, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	case "Seq":
		sortorder = "Seq"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	total, _ := query.Count()
	query.RelatedSel().OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// MealBatchDelete 批量删除
func MealCarouseBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(MealCarouseTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

