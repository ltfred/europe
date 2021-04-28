package utils

import (
	"fmt"
	"time"
)

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