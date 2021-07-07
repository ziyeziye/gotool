package test

import (
	"fmt"
	"github.com/druidcaesa/gotool"
	"testing"
	"time"
)

func TestFormatToString(t *testing.T) {
	now := gotool.DateUtil.Now()
	toString := gotool.DateUtil.FormatToString(now, "YYYY-MM-DD hh:mm:ss")
	fmt.Println(toString)
	toString = gotool.DateUtil.FormatToString(now, "YYYYMMDD hhmmss")
	fmt.Println(toString)
}

func TestDate_UnixToTime(t *testing.T) {
	fmt.Println(gotool.DateUtil.UnixToTime(gotool.DateUtil.Now().Unix()))
}

func TestDate_IsZero(t *testing.T) {
	t2 := time.Time{}
	zero := gotool.DateUtil.IsZero(t2)
	fmt.Println(zero)
	zero = gotool.DateUtil.IsZero(gotool.DateUtil.Now())
	fmt.Println(zero)
}

func TestInterpretStringToTimestamp(t *testing.T) {
	timestamp, err := gotool.DateUtil.InterpretStringToTimestamp("2021-05-04 15:12:59", "YYYY-MM-DD hh:mm:ss")
	if err != nil {
		gotool.Logs.ErrorLog().Println(err.Error())
	}
	fmt.Println(timestamp)
}

func TestUnixToTime(t *testing.T) {
	unix := gotool.DateUtil.Now().Unix()
	fmt.Println("时间戳----------------------->", unix)
	toTime := gotool.DateUtil.UnixToTime(unix)
	fmt.Println(toTime)
}

func TestGetWeekDay(t *testing.T) {
	now := gotool.DateUtil.Now()
	day := gotool.DateUtil.GetWeekDay(now)
	fmt.Println("今天是-----------------周", day)
}

//时间计算
func TestTimeAddOrSub(t *testing.T) {
	now := gotool.DateUtil.Now()
	fmt.Println("现在时间是--------------------->", now)
	sub := gotool.DateUtil.MinuteAddOrSub(now, 10)
	fmt.Println("分钟计算结果-------------------->", sub)
	sub = gotool.DateUtil.MinuteAddOrSub(now, -10)
	fmt.Println("分钟计算结果-------------------->", sub)
	sub = gotool.DateUtil.HourAddOrSub(now, 10)
	fmt.Println("小时计算结果-------------------->", sub)
	sub = gotool.DateUtil.HourAddOrSub(now, -10)
	fmt.Println("小时计算结果-------------------->", sub)
	sub = gotool.DateUtil.DayAddOrSub(now, 10)
	fmt.Println("天计算结果-------------------->", sub)
	sub = gotool.DateUtil.DayAddOrSub(now, -10)
	fmt.Println("天计算结果-------------------->", sub)
}
