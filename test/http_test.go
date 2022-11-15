package test

import (
	"fmt"
	"github.com/ziyeziye/gotool"
	"testing"
)

func TestGet(t *testing.T) {
	get, err := gotool.HttpUtils.Debug(true).Get("http://www.baidu.com")
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
	}
	s, _ := get.Content()
	fmt.Println(s)
}
