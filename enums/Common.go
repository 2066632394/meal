package enums

type JsonResultCode int

const (
	JRCodeSucc JsonResultCode = iota
	JRCodeFailed
	JRCode302 = 302 //跳转至地址
	JRCode401 = 401 //未授权访问
)

const (
	Deleted = iota - 1
	Disabled
	Enabled
)

const (
	Breakfast = iota
	Lunch
	Dinner
	Meeting
	TakeOut
)

const (
	WEIXINUSER = "weixinuser"
)

const (
	MealToday = iota
	MealTodayTakeOut
	MealWeek
)