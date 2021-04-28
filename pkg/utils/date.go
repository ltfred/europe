package utils

import (
	"fmt"
	"time"
)

// 2021-04-28 16  --> "2021042816"
func GetHourDigitString(t time.Time) string {
	var year, month, day = t.Date()
	var hour = t.Hour()
	return fmt.Sprintf("%04d%02d%02d%02d", year, month, day, hour)
}

// 获取当前时间是第几周
func GetWeekByDate(t time.Time) int {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	fmt.Println(yearFirstDay)
	firstDayInWeek := int(yearFirstDay.Weekday())
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1

	}
	var week int
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
	return week

}

// 判断是否为相同的一天
func IsSameDay(a, b time.Time) bool {
	var aYear, aMonth, aDay = a.Date()
	var bYear, bMonth, bDay = b.Date()
	return aYear == bYear && aMonth == bMonth && aDay == bDay
}

// 获取当天0点和23:59:59
func GetZeroAndLastTime(d time.Time) (zeroTime, lastTime time.Time) {
	zeroTime = time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
	lastTime = time.Date(d.Year(), d.Month(), d.Day(), 23, 59, 59, 0, d.Location())
	return
}
