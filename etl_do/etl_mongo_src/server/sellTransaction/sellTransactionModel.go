package sellTransaction

import (
	"../../utils"
	"gopkg.in/mgo.v2/bson"
)

type (
	// 供应结构体
	SellTransaction struct {
		ID                    bson.ObjectId `json:"_id" bson:"_id"`                                               //ObjectID(必填)
		Status                int           `json:"status" bson:"status"`                                         //状态(必填)
		Version               int           `json:"version" bson:"version"`                                       //版本(必填)
		Timestamp             int64         `json:"timestamp" bson:"timestamp,omitempty"`                         //时间戳
		TransactionID         string        `json:"transactionId" bson:"transactionId,omitempty"`                 //唯一标识
		Source                string        `json:"source" bson:"source,omitempty"`                               //来源
		SourceURL             string        `json:"sourceUrl" bson:"sourceUrl,omitempty"`                         //来源链接
		Content               string        `json:"content" bson:"content,omitempty"`                             //内容
		ContentURL            string        `json:"contentUrl" bson:"contentUrl,omitempty"`                       //详情链接
		Title                 string        `json:"title" bson:"title,omitempty"`                                 //标题
		Price                 string        `json:"price" bson:"price,omitempty"`                                 //价格
		UnitName              string        `json:"unitName" bson:"unitName,omitempty"`                           //单位
		PublishTime           string        `json:"publishTime" bson:"publishTime,omitempty"`                     //发布时间
		ExpiryDate            string        `json:"expiryDate" bson:"expiryDate,omitempty"`                       //有效期
		UserName              string        `json:"userName" bson:"userName,omitempty"`                           //用户名
		QQ                    string        `json:"qq" bson:"qq,omitempty"`                                       //QQ
		Phone                 string        `json:"phone" bson:"phone,omitempty"`                                 //手机
		Cellphone             string        `json:"cellphone" bson:"cellphone,omitempty"`                         //电话
		Address               string        `json:"address" bson:"address,omitempty"`                             //地址
		PickTime              string        `json:"pickTime" bson:"pickTime,omitempty"`                           //爬取时间
		ParentID              string        `json:"parent_id" bson:"parent_id,omitempty"`                         //父ID
		InfoFlag              string        `json:"info_flag" bson:"info_flag,omitempty"`                         //用于检索
		TitleMd5              string        `json:"titleMd5" bson:"titleMd5,omitempty"`                           //
		ContentMd5            string        `json:"contentMd5" bson:"contentMd5,omitempty"`                       //
		Area                  string        `json:"area" bson:"area,omitempty"`                                   //地区（陕西农业网供求特殊需求）
		ArticleClassification string        `json:"articleClassification" bson:"articleClassification,omitempty"` //文章分类（陕西农业网供求特殊需求）
	}
	// 请求体结构体
	QueryParam struct {
		PageSize    int                    `json:"pageSize"`    //分页大小
		CurrentPage int                    `json:"currentPage"` //当前页码
		Conf        map[string]interface{} `json:"conf"`        //查询参数
	}
)

var conn = utils.GetConn("sellTransaction")
