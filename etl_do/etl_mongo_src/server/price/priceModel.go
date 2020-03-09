package price

import (
	"../../utils"
	"gopkg.in/mgo.v2/bson"
)

type (

	// 价格结构体
	Price struct {
		ID          bson.ObjectId `json:"_id" bson:"_id"`
		PublishDate string        `json:"publishDate" bson:"publishDate,omitempty"`
		Marketname  string        `json:"marketname" bson:"marketname,omitempty"`

		Name      string `json:"name" bson:"name,omitempty"`
		PriceUnit string `json:"priceUnit" bson:"priceUnit,omitempty"`
		AvgPrice  string `json:"avgPrice" bson:"avgPrice,omitempty"`
		Area      string `json:"area" bson:"area,omitempty"`
		InfoFlag  string `json:"info_flag" bson:"info_flag,omitempty"`
		ParentID  string `json:"parent_id" bson:"parent_id,omitempty"`
		Timestamp int64  `json:"timestamp" bson:"timestamp"`
		Useid     string `json:"id" bson:"id"`
	}
	// 请求体结构体
	QueryParam struct {
		PageSize    int                    `json:"pageSize"`    //分页大小
		CurrentPage int                    `json:"currentPage"` //当前页码
		Conf        map[string]interface{} `json:"conf"`        //查询参数
	}
)

var conn = utils.GetConn("price")
