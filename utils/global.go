package utils

func GetDayName(tt int) string {
	//tt := int(time.Unix(t,0).Weekday())
	name := "星期一"
	switch tt {
	case 1:
		name = "星期一"
	case 2:
		name = "星期二"
	case 3:
		name = "星期三"
	case 4:
		name = "星期四"
	case 5:
		name = "星期五"
	}
	return name
}