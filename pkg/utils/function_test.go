package utils

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncRound(t *testing.T) {
	value := Round(12.564, 1)
	assert.Equal(t, 12.6, value)

	value = Round(12.514, 1)
	assert.Equal(t, 12.5, value)

	value = Round(12.5, 0)
	assert.Equal(t, float64(13), value)

	value = Round(12.5, -1)
	assert.Equal(t, float64(10), value)
}

func TestFuncValueIsInSlice(t *testing.T) {
	s := []int{1, 2, 3}
	assert.True(t, true, ValueIsInSlice(s, 3))
	assert.False(t, false, ValueIsInSlice(s, 6))
}

func TestFuncKeyIsInMap(t *testing.T) {
	m := map[string]interface{}{"key1": 1, "key2": 2}
	assert.True(t, true, KeyIsInMap("key1", m))
	assert.False(t, false, KeyIsInMap("key3", m))
}

func TestGetUUID(t *testing.T) {
	fmt.Println(GetUUID())
}

func TestGetEnv(t *testing.T) {
	assert.Equal(t, "111", GetEnv("test", "111"))

	err := os.Setenv("test", "222")
	assert.Equal(t, nil, err)
	assert.Equal(t, "222", GetEnv("test", "111"))
}

func TestGetFloatDecimalNumByStr(t *testing.T) {
	str, err := GetFloatDecimalNumByStr("3.141")
	if err != nil {
		log.Fatalf("%v", err)
	}
	assert.Equal(t, 3, str)
}