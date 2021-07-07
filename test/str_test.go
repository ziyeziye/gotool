package test

import (
	"fmt"
	"github.com/druidcaesa/gotool"
	"testing"
)

func TestRemoveSuffix(t *testing.T) {
	fullFilename := "test.txt"
	suffix, _ := gotool.StrUtils.RemoveSuffix(fullFilename)
	fmt.Println(suffix)
	fullFilename = "/root/home/test.txt"
	suffix, _ = gotool.StrUtils.RemoveSuffix(fullFilename)
	fmt.Println(suffix)
}

func TestStringReplacePlaceholder(t *testing.T) {
	s := "你是我的{},我是你的{}"
	placeholder, err := gotool.StrUtils.ReplacePlaceholder(s, "唯一", "所有")
	if err == nil {
		fmt.Println(placeholder)
	}
}

func TestGetSuffix(t *testing.T) {
	fullFilename := "test.txt"
	suffix, _ := gotool.StrUtils.GetSuffix(fullFilename)
	fmt.Println(suffix)
	fullFilename = "/root/home/test.txt"
	suffix, _ = gotool.StrUtils.GetSuffix(fullFilename)
	fmt.Println(suffix)
}

func TestHasStr(t *testing.T) {
	str := ""
	empty := gotool.StrUtils.HasEmpty(str)
	fmt.Println(empty)
	str = "11111"
	empty = gotool.StrUtils.HasEmpty(str)
	fmt.Println(empty)
}
