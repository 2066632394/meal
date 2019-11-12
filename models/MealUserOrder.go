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
	QueryType int32
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
	User          *MealUser  `orm:"rel(one);on_delete(do_nothing)"`
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
	if params.StatusType != -1 {
		query = query.Filter("status", params.StatusType)
	}
	if params.MealDate != 0 {
		query = query.Filter("meal_date", params.MealDate)
	}
	if params.UserId != 0 {
		query = query.Filter("user_id", params.UserId)
	}
	if params.NameLike != "" {
		query = query.Filter("meal_code",params.NameLike)
	}
	total, _ := query.Count()
	query.RelatedSel().OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
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
func MealUserOrderBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(MealUserOrderTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}
func AddOrder(params *MealUserOrderQueryParam) (*ResponseOrder,bool,string,error) {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		o.Rollback()
		return nil,false,"",err
	}
	response := ResponseOrder{}
	orderdetail := make([]*OrderDetail,0)
	total := int64(0)
	mealIds := ""
	mealmap := MealAll()
	for k,v := range params.Ids {
		var oo OrderDetail
		var dailymeal DailyMeal
		vs := strings.Split(v,"-")
		if len(vs) != 2 {
			o.Rollback()
			return nil,false,"",errors.New("数据格式异常")
		}
		var nmeal Meal
		nmeal.Id = utils.ToInt64(vs[0])
		dailymeal.Meal = &nmeal
		dailymeal.MealDate = utils.GetNow()

		if err := o.Read(&dailymeal,"MealDate","meal_id");err != nil {
			o.Rollback()
			return nil,false,"",errors.New("此菜单不在今日菜谱上")
		}
		if k == 0 {
			if len(params.Ids) > 1 {
				mealIds = utils.ToString(v) + ","
			}else {
				mealIds = utils.ToString(v)
			}

		} else if k == len(params.Ids)-1 {
			mealIds = mealIds + utils.ToString(v)
		} else {
			mealIds = mealIds + utils.ToString(v)+","
		}
		if err := o.Read(&nmeal);err != nil {
			o.Rollback()
			return nil,false,"",err
		}
		nmeal.Sold += utils.ToInt64(vs[1])
		if _,err := o.Update(&nmeal,"sold");err != nil {
			o.Rollback()
			return nil,false,"",err
		}
		total += utils.ToInt64(vs[1]) * utils.ToInt64(mealmap[nmeal.Id].Price)
		oo.MealId = nmeal.Id
		oo.MealName = mealmap[nmeal.Id].MealName
		oo.MealNums = utils.ToInt32(vs[1])
		oo.MealAmount = utils.ToString(utils.ToInt64(vs[1]) * utils.ToInt64(mealmap[nmeal.Id].Price))
		orderdetail = append(orderdetail,&oo)
	}
	response.OrderDetail = orderdetail

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
	req.Total = utils.ToString(total)
	response.UserOrder = &req
	if id,err := o.Insert(&req);err != nil && id ==0{
		o.Rollback()
		return &response,false,code,err
	} else {
		o.Commit()
		return &response,true,code,nil
	}


}

// MealBatchDelete 批量删除
func MealUserOrderBatchUpdate(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(MealUserOrderTBName())
	num, err := query.Filter("id__in", ids).Update(orm.Params{
		"status": 1,
	})
	return num, err
}