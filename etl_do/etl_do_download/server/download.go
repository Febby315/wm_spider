package server

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"../utils"
)

//DownLoad a
func DownLoad(conf string) {
	param := map[string]interface{}{}
	param["appid"] = utils.GetStringValue("appflag_interface", "appid")
	param["passwd"] = utils.GetStringValue("appflag_interface", "passwd")
	param["tablename"] = "downImg"

	//获取查询的数据位置
	sysdatapoint := utils.HTTPPost(utils.GetStringValue("from_interface", "db_src_point")+"/sysdatapoint/findone", bson.M{"conf": param})

	sysdatapointmap := map[string]interface{}{}
	err := json.Unmarshal([]byte(sysdatapoint), &sysdatapointmap)

	//同步某些条件的数据
	qupdownloadparam := map[string]interface{}{}
	lastsysid := ""
	if err == nil && sysdatapointmap != nil && sysdatapointmap["last_sys_id"] != nil {
		lastsysid = sysdatapointmap["last_sys_id"].(string)
	}
	pointid := utils.GetStringValue("from_interface", "db_src_updownload_point_id")
	if pointid != "-1" {
		lastsysid = pointid
	}

	//组装flag的查询条件
	confmap := map[string]interface{}{}
	if conf != "" {
		confmap["info_flag"] = conf
	}

	if lastsysid != "" {
		confmap["_id"] = lastsysid
	}
	qupdownloadparam["conf"] = confmap

	//获取数据updownload
	result := utils.HTTPPost(utils.GetStringValue("from_interface", "db_src_server_url")+"/downImg/list", qupdownloadparam) //根据起始位置查询数据
	type resultObj struct {
		Data []map[string]interface{} `json:"data"`
	}
	resultarray := resultObj{}
	err = json.Unmarshal([]byte(result), &resultarray)
	if err == nil {
		_id := "0"
		for _, value := range resultarray.Data {
			_id = value["_id"].(string)
			imgurl := value["img_url"].(string)
			imgname := value["img_name"].(string)
			value["img_url"] = strings.TrimPrefix(strings.Replace(imgurl, imgname, "", -1), "{IMG}")
			utils.HTTPPost(utils.GetStringValue("to_interface", "db_tar_updownload_url"), value)
			fmt.Println(_id)
		}
		if _id != "0" {
			param["last_sys_id"] = _id
			utils.HTTPPost(utils.GetStringValue("from_interface", "db_src_point")+"/sysdatapoint/edit", param) //修改起始位置
		}

	}

	//fmt.Println(result)

}
