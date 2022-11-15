package test

import (
	"github.com/ziyeziye/gotool"
	"testing"
)

func TestLogs(t *testing.T) {
	gotool.Logs.ErrorLog().Println("Error日志测试")
	gotool.Logs.InfoLog().Println("Info日志测试")
	gotool.Logs.DebugLog().Println("Debug日志测试")
}
