package main

import (
	"log"
	"net/http"

	"./utils"
	"github.com/julienschmidt/httprouter"

	serviceBuyTransaction "./server/buyTransaction"
	serviceDownimg "./server/downimg"
	servicePrice "./server/price"
	serviceRecruit "./server/recruit"
	serviceSellTransaction "./server/sellTransaction"
	serviceSysappflag "./server/sysappflag"
	serviceSysdatapoint "./server/sysdatapoint"
	serviceTagArticle "./server/tagArticle"
)

//Port 端口
var Port = utils.GetStringValue("api", "port")

// R Instantiate a new router
var R = httprouter.New()

func main() {
	log.Printf("listen to localhost:%s\n", Port)
	http.ListenAndServe(":"+Port, R)
}
func init() {
	R.POST(utils.GetStringValue("api", "sysappflag")+"/findone", serviceSysappflag.FindOne)     //允许检索查询
	R.POST(utils.GetStringValue("api", "sysdatapoint")+"/findone", serviceSysdatapoint.FindOne) //同步进度查询
	R.POST(utils.GetStringValue("api", "sysdatapoint")+"/edit", serviceSysdatapoint.Edit)       //同步进度编辑
	R.POST(utils.GetStringValue("api", "tagArticle")+"/list", serviceTagArticle.List)           //文章查询接口
	R.POST(utils.GetStringValue("api", "sellTransaction")+"/list", serviceSellTransaction.List) //供应查询接口
	R.POST(utils.GetStringValue("api", "buyTransaction")+"/list", serviceBuyTransaction.List)   //求购查询接口
	R.POST(utils.GetStringValue("api", "price")+"/list", servicePrice.List)                     //价格查询接口
	R.POST(utils.GetStringValue("api", "recruit")+"/list", serviceRecruit.List)                 //招聘查询接口
	R.POST(utils.GetStringValue("api", "downImg")+"/list", serviceDownimg.List)                 //附件查询接口
}
