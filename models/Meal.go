package models

import (
	"github.com/astaxie/beego/orm"
)

// TableName 设置Meal表名
func (a *Meal) TableName() string {
	return MealTBName()
}

// MealQueryParam 用于搜索的类
type MealQueryParam struct {
	BaseQueryParam
	NameLike string
}

// Meal 实体类
type Meal struct {
	Id        int64
	MealName      string `orm:"size(32)"`
	MealImg string
	MealDesc  string
	MealType   *MealType `orm:"rel(one);on_delete(do_nothing)"`
	Sold    int64
	Price   string
	RealPrice string
	IsOut    int32
	Score     int32
	ScoreList string
	Seq  int32
	Time int64
}
type ReqMeal struct {
	Id        int64
	MealName      string `orm:"size(32)"`
	MealImg string
	MealDesc  string
	MealType  int64
	Score     int32
	ScoreList string
	IsOut int32
	Price string
	Seq  int32
	Time int64
}


type ResponseOrder struct {
	UserOrder *MealUserOrder
	OrderDetail []*OrderDetail
}

type OrderDetail struct {
	MealId int64
	MealName string
	MealNums int32
	MealAmount string
}

// MealPageList 获取分页数据
func MealPageList(params *MealQueryParam) ([]*Meal, int64) {
	query := orm.NewOrm().QueryTable(MealTBName())
	data := make([]*Meal, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	case "Score":
		sortorder = "Score"
	case "Seq":
		sortorder = "Seq"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	if params.NameLike != "" {
		query = query.Filter("MealName__istartswith", params.NameLike)
	}

	total, _ := query.Count()
	query.RelatedSel().OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// MealBatchDelete 批量删除
func MealBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(MealTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

// MealOne 获取单条
func MealOne(id int64) (*Meal, error) {
	o := orm.NewOrm()
	m := Meal{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func MealExistName(name string) bool {
	query := orm.NewOrm().QueryTable(MealTBName())
	return query.Filter("meal_name",name).Exist()
}

//AddMeal 添加菜单
func AddMeal(params *Meal) (bool,int64,error) {
	o := orm.NewOrm()
	return o.ReadOrCreate(&params, "Name")
}

func MealAll() map[int64]*Meal {
	query := orm.NewOrm().QueryTable(MealTBName())
	data := make([]*Meal, 0)
	mapdata := make(map[int64]*Meal, 0)
	query.All(&data)
	for _,v := range data {
		if _,ok := mapdata[v.Id];!ok {
			mapdata[v.Id] = v
		}
	}
	return mapdata
}