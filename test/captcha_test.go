package test

import (
	"encoding/base64"
	"fmt"
	"github.com/ziyeziye/gotool"
	"testing"
)

func TestCaptcha(t *testing.T) {
	str := gotool.CaptchaUtils.GetRandStr(6)
	fmt.Println("生成验证码字符串------------------->", str)
	//生成图片byte数据，可以根据需要转成base64或者是image图片文件
	text := gotool.CaptchaUtils.ImgText(100, 40, str)
	sourcestring := base64.StdEncoding.EncodeToString(text)
	fmt.Println(sourcestring)
}
