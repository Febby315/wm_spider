package recruit

import (
	"../../utils"
	"gopkg.in/mgo.v2/bson"
)

type (
	// 招聘结构体
	Recruit struct {
		ID              bson.ObjectId `json:"_id" bson:"_id"`                                   //ObjectID(必填)
		Status          int           `json:"status" bson:"status"`                             //状态(必填)
		Version         int           `json:"version" bson:"version"`                           //版本(必填)
		Timestamp       int64         `json:"timestamp" bson:"timestamp,omitempty"`             //时间戳
		Title           string        `json:"title" bson:"title,omitempty"`                     //标题
		Content         string        `json:"content" bson:"content,omitempty"`                 //内容详情
		ContentUrl      string        `json:"contentUrl" bson:"contentUrl,omitempty"`           //详情链接
		SourceUrl       string        `json:"sourceUrl" bson:"sourceUrl,omitempty"`             //来源网址
		WebSource       string        `json:"webSource" bson:"webSource,omitempty"`             //来源网站
		PickTime        string        `json:"pickTime" bson:"pickTime,omitempty"`               //爬取时间
		PublishTime     string        `json:"publicTime" bson:"publicTime,omitempty"`           //发布时间
		ExpiryDate      string        `json:"expiryDate" bson:"expiryDate,omitempty"`           //截止时间
		Province        string        `json:"privice" bson:"privice,omitempty"`                 //省份
		City            string        `json:"city" bson:"city,omitempty"`                       //城市
		WorkType        string        `json:"workType" bson:"workType,omitempty"`               //行业类型*
		WorkName        string        `json:"workName" bson:"workName,omitempty"`               //职位名称*
		Age             string        `json:"age" bson:"age,omitempty"`                         //年龄要求
		Sex             string        `json:"sex" bson:"sex,omitempty"`                         //性别要求
		Experience      string        `json:"experience" bson:"experience,omitempty"`           //工作经验
		CultureLevel    string        `json:"cultureLevel" bson:"cultureLevel,omitempty"`       //文化水平
		MinSalary       string        `json:"minSalary" bson:"minSalary,omitempty"`             //最小薪资
		MaxSalary       string        `json:"maxSalary" bson:"maxSalary,omitempty"`             //最大薪资
		WorkAddress     string        `json:"workAddress" bson:"workAddress,omitempty"`         //工作地点
		RecruitNumber   string        `json:"recuritNumber" bson:"recuritNumber,omitempty"`     //招聘人数
		Welfare         string        `json:"welfare" bson:"welfare,omitempty"`                 //福利待遇
		Linkman         string        `json:"linkman" bson:"linkman,omitempty"`                 //联系人
		Phone           string        `json:"phone" bson:"phone,omitempty"`                     //联系电话
		PhoneImg        string        `json:"phoneImg" bson:"phoneImg,omitempty"`               //联系图片
		UnitName        string        `json:"unitName" bson:"unitName,omitempty"`               //单位名称
		CompanyInfo     string        `json:"companyInfo" bson:"companyInfo,omitempty"`         //企业信息
		ConditionPhotos string        `json:"conditionPhotos" bson:"conditionPhotos,omitempty"` //环境照片
		InfoFlag        string        `json:"info_flag" bson:"info_flag,omitempty"`
		PkID            string        `json:"id" bson:"id,omitempty"`
		// JobTitle        string        `json:"jobTitle" bson:"jobTitle,omitempty"`               //工作标题*
		// Jocation        string        `json:"location" bson:"location,omitempty"`               //      *
	}
	// 请求体结构体
	QueryParam struct {
		PageSize    int                    `json:"pageSize"`    //分页大小
		CurrentPage int                    `json:"currentPage"` //当前页码
		Conf        map[string]interface{} `json:"conf"`        //查询参数
	}
)

//获取表建立的链接
var conn = utils.GetConn("recruit")
