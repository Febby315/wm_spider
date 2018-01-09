package spiderconfig

import (
	"../../utils"
	"gopkg.in/mgo.v2/bson"
)

type (
	//SpiderConfig 结构体
	SpiderConfig struct {
		ID               bson.ObjectId          `json:"_id" bson:"_id"`                                         //ObjectID(必填)
		Status           int                    `json:"status" bson:"status"`                                   //状态(必填)
		Version          int                    `json:"version" bson:"version"`                                 //版本(必填)
		Remark           string                 `json:"remark" bson:"remark,omitempty"`                         //备注
		PageCount        int                    `json:"page_count" bson:"page_count,omitempty"`                 //当前页号
		MorePageModel    string                 `json:"more_page_model" bson:"more_page_model,omitempty"`       //分页模板
		CleanRuleContent string                 `json:"clean_rule_content" bson:"clean_rule_content,omitempty"` //详情清洗脚本
		CleanRuleList    string                 `json:"clean_rule_list" bson:"clean_rule_list,omitempty"`       //列表清洗脚本
		PathPre          string                 `json:"path_pre" bson:"path_pre,omitempty"`                     //图片前缀
		WebType          string                 `json:"web_type" bson:"web_type,omitempty"`                     //站点信息类型
		TableInfo        map[string]interface{} `json:"table_info" bson:"table_info,omitempty"`                 //扩展字段
	}
	//分页 结构体
	QueryParam struct {
		PageSize    int                    `json:"pageSize"`    //分页大小
		CurrentPage int                    `json:"currentPage"` //当前页号
		Conf        map[string]interface{} `json:"conf"`        //数据集合
	}
)

//获取表建立的链接
var conn = utils.GetConn("db_spider_config")
