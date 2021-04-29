package utils

import (
	"errors"
	"fmt"
	"time"
)

// GetHourDigitString 2021-04-28 16  --> "2021042816"
func GetHourDigitString(t time.Time) string {
	var year, month, day = t.Date()
	var hour = t.Hour()
	return fmt.Sprintf("%04d%02d%02d%02d", year, month, day, hour)
}

// GetWeekByDate 获取当前时间是第几周
func GetWeekByDate(t time.Time) ( week int) {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	fmt.Println(yearFirstDay)
	firstDayInWeek := int(yearFirstDay.Weekday())
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		weekOverDay := (yearDay - firstWeekDays) % 7
		if weekOverDay == 0 {
			week = (yearDay-firstWeekDays)/7 + 1
		} else {
			week = (yearDay-firstWeekDays)/7 + 2
		}
	}
	return
}

// IsSameDay 判断是否为相同的一天
func IsSameDay(a, b time.Time) bool {
	var aYear, aMonth, aDay = a.Date()
	var bYear, bMonth, bDay = b.Date()
	return aYear == bYear && aMonth == bMonth && aDay == bDay
}

// GetZeroAndLastTime 获取当天0点和23:59:59
func GetZeroAndLastTime(d time.Time) (zeroTime, lastTime time.Time) {
	zeroTime = time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	lastTime = time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
	return
}

// TimeToStr 时间类型转字符串
func TimeToStr(t time.Time, format ...string) string {
	if len(format) > 1 {
		panic(errors.New("只支持一个时间格式"))
	}
	if len(format) == 1 {
		return t.Format(format[0])
	}
	return t.Format("2006-01-02 15:04:05")
}

// StrToTime 字符串转时间类型
func StrToTime(tStr string, format ...string) (t time.Time, err error) {
	if len(format) > 1 {
		panic(errors.New("只支持一个时间格式"))
	}
	if len(format) == 1 {
		return time.Parse(format[0], tStr)
	}
	return time.Parse("2006-01-02 15:04:05", tStr)
}

// 获取明天的时间
func Tomorrow(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

// 获取昨天的时间
func Yesterday(t time.Time) time.Time {
	return t.AddDate(0, 0, -1)
}