package enums

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


type ReqMaintain struct {
	Name string
	Phone string
	Desc string
	Type string
	Ext string
}