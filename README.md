## 基于beego的后台系统

### 1、/weixin/wexlogin
|参数|类型|是否必填|说明|
|-|-|-|-|
|code|string|是|wx.login返回的code|

response
```json
{
    "code":0,
    "msg":"ok",
    "data":[{"openid":"string" ,"accesstoken":"string"}]
}
```

### 2、/dailymeal/list

|参数|类型|是否必填|说明|
|-|-|-|-|
||string|是|wx.login返回的code|