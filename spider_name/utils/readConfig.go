package utils

import (
	"log"
	"net/http"
	"strconv"

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
		log.Println("字符串转换成整数失败")
	}
	return val
}

//允许跨域及访问日志
func EnableXDA(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	log.Println(r.Method, r.URL)
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")
	return w
}
