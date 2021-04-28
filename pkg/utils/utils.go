package utils

import "os"

// 获取环境变量，不存在使用默认值
func GetEnv(key, defaultValue string) (value string) {
	value = os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return
}