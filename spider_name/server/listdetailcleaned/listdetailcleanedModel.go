package listdetailcleaned

import (
	"../../utils"
	"gopkg.in/mgo.v2/bson"
)

type (
	//Listdetailcleaned 结构体
	Listdetailcleaned struct {
		ID               bson.ObjectId          `json:"_id" bson:"_id"`                                     //ObjectID(必填)
		DealStatus       int                    `json:"dealStatus" bson:"dealStatus"`                       //状态(必填)
		Version          int                    `json:"version" bson:"version"`                             //版本(必填)
		Timestamp        int64                  `json:"timestamp" bson:"timestamp,omitempty"`               //时间戳
		Remark           string                 `json:"remark" bson:"remark,omitempty"`                     //备注
		PathPre          string                 `json:"path_pre" bson:"path_pre,omitempty"`                 //图片前缀
		WebType          string                 `json:"web_type" bson:"web_type,omitempty"`                 //站点信息类型
		TableInfo        map[string]interface{} `json:"table_info" bson:"table_info,omitempty"`             //扩展字段
		OperateDate      string                 `json:"operate_date" bson:"operate_date,omitempty"`         //创建时间
		ParentID         string                 `json:"parent_id" bson:"parent_id,omitempty"`               //父ID
		ListURL          string                 `json:"list_url" bson:"list_url,omitempty"`                 //列表链接
		DetailURL        string                 `json:"detail_url" bson:"detail_url,omitempty"`             //详情链接
		PubTime          string                 `json:"pub_time" bson:"pub_time,omitempty"`                 //发布时间
		Title            string                 `json:"title" bson:"title,omitempty"`                       //标题
		ContentSour      string                 `json:"content_sour" bson:"content_sour,omitempty"`         //详情内容(图片地址未处理)
		ListImageSour    string                 `json:"listImageSour" bson:"listImageSour,omitempty"`       //列表图片集合
		Summary          string                 `json:"summary" bson:"summary,omitempty"`                   //内容摘要
		ContentImageSour string                 `json:"contentImageSour" bson:"contentImageSour,omitempty"` //详情图片集合
		Content          string                 `json:"content" bson:"content,omitempty"`                   //详情内容(图片地址已处理)
		ListShowImage    string                 `json:"listShowImage" bson:"listShowImage,omitempty"`       //所有图片
		AttachFilesSour  string                 `json:"attachFilesSour" bson:"attachFilesSour,omitempty"`   //附件
	}
	//QueryParam 分页 结构体
	QueryParam struct {
		PageSize    int                    `json:"pageSize"`    //分页大小
		CurrentPage int                    `json:"currentPage"` //当前页号
		Conf        map[string]interface{} `json:"conf"`        //数据集合
	}
)

//获取表建立的链接
var conn = utils.GetConn("list_detail_cleaned")
