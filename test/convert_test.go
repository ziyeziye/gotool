package test

import (
	"fmt"
	"github.com/ziyeziye/gotool"
	"testing"
	"time"
)

func TestConvertTest(t *testing.T) {
	calendar := gotool.ConvertUtils.GregorianToLunarCalendar(2020, 2, 1)
	fmt.Println(calendar)
	gregorian := gotool.ConvertUtils.LunarToGregorian(calendar[0], calendar[1], calendar[2], false)
	fmt.Println(gregorian)
	days := gotool.ConvertUtils.GetLunarYearDays(2021)
	fmt.Println(days)
}

func TestDate(t *testing.T) {
	now := time.Now()
	sub := gotool.DateUtil.MinuteAddOrSub(now, -12)
	fmt.Println(sub)
	sub = gotool.DateUtil.MinuteAddOrSub(now, 12)
	fmt.Println(sub)
	orSub := gotool.DateUtil.HourAddOrSub(now, 12)
	fmt.Println(orSub)
	orSub = gotool.DateUtil.HourAddOrSub(now, -12)
	fmt.Println(orSub)
	duration := sub.Sub(orSub)
	fmt.Println(duration)
}
