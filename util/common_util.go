package util

import "time"

/**
* 获取当前时区时间
 */
func GetNowTime() time.Time {
	l, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(l)
}

/**
*截取字符串
 */
func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	return string(r[start:end])
}
