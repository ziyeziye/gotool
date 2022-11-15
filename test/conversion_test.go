package test

import (
	"fmt"
	"github.com/ziyeziye/gotool"
	"testing"
)

func TestStrToInt(t *testing.T) {
	toInt := gotool.TypeConversion.StrToInt("32e", 56)
	fmt.Println("出现异常，返回默认值：", toInt)
	toInt = gotool.TypeConversion.StrToInt("32e")
	fmt.Println("出现异常，无默认值：", toInt)
}

func TestIntToStr(t *testing.T) {
	var num int64 = 1344
	str := gotool.TypeConversion.IntToStr(num)
	fmt.Println("int64转字符串", str)
	str = gotool.TypeConversion.IntToStr("num")
	fmt.Println("直接输入字符串，会直接将字符串返回", str)
}

func TestStrToFloat64(t *testing.T) {
	toFloat32 := gotool.TypeConversion.StrToFloat32("12.78")
	fmt.Println(toFloat32)
	toFloat32 = gotool.TypeConversion.StrToFloat32("12.78e")
	fmt.Println(toFloat32)
}
