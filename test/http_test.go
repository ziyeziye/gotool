package test

import (
	"fmt"
	"github.com/druidcaesa/gotool"
	"testing"
)

func TestGet(t *testing.T) {
	get, err := gotool.HttpUtils.Debug(true).Get("http://192.168.1.115:8080/user/list ")
	if err != nil {
		gotool.Logs.ErrorLog().Println(err)
	}
	s := string(get)
	fmt.Println(s)
}
