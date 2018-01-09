package article

import (
	"../../utils"

	"gopkg.in/mgo.v2/bson"
)

type (
	//Article 结构体
	Article struct {
		ID                     bson.ObjectId `json:"_id" bson:"_id"`                                                 //ObjectID(必填)
		DealStatus             int           `json:"dealStatus" bson:"dealStatus"`                                   //状态(必填)
		Version                int           `json:"version" bson:"version"`                                         //版本(必填)
		ArticleID              string        `json:"articleID" bson:"articleID,omitempty"`                           //文章ID
		Timestamp              int64         `json:"timestamp" bson:"timestamp,omitempty"`                           //时间戳
		Source                 string        `json:"source" bson:"source,omitempty"`                                 //来源
		SourceURL              string        `json:"sourceUrl" bson:"sourceUrl,omitempty"`                           //来源链接
		ArticleURL             string        `json:"articleURL" bson:"articleURL,omitempty"`                         //文章链接
		ArticleCreateDateTime  string        `json:"articleCreateDateTime" bson:"articleCreateDateTime,omitempty"`   //文章创建时间
		ArticleTitle           string        `json:"articleTitle" bson:"articleTitle,omitempty"`                     //文章标题
		ArticleContent         string        `json:"articleContent" bson:"articleContent,omitempty"`                 //文章内容
		ArticleImageList       string        `json:"articleImageList" bson:"articleImageList,omitempty"`             //文章图片列表
		ContentImageList       string        `json:"contentImageList" bson:"contentImageList,omitempty"`             //内容图片列表
		ArticleRefineTimestamp int64         `json:"articleRefineTimestamp" bson:"articleRefineTimestamp,omitempty"` //文章发布时间戳
		SourceHead             string        `json:"sourceHead" bson:"sourceHead,omitempty"`                         //
		WebSource              string        `json:"webSource" bson:"webSource,omitempty"`                           //来源网站
		ArticleRefineTime      string        `json:"articleRefineTime" bson:"articleRefineTime,omitempty"`           //文章发布时间
		ArticleAbstract        string        `json:"articleAbstract" bson:"articleAbstract,omitempty"`               //文章摘要
		ArticleClassification  string        `json:"articleClassification" bson:"articleClassification,omitempty"`   //文章分类
		ArticleNavTitle        string        `json:"articleNavTitle" bson:"articleNavTitle,omitempty"`               //文章栏目标题
		SourceTag              string        `json:"sourceTag" bson:"sourceTag,omitempty"`                           //源标签
		EsNewsClassify         string        `json:"esNewsClassify" bson:"esNewsClassify,omitempty"`                 //新闻分类
		ParentID               string        `json:"parent_id" bson:"parent_id,omitempty"`                           //父ID
		Infoflag               string        `json:"info_flag" bson:"info_flag,omitempty"`                           //用于检索
		ExtcolumnValue         string        `json:"ext_column_value" bson:"ext_column_value,omitempty"`             //
		Area                   string        `json:"area" bson:"area,omitempty"`                                     //地区
		PrivateField           string        `json:"privateField" bson:"privateField,omitempty"`                     //私有字段
		AttachFilesList        string        `json:"attachFilesList" bson:"attachFilesList,omitempty"`               //附件列表
	}
	//QueryParam 分页 结构体
	QueryParam struct {
		PageSize    int                    `json:"pageSize"`    //分页大小
		CurrentPage int                    `json:"currentPage"` //当前页号
		Conf        map[string]interface{} `json:"conf"`        //数据集合
	}
)

//获取表建立的链接
var conn = utils.GetConn("article")
