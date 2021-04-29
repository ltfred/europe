package utils

import (
	"math"
)

// 返回将 val 根据指定精度 precision（十进制小数点后数字的数目）进行四舍五入的结果。precision 也可以是负数或零。
func FuncRound(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

// value是否存在切片中
func FuncValueIsInSlice(slice []int, value int) (isIn bool) {
	for _, v := range slice {
		if v == value {
			isIn = true
			return
		}
	}
	return
}

// map中是否存在key
func FuncKeyIsInMap(key string, dict map[string]interface{}) (isIn bool) {
	if _, isOk := dict[key]; isOk {
		isIn = true
	}
	return
}