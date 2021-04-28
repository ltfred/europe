package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 2021-04-28 16  --> "2021042816"
func TestGetHourDigitString(t *testing.T) {
	date, err := time.Parse("2006-01-02 15", "2021-04-28 16")
	assert.Equal(t, nil, err)
	dateStr := GetHourDigitString(date)
	assert.Equal(t, "2021042816", dateStr)
}

func TestGetWeekByDate(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2021-04-28")
	assert.Equal(t, nil, err)
	weeky := GetWeekByDate(date)
	assert.Equal(t, 18, weeky)
}