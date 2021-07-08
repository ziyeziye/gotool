gotool
=======
gotool是一个小而全的Golang工具集，主要是将日常开发中常用的到方法进行提炼集成，避免重复造轮子，提高工作效率，每一个方法都是作者经过工作经验，和从以往的项目中提炼出来的。

### 如何使用gotool呢？

### 安装

go get github.com/druidcaesa/gotool

go.mod github.com/druidcaesa/gotool

### 引入

```go
import "github.com/druidcaesa/gotool"
```

StrUtils
=======
golang一个string常用工具集，基本涵盖了开发中经常用到的工具，目前正在不端的完善中

#### gotool.StrUtils.ReplacePlaceholder 占位符替换

```go
func TestStringReplacePlaceholder(t *testing.T) {
s := "你是我的{},我是你的{}"
placeholder, err := gotool.StrUtils.ReplacePlaceholder(s, "唯一", "所有")
if err == nil {
fmt.Println(placeholder)
}
}
//out
== = RUN   TestStringReplacePlaceholder
你是我的唯一, 我是你的所有
--- PASS: TestStringReplacePlaceholder (0.00s)
PASS
```

#### gotool.StrUtils.RemoveSuffix 去除文件扩展名获取文件名

```go
func TestRemoveSuffix(t *testing.T) {
fullFilename := "test.txt"
suffix, _ := gotool.StrUtils.RemoveSuffix(fullFilename)
fmt.Println(suffix)
fullFilename = "/root/home/test.txt"
suffix, _ = gotool.StrUtils.RemoveSuffix(fullFilename)
fmt.Println(suffix)
}
//out
== = RUN   TestRemoveSuffix
test
test
--- PASS: TestRemoveSuffix (0.00s)
PASS
```

#### gotool.StrUtils.GetSuffix 获取文件扩展名

```go
func TestGetSuffix(t *testing.T) {
fullFilename := "test.txt"
suffix, _ := gotool.StrUtils.GetSuffix(fullFilename)
fmt.Println(suffix)
fullFilename = "/root/home/test.txt"
suffix, _ = gotool.StrUtils.GetSuffix(fullFilename)
fmt.Println(suffix)
}
//out
== = RUN   TestGetSuffix
.txt
.txt
--- PASS: TestGetSuffix (0.00s)
PASS
```

#### gotool.StrUtils.HasEmpty 判断字符串是否未空，我空返回ture

```go
func TestHasStr(t *testing.T) {
str := ""
empty := gotool.StrUtils.HasEmpty(str)
fmt.Println(empty)
str = "11111"
empty = gotool.StrUtils.HasEmpty(str)
fmt.Println(empty)
}
//out
== = RUN   TestHasStr
true
false
--- PASS: TestHasStr (0.00s)
PASS

```

DateUtil
=======
golang一个时间操作工具集，基本涵盖了开发中经常用到的工具，目前正在不端的完善中

#### gotool.DateUtil.FormatToString 时间格式化成字符串

```go
func TestFormatToString(t *testing.T) {
now := gotool.DateUtil.Now()
toString := gotool.DateUtil.FormatToString(&now, "YYYY-MM-DD hh:mm:ss")
fmt.Println(toString)
toString = gotool.DateUtil.FormatToString(&now, "YYYYMMDD hhmmss")
fmt.Println(toString)
}
//年月日对应YYYY MM DD 时分秒 hhmmss 可进行任意组合 比如 YYYY  hh   YYYY-DD hh:mm 等
//out
== = RUN   TestFormatToString
2021-07-07 16:13:30
20210707 161330
--- PASS: TestFormatToString (0.00s)
PASS
```

#### gotool.DateUtil.IsZero 判断时间是否为空

```go
//时间为空 true 否则 false
func TestDate_IsZero(t *testing.T) {
t2 := time.Time{}
zero := gotool.DateUtil.IsZero(t2)
fmt.Println(zero)
zero = gotool.DateUtil.IsZero(gotool.DateUtil.Now())
fmt.Println(zero)
}
//out
== = RUN   TestDate_IsZero
true
false
--- PASS: TestDate_IsZero (0.00s)
PASS
```

#### gotool.DateUtil.Now 获取当前时间 等同于time.Now(),为了统一化所以将此方法也纳入到工具中

#### gotool.DateUtil.InterpretStringToTimestamp 字符串格式化成时间类型

```go
//参数一 需要格式化的时间字符串 参数二 字符串格式，需要和需格式化字符串格式一致 
//如 2021-6-4 对应YYYY-MM-DD  2021.6.4 对应YYYY.MM.DD
func TestInterpretStringToTimestamp(t *testing.T) {
timestamp, err := gotool.DateUtil.InterpretStringToTimestamp("2021-05-04 15:12:59", "YYYY-MM-DD hh:mm:ss")
if err != nil {
gotool.Logs.ErrorLog().Println(err.Error())
}
fmt.Println(timestamp)
}
//out
== = RUN   TestInterpretStringToTimestamp
1620112379
--- PASS: TestInterpretStringToTimestamp (0.00s)
PASS
```

#### gotool.DateUtil.UnixToTime 时间戳转时间

```go
func TestUnixToTime(t *testing.T) {
unix := gotool.DateUtil.Now().Unix()
fmt.Println("时间戳----------------------->", unix)
toTime := gotool.DateUtil.UnixToTime(unix)
fmt.Println(toTime)
}
//out
== = RUN   TestUnixToTime
时间戳-----------------------> 1625645682
2021-07-07 16:14:42 +0800 CST
--- PASS: TestUnixToTime (0.00s)
PASS
```

#### gotool.DateUtil.GetWeekDay 获取星期几

```go
func TestGetWeekDay(t *testing.T) {
now := gotool.DateUtil.Now()
day := gotool.DateUtil.GetWeekDay(now)
fmt.Println("今天是-----------------周", day)
}
//out
== = RUN   TestGetWeekDay
今天是-----------------周 3
--- PASS: TestGetWeekDay (0.00s)
PASS
```

#### gotool.DateUtil.MinuteAddOrSub,HourAddOrSub,DayAddOrSub 时间计算工具

```go
//时间计算
func TestTimeAddOrSub(t *testing.T) {
now := gotool.DateUtil.Now()
fmt.Println("现在时间是--------------------->", now)
sub := gotool.DateUtil.MinuteAddOrSub(now, 10)
fmt.Println("分钟计算结果-------------------->", sub)
sub = gotool.DateUtil.MinuteAddOrSub(now, -10)
fmt.Println("分钟计算结果-------------------->", sub)
sub = gotool.DateUtil.HourAddOrSub(now, 10)
fmt.Println("小时计算结果-------------------->", sub)
sub = gotool.DateUtil.HourAddOrSub(now, -10)
fmt.Println("小时计算结果-------------------->", sub)
sub = gotool.DateUtil.DayAddOrSub(now, 10)
fmt.Println("天计算结果-------------------->", sub)
sub = gotool.DateUtil.DayAddOrSub(now, -10)
fmt.Println("天计算结果-------------------->", sub)
}
//现在时间是---------------------> 2021-07-07 11:18:17.8295757 +0800 CST m=+0.012278001
//分钟计算结果--------------------> 2021-07-07 11:28:17.8295757 +0800 CST m=+600.012278001
//分钟计算结果--------------------> 2021-07-07 11:08:17.8295757 +0800 CST m=-599.987721999
//小时计算结果--------------------> 2021-07-07 21:18:17.8295757 +0800 CST m=+36000.012278001
//小时计算结果--------------------> 2021-07-07 01:18:17.8295757 +0800 CST m=-35999.987721999
//天计算结果--------------------> 2021-07-17 11:18:17.8295757 +0800 CST m=+864000.012278001
//天计算结果--------------------> 2021-06-27 11:18:17.8295757 +0800 CST m=-863999.987721999
```

ConvertUtils 公历转农历工具
============

#### gotool.ConvertUtils.GregorianToLunarCalendar(公历转农历),GetLunarYearDays(农历转公历),GetLunarYearDays(获取农历这一年农历天数)

```go
func TestConvertTest(t *testing.T) {
calendar := gotool.ConvertUtils.GregorianToLunarCalendar(2020, 2, 1)
fmt.Println(calendar)
gregorian := gotool.ConvertUtils.LunarToGregorian(calendar[0], calendar[1], calendar[2], false)
fmt.Println(gregorian)
days := gotool.ConvertUtils.GetLunarYearDays(2021)
fmt.Println(days)
}
//[2020 1 8]
//[2020 2 1]
//354
```

#### gotool.ConvertUtils.GetSolarMonthDays(2021,7)获取公历某月天数 2021年7月天数

#### gotool.ConvertUtils.IsLeapYear(2021)获取某年是否是瑞年 true是 false不是

#### gotool.ConvertUtils.GetLeapMonth(2021)获取某年闰月月份

BcryptUtils 机密和解密工具
==========

#### gotool.BcryptUtils.Generate 加密处理，多用于密码进行加密后数据库存储使用，不可逆

#### gotool.BcryptUtils.CompareHash 加密后和未加密密码对比，多用于登录验证使用

```go
func TestGenerate(t *testing.T) {
//进行加密
generate := gotool.BcryptUtils.Generate("123456789")
fmt.Println(generate)
//进行加密对比
hash := gotool.BcryptUtils.CompareHash(generate, "123456789")
fmt.Println(hash)
}
//out
== = RUN   TestGenerate
$2a$10$IACJ6zGuNuzaumrvDz58Q.vJUzz4JGqougYKrdCs48rQYIRjAXcU2
true
--- PASS: TestGenerate (0.11s)
PASS
```

#### gotool.BcryptUtils.MD5 md5加密

```go
func TestMd5(t *testing.T) {
md5 := gotool.BcryptUtils.MD5("123456789")
fmt.Println(md5)
}
//out
== = RUN   TestMd5
25f9e794323b453885f5181f1b624d0b
--- PASS: TestMd5 (0.00s)
PASS
```

#### gotool.BcryptUtils.GenRsaKey(获取公钥和私钥),

#### RsaSignWithSha256(进行签名),

#### RsaVerySignWithSha256(验证),

#### RsaEncrypt(公钥加密),

#### RsaDecrypt(私钥解密)

```go
func TestRsa(t *testing.T) {
//rsa 密钥文件产生
fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
prvKey, pubKey := gotool.BcryptUtils.GenRsaKey()
fmt.Println(string(prvKey))
fmt.Println(string(pubKey))

fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
var data = "我是被加密的数据记住我哦-------------------------------"
fmt.Println("对消息进行签名操作...")
signData := gotool.BcryptUtils.RsaSignWithSha256([]byte(data), prvKey)
fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
fmt.Println("\n对签名信息进行验证...")
if gotool.BcryptUtils.RsaVerySignWithSha256([]byte(data), signData, pubKey) {
fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
}

fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
ciphertext := gotool.BcryptUtils.RsaEncrypt([]byte(data), pubKey)
fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
sourceData := gotool.BcryptUtils.RsaDecrypt(ciphertext, prvKey)
fmt.Println("私钥解密后的数据：", string(sourceData))
}
//out
== = RUN   TestRsa
-------------------------------获取RSA公私钥-----------------------------------------
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCgHh1ZYFjlxrIJYjjWGFaLwC8Oov8KqyMtHa+GauF121dperr3
i46JyDoskoSBhbkmqv70LMNrjqVdttdIsC0BtH9ThWLBwKVLH56EqfzqlzClKZEh
WTNaddCSuxoZpN33mwS82DCjZe3d7nAPdEGD5pSBx6TVt5bG1c3NVAmcBQIDAQAB
AoGAWc5KO9T0R3xYYzb6Fer0r9GNEzKMxdkTE7zws/3Cky4BKyIxN6LIwbLSHinX
tCXioTOLaDyrJuqNCbEBsr1NoCIPxnswA5Jm5QDYO5x9aq4u8+SZ9KqLbMrO1JDS
ZV7Cbiblz1xNMfdVIvlVjD5RdEotbYTbHI2CZUoPsjHng8kCQQDHi6TJYJqvej8r
q46ZceuWHDgE81Wu16RrA/kZKi6MJAApQtfO/4HM6W/ImbCjZ2rSYxqnAlIg/GxA
dT6iJABjAkEAzWra06RyhGm3lk9j9Xxc0YPX6VX7qT5Iq6c7ry1JtTSeVOksKANG
elaNnCj8YYUgj7BeBBcMMvLd39hP1h06dwJAINTyHQwfB2ZGxImqocajq4QjF3Vu
EKF8dPsnXiOZmwdFW4Sa+30Av1VdRhU7gfc/FTSnKvlvx+ugaA6iao0f3wJBALa8
sTCH4WwUE8K+m4DeAkBMVn34BKnZg5JYcgrzcdemmJeW2rY5u6/HYbCi8WnboUzS
K8Dds/d7AJBKgTNLyx8CQBAeU0St3Vk6SJ6H71J8YtVxlRGDjS2fE0JfUBrpI/bg
r/QI8yMAMyFkt1YshN+UNWJcvR5SXQnyT/udnWJIdg4 =
-----END RSA PRIVATE KEY-----

-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCgHh1ZYFjlxrIJYjjWGFaLwC8O
ov8KqyMtHa+GauF121dperr3i46JyDoskoSBhbkmqv70LMNrjqVdttdIsC0BtH9T
hWLBwKVLH56EqfzqlzClKZEhWTNaddCSuxoZpN33mwS82DCjZe3d7nAPdEGD5pSB
x6TVt5bG1c3NVAmcBQIDAQAB
-----END PUBLIC KEY-----

-------------------------------进行签名与验证操作-----------------------------------------
对消息进行签名操作...
消息的签名信息：  1fcf20c4fb548c8fc0906e369287feb84c861bf488d822d78a0eada23d1af66ed3a12e9440d7181b1748fd0ad805222cf2ce7ce1f6772b330ef11b717700ba26945dda9d749a5c4d8c108ede103c17bed92801a4c3fbc1ebf38d10bf4bf183713eeb7f429acc3dcc3812366a324737f756720f3f9e75ddda1e024a7698b89163

对签名信息进行验证...
签名信息验证成功，确定是正确私钥签名！！
-------------------------------进行加密解密操作-----------------------------------------
公钥加密后的数据： 637b05798c1cf95cfcc63adf228645c77f8e9a9f34b12b722e6938ded8c9560a0430171a4f940d3fb2d96bc6c470c80a817d81f4a2669b991adbff5b22b335129e514c921083ce3e64c1c876409d9b763d5e8e269797283ef951a364da6a59a1d8454391356cb0cd0808092e9dd7ac371f9247a43760f3c82b7ad26a32a7a807
私钥解密后的数据： 我是被加密的数据记住我哦-------------------------------
--- PASS: TestRsa (0.02s)
PASS
```

ZipUtils
=======
压缩和解压工具，可单文件压缩也可进行目录压缩，或者跨目录压缩

#### gotool.ZipUtils.Compress

- files 文件数组 可以是多目录文件
- dest 压缩文件存放地址

```go
func TestCompress(t *testing.T) {
open, err := os.Open("/Users/fanyanan/Downloads/gotool")
if err != nil {
t.Fatal(err)
}
files := []*os.File{open}
flag, err := gotool.ZipUtils.Compress(files, "/Users/fanyanan/Downloads/test.zip")
if err != nil {
t.Fatal(err)
}
if flag {
fmt.Println("压缩成功")
}
}
//out
== = RUN   TestCompress
压缩成功
--- PASS: TestCompress (0.12s)
PASS
```

#### gotool.ZipUtils.DeCompress

- zipFile 压缩包路径
- dest 要解压的路径

```go
func TestDeCompress(t *testing.T) {
	compress, err := gotool.ZipUtils.DeCompress("/Users/fanyanan/Downloads/test.zip", "/Users/fanyanan/Downloads")
	if err != nil {
		t.Fatal(err)
	}
	if compress {
		fmt.Println("解压成功")
	}
}
//out
=== RUN   TestDeCompress
解压成功
--- PASS: TestDeCompress (0.44s)
PASS
```

[comment]: <> (FileUtils)

[comment]: <> (=======)

[comment]: <> (文件操作工具，让做io操作更简单各个方便)



Logs 日志打印工具
=======

#### gotool.Logs.ErrorLog 异常日志

#### gotool.Logs.InfoLog 加载日志

#### gotool.Logs.DebugLog 调试日志

```go
func TestLogs(t *testing.T) {
gotool.Logs.ErrorLog().Println("Error日志测试")
gotool.Logs.InfoLog().Println("Info日志测试")
gotool.Logs.DebugLog().Println("Debug日志测试")
}
//out
== = RUN   TestLogs
[ERROR] 2021/07/07 15:58:10 logs_test.go:9: Error日志测试
[INFO] 2021/07/07 15:58:10 logs_test.go:10: Info日志测试
[DEBUG] 2021/07/07 15:58:10 logs_test.go:11: Debug日志测试
--- PASS: TestLogs (0.00s)
PASS
```

PageUtils 分页工具
=========

#### gotool.PageUtils.Paginator 彩虹分页

```go
func TestPage(t *testing.T) {
paginator := gotool.PageUtils.Paginator(5, 20, 500)
fmt.Println(paginator)
}
//out
== = RUN   TestPage
map[AfterPage:6 beforePage:4 currPage:5 pages:[3 4 5 6 7] totalPages:25]
--- PASS: TestPage (0.00s)
PASS
//说明 AfterPage 下一页  beforePage前一页 currPage当前页 pages页码 totalPages总页数
```

IdUtils
=======
id生成工具，可生成字符串id和int类型id，根据需要选择自己需要的生成规则

#### gotool.IdUtils.IdUUIDToTime 根据时间生成的UUID规则，入参 true消除“-”false保留“-”

```go
func TestUUID(t *testing.T) {
time, err := gotool.IdUtils.IdUUIDToTime(true)
if err == nil {
fmt.Println("根据时间生成UUID去除--------------------->'-'----->", time)
}
time, err = gotool.IdUtils.IdUUIDToTime(false)
if err == nil {
fmt.Println("根据时间生成不去除--------------------->'-'----->", time)
}
}
//out
== = RUN   TestUUID
根据时间生成UUID去除--------------------->'-'-----> 6fb94fe4dfd511ebbc4418c04d462680
根据时间生成不去除--------------------->'-'-----> 6fb9c783-dfd5-11eb-bc44-18c04d462680
--- PASS: TestUUID (0.00s)
PASS
```

#### gotool.IdUtils.IdUUIDToRan 根据随机数生成的UUID推荐使用本方法，并发不会出现重复现象入，参 true消除“-”false保留“-”

```go
    time, err := gotool.IdUtils.IdUUIDToTime(true)
if err == nil {
fmt.Println("根据时间生成UUID去除--------------------->'-'----->", time)
}
time, err = gotool.IdUtils.IdUUIDToTime(false)
if err == nil {
fmt.Println("根据时间生成不去除--------------------->'-'----->", time)
}
//out
== = RUN   TestUUID
根据随机数生成UUID去除--------------------->'-'-----> cf5bcdc585454cda95447aae186d14e6
根据随机数生成不去除--------------------->'-'-----> 72035dba-d45f-480f-b1fd-508d1e036f71
--- PASS: TestUUID (0.00s)
PASS
```

#### gotool.IdUtils.CreateCaptcha 生成随机数id，int类型，入参int 1-18，超过18后会造成int超过长度

```go
func TestCreateCaptcha(t *testing.T) {
captcha, err := gotool.IdUtils.CreateCaptcha(18)
if err != nil {
fmt.Println(err)
}
fmt.Println("18位------------------------------------------>", captcha)
captcha, err = gotool.IdUtils.CreateCaptcha(10)
if err != nil {
fmt.Println(err)
}
fmt.Println("10位------------------------------------------>", captcha)
}
//out
== = RUN   TestCreateCaptcha
18位------------------------------------------> 492457482855750014
10位------------------------------------------> 2855750014
--- PASS: TestCreateCaptcha (0.00s)
PASS
```

#### gotool.IdUtils.GetIdWork根据时间戳在加以计算获取int64的id 长度16位

```go
func TestGetIdWork(t *testing.T) {
work := gotool.IdUtils.GetIdWork()
fmt.Println("根据时间戳在加以计算获取int64类型id-------->", work)
}
//out
== = RUN   TestGetIdWork
根据时间戳在加以计算获取int64类型id--------> 1625746675366450
--- PASS: TestGetIdWork (0.00s)
PASS
```

HttpUtils
=======

golang 的一个简单的“HTTP 请求”包 `GET` `POST` `DELETE` `PUT`

### 我们如何使用HttpUtils？

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
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest", "address=beijing")

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
req.SetTimeout(5) //default 30s
```

#### 面向对象的操作模式

```go
req := gotool.HttpUtils.NewRequest().
Debug(true).
SetHeaders(map[string]string{
"Content-Type": "application/x-www-form-urlencoded",
}).SetTimeout(5)
resp, err := req.Get("http://127.0.0.1")

resp,err := gotool.HttpUtils.NewRequest().Get("http://127.0.0.1")
```

### GET

#### 查询参数

```go
resp, err := req.Get("http://127.0.0.1:8000")
resp, err := req.Get("http://127.0.0.1:8000",nil)
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest")
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest", "address=beijing")

OR

resp, err := gotool.HttpUtils.Get("http://127.0.0.1:8000")
resp, err := gotool.HttpUtils.Debug(true).SetHeaders(map[string]string{}).Get("http://127.0.0.1:8000")
```

#### 多参数

```go
resp, err := req.Get("http://127.0.0.1:8000?id=10&title=HttpRequest", map[string]interface{}{
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
resp, err := gotool.HttpUtils.JSON().Post("http://127.0.0.1:8000", map[string]interface{}{"title":"github"})
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
resp, err := req.Upload("http://127.0.0.1:8000/upload", "/root/demo.txt", "uploadFile")
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
Request: GET http: //127.0.0.1:8000?name=iceview&age=19&score=100
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
