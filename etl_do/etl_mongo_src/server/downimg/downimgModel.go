package downimg

import (
	"../../utils"
	"gopkg.in/mgo.v2/bson"
)

type (
	// 附件结构体
	Downimg struct {
		ID bson.ObjectId `json:"_id" bson:"_id"`

		ImgURL     string `json:"img_url" bson:"img_url,omitempty"`
		ImgName    string `json:"img_name" bson:"img_name,omitempty"`
		Timestamp  int64  `json:"timestamp" bson:"timestamp,omitempty"`
		CreateTime string `json:"create_time" bson:"create_time,omitempty"`
		UpdateTime string `json:"update_time" bson:"update_time,omitempty"`
		ImgSrc     string `json:"img_src" bson:"img_src,omitempty"`
		DealStatus int    `json:"dealStatus" bson:"dealStatus"`
		Version    int    `json:"version" bson:"version"`
	}
	// 请求体结构体
	QueryParam struct {
		PageSize    int                    `json:"pageSize"`    //分页大小
		CurrentPage int                    `json:"currentPage"` //当前页码
		Conf        map[string]interface{} `json:"conf"`        //查询参数
	}
)

var conn = utils.GetConn("downImg")
