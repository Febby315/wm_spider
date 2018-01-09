package listsource

import (
	"../../utils"
	"gopkg.in/mgo.v2/bson"
)

type (
	//ListSource 结构体
	ListSource struct {
		ID               bson.ObjectId          `json:"_id" bson:"_id"`                                         //ObjectID(必填)
		Status           int                    `json:"status" bson:"status"`                                   //状态(必填)
		Version          int                    `json:"version" bson:"version"`                                 //版本(必填)
		Timestamp        int64                  `json:"timestamp" bson:"timestamp,omitempty"`                   //时间戳
		Remark           string                 `json:"remark" bson:"remark,omitempty"`                         //备注
		CleanRuleContent string                 `json:"clean_rule_content" bson:"clean_rule_content,omitempty"` //清洗详情脚本
		CleanRuleList    string                 `json:"clean_rule_list" bson:"clean_rule_list,omitempty"`       //清洗列表脚本
		PathPre          string                 `json:"path_pre" bson:"path_pre,omitempty"`                     //图片前缀
		WebType          string                 `json:"web_type" bson:"web_type,omitempty"`                     //站点信息类型
		TableInfo        map[string]interface{} `json:"table_info" bson:"table_info,omitempty"`                 //扩展字段
		ListContent      string                 `json:"list_content" bson:"list_content,omitempty"`             //列表内容
		OperateDate      string                 `json:"operate_date" bson:"operate_date,omitempty"`             //创建时间
		ParentID         string                 `json:"parent_id" bson:"parent_id,omitempty"`                   //父ID
		ListURL          string                 `json:"list_url" bson:"list_url,omitempty"`                     //列表链接
		CurrentPage      int                    `json:"current_page" bson:"current_page,omitempty"`             //当前页号
	}
	//分页 结构体
	QueryParam struct {
		PageSize    int                    `json:"pageSize"`    //分页大小
		CurrentPage int                    `json:"currentPage"` //当前页号
		Conf        map[string]interface{} `json:"conf"`        //数据集合
	}
)

//获取表建立的链接
var conn = utils.GetConn("list_source")
