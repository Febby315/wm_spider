package sysdatapoint

import (
	"../../utils"
	"gopkg.in/mgo.v2/bson"
)

type (
	// 同步进度结构体
	Sysdatapoint struct {
		ID             bson.ObjectId `json:"_id" bson:"_id"`                             //唯一标识
		Appid          string        `json:"appid" bson:"appid"`                         //程序标识
		Passwd         string        `json:"passwd" bson:"passwd"`                       //访问密码
		Tablename      string        `json:"tablename" bson:"tablename"`                 //当前表名
		LastSysID      string        `json:"last_sys_id" bson:"last_sys_id"`             //最后同步ID
		CreatedDate    string        `json:"created_date" bson:"created_date,omitempty"` //创建时间
		LastUpdateDate string        `json:"last_update_date" bson:"last_update_date"`   //最后更新时间
		Version        int           `json:"version" bson:"version"`                     //版本号
	}
	// 请求体结构体
	QueryParam struct {
		PageSize    int                    `json:"pageSize"`    //分页大小
		CurrentPage int                    `json:"currentPage"` //当前页码
		Conf        map[string]interface{} `json:"conf"`        //查询参数
	}
)

var conn = utils.GetConn("sys_data_point")
