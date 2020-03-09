package utils

import (
	"fmt"
	"strconv"

	"github.com/widuu/goini"
)

//ReadConfig duqu wenjian
var ReadConfig = goini.SetConfig("./config.ini")

//GetStringValue  get keyvalue
func GetStringValue(secName string, keyName string) string {
	return ReadConfig.GetValue(secName, keyName)
}

//GetIntValue  get keyvalue
func GetIntValue(secName string, keyName string) int {
	a := ReadConfig.GetValue(secName, keyName)
	b, error := strconv.Atoi(a)
	if error != nil {
		fmt.Println("字符串转换成整数失败")
		b = 0
	}
	return b
}
