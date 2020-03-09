package sysappflag

import (
	"../../utils"
	"gopkg.in/mgo.v2/bson"
)

type (
	// 查询权限结构体
	Sysappflag struct {
		ID             bson.ObjectId `json:"_id" bson:"_id"`                       //唯一标识
		Appid          string        `json:"appid" bson:"appid"`                   //程序标识
		Passwd         string        `json:"passwd" bson:"passwd"`                 //访问密码
		PriceConf      string        `json:"priceConf" bson:"priceConf"`           //价格同步配置
		BuyConf        string        `json:"buyConf" bson:"buyConf"`               //求购同步配置
		SellConf       string        `json:"sellConf" bson:"sellConf"`             //供应同步配置
		TagArticleConf string        `json:"tagArticleConf" bson:"tagArticleConf"` //文章同步配置
		RecruitConf    string        `json:"recruitConf" bson:"recruitConf"`       //招聘同步配置
		UpdownloadConf string        `json:"updownloadConf" bson:"updownloadConf"` //附件同步配置
	}
	// 请求体结构体
	QueryParam struct {
		PageSize    int                    `json:"pageSize"`    //分页大小
		CurrentPage int                    `json:"currentPage"` //当前页码
		Conf        map[string]interface{} `json:"conf"`        //查询参数
	}
)

var conn = utils.GetConn("sysappflag")
