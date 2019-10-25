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
	OpenId string
	Score int32
}

type MealAdvise struct {
	Id     int64
	Advise string
	Reply  string
	TagId    int32
	Score    int32
	User   *MealUser `orm:"rel(one)"`
	Time  int64
}

// CoursePageList 获取分页数据
func MealAdvisePageList(params *MealAdviseQueryParam) ([]*MealAdvise, int64) {
	query := orm.NewOrm().QueryTable(MealUserAdviseTBName())
	data := make([]*MealAdvise, 0)
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
		query = query.Filter("advise__istartswith", params.NameLike)
	}

	if params.UserId != 0 {
		query = query.Filter("user_id", params.UserId)
	}
	total, _ := query.Count()
	query.RelatedSel().OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// MealUserOne
func MealAdviseOne(id int64) (*MealAdvise, error) {
	o := orm.NewOrm()
	advise := MealAdvise{Id:id}
	err := o.Read(&advise)
	if err != nil {
		return nil, err
	}
	return &advise, nil
}

func MealAdviseAddOne(advise *MealAdvise) (bool,int64,error) {
	o := orm.NewOrm()
	if id,err := o.Insert(advise);err != nil {
		return false,id,err
	} else {
		return true, id, nil
	}

}
