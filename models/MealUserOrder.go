package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"meal/utils"
	"time"
	"meal/enums"
	"errors"
	"github.com/astaxie/beego/logs"
)

// TableName 设置MealUserOrder表名
func (a *MealUserOrder) TableName() string {
	return MealUserOrderTBName()
}

// MealUserOrderQueryParam 用于查询的类
type MealUserOrderQueryParam struct {
	BaseQueryParam
	StatusType int32
	NameLike string
	Phone int64
	UserId int64
	MealId int64
	MealDate int64
	Ids    []string
}

// MealUserOrder 实体类
type MealUserOrder struct {
	Id              int64
	User          *MealUser  `orm:"rel(one)"`
	Type 			int32
	MealIds         string
	MealDate        int64
	MealCode        string
	Total           string
	Status          int32
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
	query = query.Filter("status", params.StatusType)
	if params.MealDate != 0 {
		query = query.Filter("meal_date", params.MealDate)
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
// MealBatchDelete 批量删除
func MealBatchDeleteOrder(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(MealUserOrderTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
func AddOrder(params *MealUserOrderQueryParam) (bool,string,error) {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		o.Rollback()
		return false,"",err
	}
	mealIds := ""
	for k,v := range params.Ids {
		var dailymeal DailyMeal
		vs := strings.Split(v,"-")
		if len(vs) != 2 {
			o.Rollback()
			return false,"",errors.New("数据格式异常")
		}
		var nmeal Meal
		nmeal.Id = utils.ToInt64(vs[0])
		dailymeal.Meal = &nmeal
		dailymeal.MealDate = utils.GetNow()

		if err := o.Read(&dailymeal,"MealDate","meal_id");err != nil {
			o.Rollback()
			return false,"",errors.New("此菜单不在今日菜谱上")
		}
		if k == 0 {
			if len(params.Ids) != 1 {
				mealIds = utils.ToString(v)
			}else {
				mealIds = utils.ToString(v) + ","
			}

		} else if k == len(params.Ids)-1 {
			mealIds = mealIds + utils.ToString(v)
		} else {
			mealIds = mealIds + ","
		}
		if err := o.Read(&nmeal);err != nil {
			o.Rollback()
			return false,"",err
		}
		nmeal.Sold += utils.ToInt64(vs[1])
		if _,err := o.Update(&nmeal,"sold");err != nil {
			o.Rollback()
			return false,"",err
		}
	}

Loop:
	code := utils.RandomString(6)
	var req MealUserOrder
	req.MealCode = code
	req.MealDate = utils.GetNow()
	if err := o.Read(&req,"MealCode","MealDate"); err != nil && err != orm.ErrNoRows {
		logs.Info(err)
		goto Loop
	}
	req.Time = time.Now().Unix()
	req.MealIds = mealIds
	req.User = &MealUser{Id:params.UserId}
	req.Status = enums.OutCommit
	if id,err := o.Insert(&req);err != nil && id ==0{
		o.Rollback()
		return false,code,err
	} else {
		o.Commit()
		return true,code,nil
	}


}
