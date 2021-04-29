package utils

import (
	"encoding/hex"
	"math"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// Round 返回将 val 根据指定精度 precision（十进制小数点后数字的数目）进行四舍五入的结果。precision 也可以是负数或零。
func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

// ValueIsInSlice value是否存在切片中
func ValueIsInSlice(slice []int, value int) (isIn bool) {
	for _, v := range slice {
		if v == value {
			isIn = true
			return
		}
	}
	return
}

// KeyIsInMap map中是否存在key
func KeyIsInMap(key string, dict map[string]interface{}) (isIn bool) {
	if _, isOk := dict[key]; isOk {
		isIn = true
	}
	return
}

// VerifyIdCard 身份证号码验证
func VerifyIdCard(idCard string) bool {
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

// GetUUID 生成string类型的uuid
func GetUUID() string {
	uid, _ := uuid.New().MarshalBinary()
	return hex.EncodeToString(uid)
}