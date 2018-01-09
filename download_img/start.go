package main

import (
	"fmt"
	"net/http"

	"./server/downimg"
	"./server/updownload"
	"github.com/julienschmidt/httprouter"

	"./utils"
)

//Port a
var Port = utils.GetStringValue("api", "port")

// R Instantiate a new router
var R = httprouter.New()

func main() {
	fmt.Println("port:" + Port)
	http.ListenAndServe(":"+Port, R)
}

func init() {
	//downimg表操作
	R.POST(utils.GetStringValue("api", "downImg")+"/list", downimg.List)
	R.POST(utils.GetStringValue("api", "downImg")+"/add", downimg.Add)
	R.POST(utils.GetStringValue("api", "downImg")+"/edit", downimg.Edit)
	R.POST(utils.GetStringValue("api", "downImg")+"/findOne", downimg.FindOne)

	//新版文件下载
	R.POST(utils.GetStringValue("api", "updownload")+"/download", updownload.Download)

	// 旧版文件下载
	// R.POST(utils.GetStringValue("api", "download"), download.Down)
	// R.POST(utils.GetStringValue("api", "download")+"/makesmallimg", download.MakeSmallImg)
}
