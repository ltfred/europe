package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	value := GetEnv("test", "1")
	assert.Equal(t, "1", value)

	err := os.Setenv("test1", "2")
	assert.Equal(t, nil, err)
	value = GetEnv("test1", "1")
	assert.Equal(t, 2, value)
}

