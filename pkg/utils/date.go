package utils

import (
	"errors"
	"fmt"
	"time"
)

type Date struct {
	startTime time.Time
	endTime   time.Time
	delta     int
	format    string
}

type Handler interface {
	parse(date *Date)
}

type HandleFunc func(date *Date)

func (f HandleFunc) parse(date *Date) {
	f(date)
}

func WithFormat(f string) HandleFunc {
	return func(date *Date) {
		date.format = f
	}
}

func WithEndTime(t time.Time) HandleFunc {
	return func(date *Date) {
		date.endTime = t
	}
}

func WithDelta(d int) HandleFunc {
	return func(date *Date) {
		date.delta = d
	}
}

// GetHourDigitString 2021-04-28 16  --> "2021042816"
func GetHourDigitString(t time.Time) string {
	var year, month, day = t.Date()
	var hour = t.Hour()
	return fmt.Sprintf("%04d%02d%02d%02d", year, month, day, hour)
}

// GetWeekByDate 获取当前时间是今年第几周
func GetWeekByDate(t time.Time) (week int) {
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

// TimeCombine 获取一段日期的起止时间
// start, end := TimeCombine(time.Now())
// yesterday, today := TimeCombine(time.Now(), WithDelta(-1))
func TimeCombine(startTime time.Time, opts ...Handler) (start time.Time, end time.Time) {
	f := &Date{}
	for _, v := range opts {
		v.parse(f)
	}
	if f.delta != 0 && f.endTime != *new(time.Time) {
		panic(errors.New("can not specify end_date and delta_days at the same time"))
	}
	if f.delta != 0 {
		if f.delta >= 0 {
			start = startTime
			end = startTime.AddDate(0, 0, f.delta)
		} else {
			end = startTime
			start = startTime.AddDate(0,0, f.delta)
		}
		start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0,0,0, start.Location())
		end = time.Date(end.Year(), end.Month(), end.Day(), 23, 59,59,0, end.Location())
		return
	}
	if f.endTime == *new(time.Time) {
		start = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0,0,0, startTime.Location())
		end = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 23, 59,59,0, startTime.Location())
		return
	}
	start = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0,0,0, startTime.Location())
	end = time.Date(f.endTime.Year(), f.endTime.Month(), f.endTime.Day(), 23, 59,59,0, startTime.Location())
	return
}

// TimeToStr 时间类型转字符串
// TimeToStr(time.Now())
// TimeToStr(time.Now(), WithFormat("2006-01-02"))
func TimeToStr(t time.Time, format ...Handler) string {
	f := &Date{}
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
	f := &Date{}
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
