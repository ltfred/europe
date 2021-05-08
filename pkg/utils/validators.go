package utils

import (
	"regexp"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// IsIdCard 身份证号码验证
func IsIdCard(idCard string) bool {
	idCardLen := strings.Count(idCard, "")
	if idCardLen-1 == 18 {
		regular18 := "^[1-9]\\d{5}(18|19|([23]\\d))\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{3}[0-9Xx]$"
		reg := regexp.MustCompile(regular18)
		return reg.MatchString(idCard)
	} else {
		regular15 := "^[1-9]\\d{5}\\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\\d{2}[0-9Xx]$"
		reg := regexp.MustCompile(regular15)
		return reg.MatchString(idCard)
	}
}

// IsMobile 手机号验证
func IsMobile(mobile string) bool {
	regular := "^(1[3-9])\\d{9}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

// IsUUID 验证是否是UUID
func IsUUID(u string) bool {
	_, err := uuid.FromString(u)
	if err != nil {
		return true
	}
	return false
}
