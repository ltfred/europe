package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetHourDigitString(t *testing.T) {
	date, err := time.Parse("2006-01-02 15", "2021-04-28 16")
	assert.Equal(t, nil, err)
	dateStr := GetHourDigitString(date)
	assert.Equal(t, "2021042816", dateStr)
}

func TestGetWeekByDate(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2021-04-28")
	assert.Equal(t, nil, err)
	week := GetWeekByDate(date)
	assert.Equal(t, 18, week)
}

func TestIsSameDay(t *testing.T) {
	date1, _ := time.Parse("2006-01-02", "2021-04-28")
	date2, _ := time.Parse("2006-01-02", "2021-04-29")
	assert.False(t, IsSameDay(date1, date2))

	date1, _ = time.Parse("2006-01-02", "2021-04-28")
	date2, _ = time.Parse("2006-01-02", "2021-04-28")
	assert.True(t, IsSameDay(date1, date2))

	date1, _ = time.Parse("2006-01-02", "2020-04-28")
	date2, _ = time.Parse("2006-01-02", "2021-04-28")
	assert.False(t, IsSameDay(date1, date2))
}

func TestTimeToStr(t *testing.T) {
	date1, _ := time.Parse("2006-01-02", "2021-04-28")
	assert.Equal(t, "2021-04-28 00:00:00", TimeToStr(date1))
	assert.Equal(t, "2021-04-28 00", TimeToStr(date1, WithFormat("2006-01-02 15")))
}

func TestStrToTime(t *testing.T) {
	zero, _ := time.Parse("2006-01-02 15:04:05", "2021-04-28 00:00:00")
	ti, err := StrToTime("2021-04-28 00:00:00")
	assert.Equal(t, nil, err)
	assert.Equal(t, zero, ti)

	zero, _ = time.Parse("2006-01-02", "2021-04-28")
	ti, err = StrToTime("2021-04-28", WithFormat("2006-01-02"))
	assert.Equal(t, zero, ti)
}

func TestTomorrow(t *testing.T) {
	date1, _ := time.Parse("2006-01-02", "2021-04-28")
	date2, _ := time.Parse("2006-01-02", "2021-04-29")
	assert.Equal(t, date2, Tomorrow(date1))
}

func TestYesterday(t *testing.T) {
	date1, _ := time.Parse("2006-01-02", "2021-04-28")
	date2, _ := time.Parse("2006-01-02", "2021-04-29")
	assert.Equal(t, date1, Yesterday(date2))
}

func TestTimeCombine(t *testing.T) {
	dateTime, _ := time.Parse("2006-01-02", "2021-04-28")
	start, _ := time.Parse("2006-01-02 15:04:05", "2021-04-28 00:00:00")
	end, _ := time.Parse("2006-01-02 15:04:05", "2021-04-28 23:59:59")
	startTime, endTime, _ := TimeCombine(dateTime)
	assert.Equal(t, start, startTime)
	assert.Equal(t, end, endTime)

	startTime, endTime, _ = TimeCombine(dateTime, WithDelta(3))
	end, _ = time.Parse("2006-01-02 15:04:05", "2021-05-01 23:59:59")
	assert.Equal(t, end, endTime)

	startTime, endTime, _ = TimeCombine(dateTime, WithEndTime(dateTime.AddDate(0, 0, 1)))
	end, _ = time.Parse("2006-01-02 15:04:05", "2021-04-29 23:59:59")
	assert.Equal(t, end, endTime)
}
