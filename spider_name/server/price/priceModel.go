package price

import (
	"../../utils"
	"gopkg.in/mgo.v2/bson"
)

type (
	//Price 结构体
	Price struct {
		ID          bson.ObjectId `json:"_id" bson:"_id"`                           //ObjectID(必填)
		Status      int           `json:"status" bson:"status"`                     //状态(必填)
		Version     int           `json:"version" bson:"version"`                   //版本(必填)
		Useid       string        `json:"id" bson:"id,omitempty"`                   //唯一标识
		Timestamp   int64         `json:"timestamp" bson:"timestamp,omitempty"`     //时间戳
		PublishDate string        `json:"publishDate" bson:"publishDate,omitempty"` //发布时间
		Name        string        `json:"name" bson:"name,omitempty"`               //名称
		PriceUnit   string        `json:"priceUnit" bson:"priceUnit,omitempty"`     //价格单位
		AvgPrice    string        `json:"avgPrice" bson:"avgPrice,omitempty"`       //价格
		Area        string        `json:"area" bson:"area,omitempty"`               //地区
		Marketname  string        `json:"marketname" bson:"marketname,omitempty"`   //市场
		InfoFlag    string        `json:"info_flag" bson:"info_flag,omitempty"`     //用于检索
		ParentID    string        `json:"parent_id" bson:"parent_id,omitempty"`     //父ID
		Source      string        `json:"source" bson:"source,omitempty"`           //来源
		Classify    string        `json:"classify" bson:"classify,omitempty"`       //分类
		Type        string        `json:"type" bson:"type,omitempty"`               //类型
	}
	//QueryParam 分页 结构体
	QueryParam struct {
		PageSize    int                    `json:"pageSize"`    //分页大小
		CurrentPage int                    `json:"currentPage"` //当前页号
		Conf        map[string]interface{} `json:"conf"`        //数据集合
	}
)

//获取表建立的链接
var conn = utils.GetConn("price")
