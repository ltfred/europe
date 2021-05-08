package utils

import (
	"encoding/hex"
	"github.com/google/uuid"
	"math"
	"os"
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

// GetUUID 生成string类型的uuid
func GetUUID() string {
	uid, _ := uuid.New().MarshalBinary()
	return hex.EncodeToString(uid)
}

// GetEnv 获取环境变量,不存在返回默认值
func GetEnv(key, defaultValue string) (value string) {
	value = os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return
}
