HttpUtils
=======

golang 的一个简单的“HTTP 请求”包 `GET` `POST` `DELETE` `PUT`



### 安装
go get github.com/druidcaesa/gotool

go.mod github.com/druidcaesa/gotool


### 我们如何使用HttpUtils？

#### 使用 http.DefaultTransport 创建请求对象
```go
resp, err := gotool.HttpUtils.Get("http://127.0.0.1:8000")
resp, err := gotool.HttpUtils.SetTimeout(5).Get("http://127.0.0.1:8000")
resp, err := gotool.HttpUtils.Debug(true).SetHeaders(map[string]string{}).Get("http://127.0.0.1:8000")

OR

req := gotool.HttpUtils.NewRequest()
req := gotool.HttpUtils.NewRequest().Debug(true).SetTimeout(5)
resp, err := req.Get("http://127.0.0.1:8000")
resp, err := req.Get("http://127.0.0.1:8000",nil)
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest")
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest","address=beijing")

```

#### 设置请求头
```go
req.SetHeaders(map[string]string{
    "Content-Type": "application/x-www-form-urlencoded",
    "Connection": "keep-alive",
})

req.SetHeaders(map[string]string{
    "Source":"api",
})
```

#### 设置Cookies
```go
req.SetCookies(map[string]string{
    "name":"json",
    "token":"",
})

OR

gotool.HttpUtils.SetCookies(map[string]string{
    "age":"19",
}).Post()
```


#### 设置超时时间
```go
req.SetTimeout(5)  //default 30s
```

#### 面向对象的操作模式
```go
req := gotool.HttpUtils.NewRequest().
	Debug(true).
	SetHeaders(map[string]string{
	    "Content-Type": "application/x-www-form-urlencoded",
	}).SetTimeout(5)
resp,err := req.Get("http://127.0.0.1")

resp,err := gotool.HttpUtils.NewRequest().Get("http://127.0.0.1")
```

### GET

#### 查询参数
```go
resp, err := req.Get("http://127.0.0.1:8000")
resp, err := req.Get("http://127.0.0.1:8000",nil)
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest")
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest","address=beijing")

OR

resp, err := gotool.HttpUtils.Get("http://127.0.0.1:8000")
resp, err := gotool.HttpUtils.Debug(true).SetHeaders(map[string]string{}).Get("http://127.0.0.1:8000")
```


#### 多参数
```go
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest",map[string]interface{}{
    "name":  "jason",
    "score": 100,
})
defer resp.Close()

body, err := resp.Body()
if err != nil {
    return
}

return string(body)
```


### POST

```go
// Send nil
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000")

// Send integer
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000", 100)

// Send []byte
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000", []byte("bytes data"))

// Send io.Reader
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000", bytes.NewReader(buf []byte))
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000", strings.NewReader("string data"))
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000", bytes.NewBuffer(buf []byte))

// Send string
resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000", "title=github&type=1")

// Send JSON
resp, err := gotool.HttpUtils.JSON().Post("http://127.0.0.1:8000", "{\"id\":10,\"title\":\"HttpRequest\"}")

// Send map[string]interface{}{}
resp, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "HttpRequest",
})
defer resp.Close()

body, err := resp.Body()
if err != nil {
    return
}
return string(body)

resp, err := gotool.HttpUtils.Post("http://127.0.0.1:8000")
resp, err := gotool.HttpUtils.JSON().Post("http://127.0.0.1:8000",map[string]interface{}{"title":"github"})
resp, err := gotool.HttpUtils.Debug(true).SetHeaders(map[string]string{}).JSON().Post("http://127.0.0.1:8000","{\"title\":\"github\"}")
```

### 代理模式
```go
proxy, err := url.Parse("http://proxyip:proxyport")
if err != nil {
	log.Println(err)
}

resp, err := gotool.HttpUtils.Proxy(http.ProxyURL(proxy)).Get("http://127.0.0.1:8000/ip")
defer resp.Close()

if err != nil {
	log.Println("Request error：%v", err.Error())
}

body, err := resp.Body()
if err != nil {
	log.Println("Get body error：%v", err.Error())
}
log.Println(string(body))
```

### 上传文件
Params: url, filename, fileinput

```go
resp, err := req.Upload("http://127.0.0.1:8000/upload", "/root/demo.txt","uploadFile")
body, err := resp.Body()
defer resp.Close()
if err != nil {
    return
}
return string(body)
```


### 提示模式
#### 默认 false

```go
req.Debug(true)
```

#### 以标准输出打印：
```go
[HttpRequest]
-------------------------------------------------------------------
Request: GET http://127.0.0.1:8000?name=iceview&age=19&score=100
Headers: map[Content-Type:application/x-www-form-urlencoded]
Cookies: map[]
Timeout: 30s
ReqBody: map[age:19 score:100]
-------------------------------------------------------------------
```


## Json
Post JSON 请求

#### 设置请求头
```go
 req.SetHeaders(map[string]string{"Content-Type": "application/json"})
```
Or
```go
req.JSON().Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "github",
})

req.JSON().Post("http://127.0.0.1:8000", "{\"title\":\"github\",\"id\":10}")
```

#### Post 请求
```go
resp, err := req.Post("http://127.0.0.1:8000", map[string]interface{}{
    "id":    10,
    "title": "HttpRequest",
})
```

#### 打印格式化的 JSON
```go
str, err := resp.Export()
if err != nil {
   return
}
```

#### 解组 JSON
```go
var u User
err := resp.Json(&u)
if err != nil {
   return err
}

var m map[string]interface{}
err := resp.Json(&m)
if err != nil {
   return err
}
```
