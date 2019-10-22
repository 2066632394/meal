package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// TableName 设置MealUser表名
func (a *MealUser) TableName() string {
	return MealUserTBName()
}

// MealUserQueryParam 用于查询的类
type MealUserQueryParam struct {
	BaseQueryParam
	NameLike string
	Phone int64
}

// MealUser 实体类
type MealUser struct {
	Id              int64
	OpenId 	        string
	Name           	string
	Phone 			int64
	NickName        string
	Time            int64
	SessionKey      string
	AccessToken     string
}

// MealUserPageList 获取分页数据
func MealUserPageList(params *MealUserQueryParam) ([]*MealUser, int64) {
	query := orm.NewOrm().QueryTable(MealUserTBName())
	data := make([]*MealUser, 0)
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

// MealUserOne 根据id获取单条
func MealUserOne(id int64) (*MealUser, error) {
	o := orm.NewOrm()
	m := MealUser{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
//是否存在，添加id，错误信息
func WeixinUserCheck(id string) (bool,int64,error) {
	o := orm.NewOrm()
	err := o.Begin()
	query := o.QueryTable(MealUserTBName())
	data := make([]*MealUser, 0)
	query.Filter("open_id",id).Limit(1).All(&data)
	if len(data) == 0 {
		nid ,err := o.Insert(&MealUser{OpenId:id,Time:time.Now().Unix()})
		if err != nil {
			o.Rollback()
			return false,0,err
		} else {
			o.Commit()
			return true,nid,nil
		}
	} else {

		o.Rollback()
		return false,data[0].Id,err
	}
}