package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncRound(t *testing.T) {
	value := FuncRound(12.564, 1)
	assert.Equal(t, 12.6, value)

	value = FuncRound(12.514, 1)
	assert.Equal(t, 12.5, value)

	value = FuncRound(12.5, 0)
	assert.Equal(t, float64(13), value)

	value = FuncRound(12.5, -1)
	assert.Equal(t, float64(10), value)
}

func TestFuncValueIsInSlice(t *testing.T) {
	s := []int{1, 2, 3}
	assert.True(t, true, FuncValueIsInSlice(s, 3))
	assert.False(t, false, FuncValueIsInSlice(s, 6))
}

func TestFuncKeyIsInMap(t *testing.T) {
	m := map[string]interface{}{"key1": 1, "key2":2}
	assert.True(t, true, FuncKeyIsInMap("key1", m))
	assert.False(t, false, FuncKeyIsInMap("key3", m))
}