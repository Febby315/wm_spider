package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"./server"
	"./utils"
)

func sysDataTimer(appflagmap map[string]interface{}) {
	_intervaltime, _ := strconv.Atoi(utils.GetStringValue("timers", "intervalTime"))

	_timer := time.NewTicker(time.Duration(_intervaltime) * time.Second)
	fmt.Println("定时间隔秒", time.Duration(_intervaltime)*time.Second)
	for {
		select {
		case <-_timer.C:
			if priceConf := appflagmap["priceConf"]; priceConf != nil {
				server.SynPriceData(priceConf.(string))
			}
		}
	}
}

func main() {
	param := map[string]interface{}{}
	param["appid"] = utils.GetStringValue("appflag_interface", "appid")
	param["passwd"] = utils.GetStringValue("appflag_interface", "passwd")

	paramConf := map[string]interface{}{}
	paramConf["conf"] = param

	//用户的权限
	appflag := utils.HTTPPost(utils.GetStringValue("appflag_interface", "db_appflag_url")+"/sysappflag/findone", paramConf)
	//获取查询条件
	appflagmap := map[string]interface{}{}
	err := json.Unmarshal([]byte(appflag), &appflagmap)
	if err != nil {

	}
	sysDataTimer(appflagmap)

}
