package utils

import (
	"fmt"
	"time"
)

type DateFormat struct {
	format string
}

type Handler interface {
	parse(format *DateFormat)
}

type HandleFunc func(format *DateFormat)

func (f HandleFunc) parse(format *DateFormat){
	f(format)
}

func WithFormat(f string) HandleFunc {
	return func(format *DateFormat) {
		format.format = f
	}
}

// GetHourDigitString 2021-04-28 16  --> "2021042816"
func GetHourDigitString(t time.Time) string {
	var year, month, day = t.Date()
	var hour = t.Hour()
	return fmt.Sprintf("%04d%02d%02d%02d", year, month, day, hour)
}

// GetWeekByDate 获取当前时间是今年第几周
func GetWeekByDate(t time.Time) ( week int) {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
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
// TimeToStr(time.Now())
// TimeToStr(time.Now(), WithFormat("2006-01-02"))
func TimeToStr(t time.Time, format ...Handler) string {
	f := &DateFormat{}
	for _, v := range format {
		v.parse(f)
	}
	if f.format != "" {
		return t.Format(f.format)
	}
	return t.Format("2006-01-02 15:04:05")
}

// StrToTime 字符串转时间类型
// StrToTime("2021-05-01")
// StrToTime("2021-05-01 15:10:30", WithFormat("2006-01-02 15:04:05"))
func StrToTime(tStr string, format ...Handler) (t time.Time, err error) {
	f := &DateFormat{}
	for _, v := range format {
		v.parse(f)
	}
	if f.format != "" {
		return time.Parse(f.format, tStr)
	}
	return time.Parse("2006-01-02 15:04:05", tStr)
}

// Tomorrow 获取明天的时间
func Tomorrow(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

// Yesterday 获取昨天的时间
func Yesterday(t time.Time) time.Time {
	return t.AddDate(0, 0, -1)
}