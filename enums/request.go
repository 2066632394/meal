package enums

import "meal/models"

type ReqLogin struct {
	Code string
}

type ReqMealList struct {
	DateType int32
}

type ReqUpUserInfo struct {
	Openid string
	Img    string
	Name   string
}

type ReqHot struct {
	Id int64
	Utype bool
}

type ResponseOrder struct {
	UserOrder *models.MealUserOrder
	OrderDetail []*OrderDetail
}

type OrderDetail struct {
	MealId int64
	MealName string
	MealNums int32
	MealAmount string
}