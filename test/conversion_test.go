package test

import (
	"fmt"
	"github.com/druidcaesa/gotool"
	"testing"
)

func TestStrToInt(t *testing.T) {
	toInt := gotool.TypeConversion.StrToInt("32e", 56)
	fmt.Println(toInt)
}

func TestIntToStr(t *testing.T) {
	var num int64 = 1344
	str := gotool.TypeConversion.IntToStr(num)
	fmt.Println(str)
}
