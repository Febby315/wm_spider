package utils

import (
	"log"
	"net/http"
	"strconv"

	"github.com/widuu/goini"
)

//ReadConfig 读取配置文件
var readConfig = goini.SetConfig("./config.ini")

//获取配置项的值
func GetStringValue(secName string, keyName string) string {
	return readConfig.GetValue(secName, keyName)
}

////获取配置项的值
func GetIntValue(secName string, keyName string) int {
	num, err := strconv.Atoi(readConfig.GetValue(secName, keyName))
	if err != nil {
		log.Println("字符串转换成整数失败")
		num = 0
	}
	return num
}

//PrintReqLog 打印请求日志
func PrintReqLog(r *http.Request) {
	log.Printf(r.Method, r.URL)
}
