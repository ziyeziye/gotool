package gotool

import (
	"github.com/ziyeziye/gotool/array"
	"github.com/ziyeziye/gotool/bcrypt"
	"github.com/ziyeziye/gotool/captcha"
	"github.com/ziyeziye/gotool/compression"
	"github.com/ziyeziye/gotool/conversion"
	"github.com/ziyeziye/gotool/convert"
	"github.com/ziyeziye/gotool/date"
	"github.com/ziyeziye/gotool/desensitized"
	"github.com/ziyeziye/gotool/id_utils"
	"github.com/ziyeziye/gotool/logs"
	"github.com/ziyeziye/gotool/openfile"
	"github.com/ziyeziye/gotool/page"
	"github.com/ziyeziye/gotool/pretty"
	"github.com/ziyeziye/gotool/request"
	"github.com/ziyeziye/gotool/str"
	"github.com/ziyeziye/gotool/tree"
)

// @Title tool
// @Description A simple, semantic and developer-friendly Golang tool development set
// @Page github.com/ziyeziye/gotool
// @Version v0.0.1
// @Author druidcaesa
// @Email hbsjzfynxm@gmail.com

var (
	StrArrayUtils     array.StrArray            //String 数据工具声明
	Logs              logs.Logs                 //log日志声明
	BcryptUtils       bcrypt.BcryptUtil         //加密工具声明
	DateUtil          date.Date                 //时间操作
	StrUtils          str.StrUtils              //字符串操作
	HttpUtils         = request.NewRequest()    //http工具
	ConvertUtils      convert.Convert           //公历转农历相关操作
	PageUtils         page.Page                 //分页插件
	IdUtils           id_utils.IdUtils          //id生成工具
	ZipUtils          compression.ZipUtils      //压缩和解压缩工具
	FileUtils         openfile.FileUtils        //IO操作工具
	CaptchaUtils      captcha.Captcha           //验证码工具
	DesensitizedUtils desensitized.Desensitized //敏感数据脱敏
	TreeUtils         tree.Tree                 //菜单树结构化工具
	PrettyUtils       pretty.PrettyUtils        //JSON打印格式化工具
	TypeConversion    conversion.Conversion     //类型转换，用于string，int，int64，float等数据转换，免去err的接收，和设置默认值
)
