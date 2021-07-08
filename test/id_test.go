package test

import (
	"fmt"
	"github.com/druidcaesa/gotool"
	"testing"
)

func TestUUID(t *testing.T) {
	time, err := gotool.IdUtils.IdUUIDToTime(true)
	if err == nil {
		fmt.Println("根据时间生成UUID去除--------------------->'-'----->", time)
	}
	time, err = gotool.IdUtils.IdUUIDToTime(false)
	if err == nil {
		fmt.Println("根据时间生成不去除--------------------->'-'----->", time)
	}
	time = gotool.IdUtils.IdUUIDToRan(true)
	if err == nil {
		fmt.Println("根据随机数生成UUID去除--------------------->'-'----->", time)
	}
	time = gotool.IdUtils.IdUUIDToRan(false)
	if err == nil {
		fmt.Println("根据随机数生成不去除--------------------->'-'----->", time)
	}
}

func TestCreateCaptcha(t *testing.T) {
	captcha, err := gotool.IdUtils.CreateCaptcha(0)
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
