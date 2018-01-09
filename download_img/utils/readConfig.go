package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/widuu/goini"
)

//ReadConfig 读取配置文件
var readConfig = goini.SetConfig("./config.ini")

//GetStringValue get keyvalue
func GetStringValue(secName string, keyName string) string {
	return readConfig.GetValue(secName, keyName)
}

//GetIntValue get keyvalue
func GetIntValue(secName string, keyName string) int {
	num, err := strconv.Atoi(readConfig.GetValue(secName, keyName))
	if err != nil {
		fmt.Println("字符串转换成整数失败")
		num = 0
	}
	return num
}

//PrintReqLog 打印请求日志
func PrintReqLog(r *http.Request) {
	t, m, u := time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL
	fmt.Printf("%s %s %s\n", t, m, u)
}
