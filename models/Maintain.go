package models

import (
	"github.com/astaxie/beego/orm"
)

// TableName 设置Maintain表名
func (a *Maintain) TableName() string {
	return MaintainTBName()
}

// MaintainQueryParam 用于搜索的类
type MaintainQueryParam struct {
	BaseQueryParam
	NameLike string
}

// Maintain 实体类
type Maintain struct {
	Id        int64
	User *MealUser `orm:"rel(one)"`
	ContractName string
	ContractPhone  string
	DeviceType   string
	Content string
	Ext    string
	Time   int64
}


// MaintainPageList 获取分页数据
func MaintainPageList(params *MaintainQueryParam) ([]*Maintain, int64) {
	query := orm.NewOrm().QueryTable(MaintainTBName())
	data := make([]*Maintain, 0)
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
		query = query.Filter("MaintainName__istartswith", params.NameLike)
	}

	total, _ := query.Count()
	query.RelatedSel().OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// MaintainBatchDelete 批量删除
func MaintainBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(MaintainTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

// MaintainOne 获取单条
func MaintainOne(id int64) (*Maintain, error) {
	o := orm.NewOrm()
	m := Maintain{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

//AddMaintain 添加菜单
func AddMaintain(params *Maintain) (int64,error) {
	o := orm.NewOrm()
	return o.Insert(&params)
}