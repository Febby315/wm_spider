package main

import (
	"fmt"
	"net/http"

	"./server/article"           //文章
	"./server/buy"               //求购
	"./server/listdetailcleaned" //洗后文章
	"./server/listdetailsource"  //文章源
	"./server/listsource"        //列表源
	"./server/price"             //价格
	"./server/recruit"           //招聘
	"./server/sell"              //供应
	"./server/spiderconfig"      //脚本配置
	"./server/tagArticle"        //标签文章
	"./utils"                    //工具箱

	"github.com/julienschmidt/httprouter"
)

//Port is This Server's port
var Port = utils.GetStringValue("api", "port")

//R is rotuer
var R = httprouter.New()

func init() {
	//站点配置
	R.POST(utils.GetStringValue("api", "spiderconfig")+"/list", spiderconfig.List) //列表
	R.POST(utils.GetStringValue("api", "spiderconfig")+"/add", spiderconfig.Add)   //增加
	R.POST(utils.GetStringValue("api", "spiderconfig")+"/edit", spiderconfig.Edit) //修改
	//已清洗详情
	R.POST(utils.GetStringValue("api", "listdetailcleaned")+"/list", listdetailcleaned.List) //列表
	R.POST(utils.GetStringValue("api", "listdetailcleaned")+"/add", listdetailcleaned.Add)   //增加
	R.POST(utils.GetStringValue("api", "listdetailcleaned")+"/edit", listdetailcleaned.Edit) //修改
	//列表源
	R.POST(utils.GetStringValue("api", "listsource")+"/list", listsource.List) //列表
	R.POST(utils.GetStringValue("api", "listsource")+"/add", listsource.Add)   //增加
	R.POST(utils.GetStringValue("api", "listsource")+"/edit", listsource.Edit) //修改
	//详情源
	R.POST(utils.GetStringValue("api", "listdetailsource")+"/list", listdetailsource.List) //列表
	R.POST(utils.GetStringValue("api", "listdetailsource")+"/add", listdetailsource.Add)   //增加
	R.POST(utils.GetStringValue("api", "listdetailsource")+"/edit", listdetailsource.Edit) //修改
	//文章接口
	R.POST(utils.GetStringValue("api", "article")+"/list", article.List) //列表
	R.POST(utils.GetStringValue("api", "article")+"/add", article.Add)   //增加
	R.POST(utils.GetStringValue("api", "article")+"/edit", article.Edit) //修改
	//价格接口
	R.POST(utils.GetStringValue("api", "price")+"/add", price.Add) //增加
	//招聘接口
	R.POST(utils.GetStringValue("api", "recruit")+"/add", recruit.Add) //增加
	//供应
	R.POST(utils.GetStringValue("api", "sell")+"/add", sell.Add) //增加
	//求购
	R.POST(utils.GetStringValue("api", "buy")+"/add", buy.Add) //增加
	//标签文章
	R.POST(utils.GetStringValue("api", "tagArticle")+"/add", tagArticle.Add) //增加
}

//程序主入口
func main() {
	fmt.Println("port:" + Port)
	http.ListenAndServe(":"+Port, R)
}
