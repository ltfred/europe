package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyIdCard(t *testing.T) {
	assert.True(t, true, IsIdCard("513436200004297890"))
}

func TestVerifyMobile(t *testing.T) {
	assert.True(t, true, IsMobile("17781256998"))
	assert.False(t, false, IsMobile("12345678910"))
}