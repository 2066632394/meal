## 基于beego的后台系统
除了wxlogin接口其他接口header Authorization 带上token   自定义openid = openid
### 1、/weixin/wxlogin
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
    "obj":[{"openid":"string" ,"accesstoken":"string"}]
}
```

### 2、/weixin/list  获取当天或者单周菜谱

|参数|类型|是否必填|说明|
|-|-|-|-|
|dateType|int|是|1是当天菜单 3是当周菜单|
request
```json
{
	"dateType":3
}
```
response

```
{
  "code": 0,
  "msg": "OK",
  "obj": {
    "list": [],
    "total": 0
  }
}
```
### 3、/weixin/secday  统计次日用餐
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

### 4、/weixin/outlist  单日外卖接口
|参数|类型|是否必填|说明|
|-|-|-|-|

request
```json
{
}
```

response
```json
{
  "code": 0,
  "msg": "ok",
  "obj": {
    "count": 0,
    "list": []
  }
}

```

### 5、/weixin/addorder  外卖预定接口
|参数|类型|是否必填|说明|
|-|-|-|-|
|ids|[]int64|是|点的菜编号数组|
request
```json
{
  "ids":[1,2,3,4]
}
```

response
```json
{
  "code": 0,
  "msg": "ok",
  "obj": 123456
}

```


### 6、/weixin/orderlist  外卖取餐列表接口
|参数|类型|是否必填|说明|
|-|-|-|-|
|ids|[]int64|是|点的菜编号数组|
request
```json
{
  "statusType":0,//0提交1已取
}
```

response
```json
{
  "code": 0,
  "msg": "ok",
  "obj": {
    "rows": [],
    "total": 0
  }
}

```

### 7、/weixin/advise  意见提交
|参数|类型|是否必填|说明|
|-|-|-|-|
|ids|[]int64|是|点的菜编号数组|
request
```json
{
  "advise":"",
  "score":""
}
```

response
```json
{
  "code": 0,
  "msg": "ok",
  "obj": nil
}

```
