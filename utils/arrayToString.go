package utils

import (
	"fmt"
	"strings"
)

// 将数组格式转化为字符串
func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
