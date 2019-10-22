package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// TableName 设置DailyMeal表名
func (a *DailyMeal) TableName() string {
	return DailyMealTBName()
}

// DailyMealQueryParam 用于搜索的类
type DailyMealQueryParam struct {
	BaseQueryParam
	NameLike string
	DateType int32
	Ddate    int64
	Dtype    int32
}

// DailyMeal 实体类
type DailyMeal struct {
	Id        int64
	Type      int32
	//MealId    int64
	Meal *Meal `orm:"rel(fk)"`
	MealDate  int64
	Seq  int32
	Time int64

}

// DailyMealPageList 获取分页数据
func DailyMealPageList(params *DailyMealQueryParam) ([]*DailyMeal, int64) {
	query := orm.NewOrm().QueryTable(DailyMealTBName())
	data := make([]*DailyMeal, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	case "MealDate":
		sortorder = "MealDate"
	case "Seq":
		sortorder = "Seq"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	if params.Ddate != 0 {
		query = query.Filter("MealDate", params.Ddate)
	}
	query = query.Filter("Type",params.Dtype)
	total, _ := query.Count()
	query.RelatedSel().OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// DailyMealBatchDelete 批量删除
func DailyMealBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(DailyMealTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

//批量添加
func DailyMealBatchAdd(ids []int,date int64 ,dtype int32) (int64,int64, error) {
	query := orm.NewOrm().QueryTable(DailyMealTBName())
	exists := make([]int,0)
	new := make([]int,0)
	add := int64(0)
	for _,id := range ids {
		num, err := query.Filter("meal_date",date).Filter( "type",dtype).Filter("meal_id",id).Count()
		if err != nil  {
			return 0,0,err

		}
		if num == 1 {
			exists = append(exists,id)
		} else {
			new = append(new,id)
		}
	}
	i,err := query.PrepareInsert()
	if err != nil {
		return 0,0,err
	}
	for _,id := range new {
		var daily DailyMeal
		daily.MealDate = date
		daily.Type = dtype
		daily.Time = time.Now().Unix()
		if daily.Meal == nil {
			daily.Meal = &Meal{}
		}
		daily.Meal.Id = int64(id)
		num,err := i.Insert(&daily)
		if err != nil {
			return 0, 0, err
		}
		if num > 0 {
			add++
		}
	}
	i.Close()
	return int64(len(exists)),add, nil
}

// DailyMealOne 获取单条
func DailyMealOne(id int64) (*DailyMeal, error) {
	o := orm.NewOrm()
	m := DailyMeal{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

//AddDailyMeal 添加菜单
func AddDailyMeal(params *DailyMeal) (bool,int64,error) {
	o := orm.NewOrm()
	err := o.Begin()
	if err != nil {
		return false, 0, err
	}
	return o.ReadOrCreate(&params, "Type","MealId","MealDate")

}