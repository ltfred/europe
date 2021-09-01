package utils

import (
	"encoding/hex"
	"math"
	"os"
	"strconv"
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

// GetFloatDecimalNumByStr 获取浮点数的小数位数
func GetFloatDecimalNumByStr(numStr string) (int, error) {
	_, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, err
	}
	arr := strings.Split(numStr, ".")
	if len(arr) < 2 {
		return 0, nil
	}
	return len(arr[1]), nil
}