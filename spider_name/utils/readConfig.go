package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/widuu/goini"
)

//ReadConfig 读取配置文件
var ReadConfig = goini.SetConfig("./config.ini")

//GetStringValue 获取字符串配置
func GetStringValue(secName string, keyName string) string {
	return ReadConfig.GetValue(secName, keyName)
}

//GetIntValue 获取整型配置
func GetIntValue(secName string, keyName string) int {
	val, err := strconv.Atoi(ReadConfig.GetValue(secName, keyName))
	if err != nil {
		fmt.Println("字符串转换成整数失败")
	}
	return val
}

//PrintReqLog 打印请求日志
func PrintReqLog(r *http.Request) {
	t, m, u := time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL
	fmt.Printf("%s %s %s\n", t, m, u)
	// switch runtime.GOOS {
	// case "windows":
	// 	fmt.Printf("%s %s %s\n", t, m, u)
	// case "linux":
	// 	d, b, f := 0, 40, 32
	// 	fmt.Printf("%c %d;%d;%dm %s %c 0m %s %s\n", 0x1B, d, b, f, t, 0x1B, m, u)
	// default:
	// 	fmt.Printf("[%s]\t%s\t%s\n", t, m, u)
	// }
}
