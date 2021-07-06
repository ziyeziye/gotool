package test

import (
	"fmt"
	"github.com/druidcaesa/gotool"
	"testing"
	"time"
)

func TestFormatToString(t *testing.T) {
	now := gotool.DateUtil.Now()
	toString := gotool.DateUtil.FormatToString(&now, "")
	fmt.Println(toString)
}

func TestDate_UnixToTime(t *testing.T) {
	fmt.Println(gotool.DateUtil.UnixToTime(gotool.DateUtil.Now().Unix()))
}

func TestDate_IsZero(t *testing.T) {
	t2 := time.Time{}
	zero := gotool.DateUtil.IsZero(&t2)
	fmt.Println(zero)
}

func TestInterpretStringToTimestamp(t *testing.T) {
	timestamp, err := gotool.DateUtil.InterpretStringToTimestamp("2021-05-04")
	if err != nil {
		gotool.Logs.ErrorLog().Println(err.Error())
	}
	fmt.Println(timestamp)
}
