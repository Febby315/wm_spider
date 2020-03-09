package tagArticle

import (
	"../../utils"
	"gopkg.in/mgo.v2/bson"
)

type (
	// 文章结构体
	TagArticle struct {
		ID                     bson.ObjectId `json:"_id" bson:"_id"`                                                 //ObjectID(必填)
		DealStatus             int           `json:"dealStatus" bson:"dealStatus"`                                   //状态(必填)
		Version                int           `json:"version" bson:"version"`                                         //版本(必填)
		ArticleID              string        `json:"articleID" bson:"articleID,omitempty"`                           //文章ID
		Source                 string        `json:"source" bson:"source,omitempty"`                                 //来源
		SourceURL              string        `json:"sourceUrl" bson:"sourceUrl,omitempty"`                           //来源链接
		ArticleURL             string        `json:"articleURL" bson:"articleURL,omitempty"`                         //文章链接
		ArticleCreateDateTime  string        `json:"articleCreateDateTime" bson:"articleCreateDateTime,omitempty"`   //文章创建时间
		ArticleTitle           string        `json:"articleTitle" bson:"articleTitle,omitempty"`                     //文章标题
		ArticleContent         string        `json:"articleContent" bson:"articleContent,omitempty"`                 //文章内容
		ContentImageList       string        `json:"contentImageList" bson:"contentImageList,omitempty"`             //内容图片集合
		ArticleImageList       string        `json:"articleImageList" bson:"articleImageList,omitempty"`             //文章图片集合
		ArticleRefineTimestamp int64         `json:"articleRefineTimestamp" bson:"articleRefineTimestamp,omitempty"` //文章发布时间戳
		KeyStatus              string        `json:"keyStatus" bson:"keyStatus,omitempty"`                           //
		KeyWordList            string        `json:"keyWordList" bson:"keyWordList,omitempty"`                       //关键字
		KeyTitleList           string        `json:"keyTitleList" bson:"keyTitleList,omitempty"`                     //标题关键字
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
		Area                   string        `json:"area" bson:"area,omitempty"`                                     //
		PrivateField           string        `json:"privateField" bson:"privateField,omitempty"`                     //
		AttachFilesList        string        `json:"attachFilesList" bson:"attachFilesList,omitempty"`               //附件
	}
	// 请求体结构体
	QueryParam struct {
		PageSize    int                    `json:"pageSize"`    //分页大小
		CurrentPage int                    `json:"currentPage"` //当前页码
		Conf        map[string]interface{} `json:"conf"`        //查询参数
	}
)

var conn = utils.GetConn("tagArticle")
