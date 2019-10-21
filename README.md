## 基于beego的后台系统

### 1、/weixin/wexlogin
|参数|类型|是否必填|说明|
|-|-|-|-|
|code|string|是|wx.login返回的code|
request
```json
{
    "code":"123123"
}

```
response
```json
{
    "code":0,
    "msg":"ok",
    "data":[{"openid":"string" ,"accesstoken":"string"}]
}
```

### 2、/dailymeal/list  获取当天或者单周菜谱

|参数|类型|是否必填|说明|
|-|-|-|-|
|dateType|int|是|1是当天菜单 3是当周菜单|
request
```json
{
	"dateType":3
}
```

### 3、/mealuser/secday  统计次日用餐
|参数|类型|是否必填|说明|
|-|-|-|-|
|tomorrow|bool|是|确认明日就餐 true|

request
```json
{
    "tomorrow":true
}
```

response
```json
{
  "code": 0,
  "msg": "谢谢参与",
  "obj": null
}
```