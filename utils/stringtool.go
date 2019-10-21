package utils

import "fmt"
import "crypto/md5"
import "math/rand"
import "time"
import "strconv"
import (
	"strings"
	"runtime/debug"
	"encoding/json"
	"unsafe"
	"regexp"
	"sort"
	"net"
	"github.com/astaxie/beego"
)

//将字符串加密成 md5
func String2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

//RandomString 在数字、大写字母、小写字母范围内生成num位的随机字符串
func RandomString(length int) string {
	// 48 ~ 57 数字
	// 65 ~ 90 A ~ Z
	// 97 ~ 122 a ~ z
	// 一共62个字符，在0~61进行随机，小于10时，在数字范围随机，
	// 小于36在大写范围内随机，其他在小写范围随机
	rand.Seed(time.Now().UnixNano())
	result := make([]string, 0, length)
	for i := 0; i < length; i++ {
		t := rand.Intn(62)
		if t < 10 {
			result = append(result, strconv.Itoa(rand.Intn(10)))
		} else if t < 36 {
			result = append(result, string(rand.Intn(26)+65))
		} else {
			result = append(result, string(rand.Intn(26)+97))
		}
	}
	return strings.Join(result, "")
}



var loc = local()

func local() *time.Location {
	loc, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		panic(err)
	}
	return loc
}

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const RoomRandomId = "0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// generate random string; fast
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandStringBytesMaskImprSrc(n int, lib string) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(lib) {
			b[i] = lib[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func RandomRoomId() string {
	return RandStringBytesMaskImprSrc(9, RoomRandomId)
}

func RandomUsername() string {
	return "chat" + RandStringBytesMaskImprSrc(10, letterBytes)
}

func NowMillionSecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func NowSecond() int64 {
	return time.Now().Unix()
}

//时间戳 毫秒
func DayStartAndEndNowMillionSecond(t time.Time) (int64, int64) {
	now := t
	year := now.Year()
	month := now.Month()
	day := now.Day()

	start := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 0, 1).Add(-time.Second)
	return start.UnixNano() / 1e6, end.UnixNano() / 1e6
}

// 2006-01-02 15:04
func MillionSecondToDateTime(ts int64) time.Time {
	sec := ts / 1000
	nsec := ts % 1000
	return time.Unix(sec, nsec)
}

func MillionSecondAddDate(tm int64, years, months, days int) int64 {
	sec := tm / 1000
	nsec := tm % 1000
	t := time.Unix(sec, nsec)
	t = t.AddDate(years, months, days)
	return t.UnixNano() / 1e6
}

func MillionSecondAddDuration(tm int64, duration time.Duration) int64 {
	sec := tm / 1000
	nsec := tm % 1000
	t := time.Unix(sec, nsec)
	t = t.Add(duration)
	return t.UnixNano() / 1e6
}

func MillionSecondToTimeString(tm int64) string {
	sec := tm / 1000
	nsec := tm % 1000
	return time.Unix(sec, nsec).Format("2006-01-02 15:04")
}

func MillionSecondToTimeString2(tm int64) string {
	sec := tm / 1000
	nsec := tm % 1000
	return time.Unix(sec, nsec).Format("2006-01-02 15:04:05")
}

func getTimeByMillionSecond(tm int64) time.Time {
	sec := tm / 1000
	nsec := tm % 1000
	return time.Unix(sec, nsec)
}

func TimeStrToMillionSecond(tm string) int64 {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", tm, loc)
	if err != nil {
		panic(err)
	}
	return t.UnixNano() / 1e6
}

// true: first > second + day
func CompareTimeInterval(first, second int64, day int) bool {
	firstTime := getTimeByMillionSecond(first)
	secondTime := getTimeByMillionSecond(second)
	return firstTime.After(secondTime.AddDate(0, 0, day))
}

func GetCurrentDays() ([]int64,error) {
	n := time.Now()
	nt := int(n.Weekday())
	day := n.Unix()
	date ,err := time.ParseInLocation("2006-01-02",time.Unix(day,0).Format("2006-01-02"),time.Local)
	if err != nil {
		return nil, err
	}
	ntime := date.Unix()
	list := make([]int64,7)

	for n := 1; n<= 7; n++ {
		s := n -nt

		list[n-1] = ntime + int64(s) * 86400
	}
	return list,nil

}

func GetNow() int64 {
	date := time.Now().Unix()

	nowdate ,err := time.ParseInLocation("2006-01-02",time.Unix(date,0).Format("2006-01-02"),time.Local)
	if err != nil {
		beego.Info("date err",err)
	}
	return nowdate.Unix()
}

func RandInt(min, max int) int {
	if min >= max {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

/*
func TimeStampIntToTimeStr(ts int64) string {
	tm := time.Unix(ts/1000, ts%1000*1000000)
	return tm.Format(TimeLayoutMillionSecond)
}

func TimeStampToTimeStr(ts string) string {
	val, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return ts
	}
	return TimeStampIntToTimeStr(val)
}

const (
	TimeLayoutMillionSecond = "2006-01-02 15:04:05.000"
	TimeLayoutSecond        = "2006-01-02 15:04:05"
)

func TimeStrToTimeStamp(tm string) string {
	if tm == "" {
		return "0"
	}
	ts, err := time.ParseInLocation(TimeLayoutMillionSecond, tm, loc)
	if err != nil {
		ts, err = time.ParseInLocation(TimeLayoutSecond, tm, loc)
		fmt.Println("TimeStrToTimeStamp: Warn time layout format")
	}
	return strconv.FormatInt(ts.UnixNano()/1000000, 10)
}*/

func RFC3339ToTimeStampMillionSecond(rfc string) int64 {
	const layout = "2006-01-02T15:04:05.000Z"
	if rfc == "" {
		return 0
	}
	ts, err := time.Parse(layout, rfc)
	if err != nil {
		return 0
	}
	return ts.UnixNano() / 1000000
}

func ParseString(format string, args ...interface{}) string {
	if len(args) == 0 {
		return format
	}
	return fmt.Sprintf(format, args...)
}

func ToBool(val interface{}) bool {
	ret := ToInt(val)
	return ret > 0
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func ToInt(val interface{}) int {
	return int(ToInt32(val))
}

func ToInt32(o interface{}) int32 {
	if o == nil {
		return 0
	}
	switch t := o.(type) {
	case int:
		return int32(t)
	case int32:
		return t
	case int64:
		return int32(t)
	case float64:
		return int32(t)
	case string:
		if o == "" {
			return 0
		}
		temp, err := strconv.ParseInt(o.(string), 10, 32)
		if err != nil {
			return 0
		}
		return int32(temp)
	default:
		return 0
	}
}

func ToInt64(val interface{}) int64 {
	if val == nil {
		return 0
	}
	switch val.(type) {
	case int:
		return int64(val.(int))
	case string:
		if val.(string) == "" {
			return 0
		}
		ret, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			debug.PrintStack()
			return 0
		}
		return ret
	case float64:
		return int64(val.(float64))
	case int64:
		return val.(int64)
	case json.Number:
		v := val.(json.Number)
		ret, err := v.Int64()
		if err != nil {
			return 0
		}
		return ret
	default:
		return 0
	}
}

func ToString(val interface{}) string {
	if val == nil {
		return ""
	}
	switch val.(type) {
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case int64:
		return strconv.FormatInt(val.(int64), 10)
	}
	return fmt.Sprintf("%v", val)
}

func ToFloat64(val interface{}) float64 {
	if val == nil {
		return 0
	}
	switch val.(type) {
	case string:
		ret, _ := strconv.ParseFloat(val.(string), 64)
		return ret
	default:
		if v, ok := val.(float64); ok {
			return v
		}
		return 0
	}
}

func StructToString(val interface{}) string {
	if val == nil {
		return ""
	}

	switch val.(type) {
	case interface{}:
		bytes, err := json.Marshal(val)
		if err != nil {
			return ""
		}
		return *(*string)(unsafe.Pointer(&bytes))
	default:
		return ""
	}
}

func StringToJobj(val interface{}) map[string]interface{} {
	var rlt = make(map[string]interface{})
	switch val.(type) {
	case string:
		err := json.Unmarshal([]byte(val.(string)), &rlt)
		if err != nil {
			return nil
		}
		return rlt
	default:
		bytes, err := json.Marshal(val)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &rlt)
		if err != nil {
			panic(err)
		}
		return rlt
	}
}

func VisitorNameSplit(src string) string {
	if len(src) >= 8 {
		return src[0:8]
	} else {
		return src[0:]
	}
}

func CheckPhoneNumber(phone string) bool {
	reg := `^1\d{10}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(phone)
}

func CheckYearly(date string) bool {
	reg := `^\d{4}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(date)
}

func PadLeft(src string, num int, s byte) string {
	if len(src) < num {
		len := num - len(src)
		data := make([]byte, len)
		for i := 0; i < len; i++ {
			data[i] = s
		}
		return fmt.Sprintf(string(data[:]) + src)
	}
	return src
}

func GetMonthStartAndEnd(date string) (int64, int64) {
	const layout = "2006-01"
	if date == "" {
		return 0, 0
	}
	ts, err := time.ParseInLocation("2006-01", date, loc)
	if err != nil {
		return 0, 0
	}
	endTs := ts.AddDate(0, 1, 0).Add(-time.Second)
	return ts.Unix(), endTs.Unix()
}

func GetFromDate(date int64,layout string) string {
	return time.Unix(date,0).Format(layout)
}


//------------------排序--------------//
type lessFunc func(p1, p2 interface{}) bool

// multiSorter implements the Sort interface, sorting the changes within.
type AutoSorter struct {
	changes []interface{}
	less    lessFunc
	desc    bool
}

// Sort sorts the argument slice according to the less functions passed to OrderedBy.
func (ms *AutoSorter) Sort(changes []interface{}) {
	ms.changes = changes
	sort.Sort(ms)
}

// OrderedBy returns a Sorter that sorts using the less functions, in order.
// Call its Sort method to sort the data.
func OrderedBy(less lessFunc, desc bool) *AutoSorter {
	return &AutoSorter{
		less: less,
		desc: desc,
	}
}

// Len is part of sort.Interface.
func (ms *AutoSorter) Len() int {
	return len(ms.changes)
}

// Swap is part of sort.Interface.
func (ms *AutoSorter) Swap(i, j int) {
	ms.changes[i], ms.changes[j] = ms.changes[j], ms.changes[i]
}

// Less is part of sort.Interface. It is implemented by looping along the
// less functions until it finds a comparison that discriminates between
// the two items (one is less than the other). Note that it can call the
// less functions twice per call. We could change the functions to return
// -1, 0, 1 and reduce the number of calls for greater efficiency: an
// exercise for the reader.
func (ms *AutoSorter) Less(i, j int) bool {
	p, q := ms.changes[i], ms.changes[j]
	if ms.desc {
		return ms.less(q, p)
	} else {
		return ms.less(p, q)
	}
}



func GetLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", nil
}
