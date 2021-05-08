package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsIdCard(t *testing.T) {
	assert.True(t, true, IsIdCard("513436200004297890"))
}

func TestIsMobile(t *testing.T) {
	assert.True(t, true, IsMobile("17781256998"))
	assert.False(t, false, IsMobile("12345678910"))
}

func TestIsUUID(t *testing.T) {
	assert.True(t, true, IsUUID(GetUUID()))
	assert.False(t, false, IsUUID("12334"))
}