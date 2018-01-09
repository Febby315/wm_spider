package updownload

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"

	"../../utils"
	"github.com/julienschmidt/httprouter"
)

var mkURL = utils.GetStringValue("uploadfile", "savepath")

// var smallMkURL = utils.GetStringValue("uploadfile", "smallsavepath")
// var smallimglength = utils.GetIntValue("uploadfile", "smallimglength")
// var smallimghight = utils.GetIntValue("uploadfile", "smallimghight")

//下载文件
func downFile(fUrl *url.URL, outpath string) (err error) {
	if !fUrl.IsAbs() || len(outpath) <= 0 {
		return errors.New("不支持非绝对URL或保存路径为空")
	}
	//访问远程文件
	if res, err := http.Get(fUrl.String()); err != nil || res.StatusCode != 200 {
		return errors.New("远程文件访问失败")
	} else {
		defer res.Body.Close()
		//读取远程文件
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.New("远程文件读取失败")
		}
		//创建输出目录
		if err = os.MkdirAll(path.Dir(outpath), os.ModePerm); err != nil {
			return errors.New("输出目录创建失败")
		}
		//创建输出文件
		outfile, err := os.Create(outpath)
		if err != nil {
			return errors.New("输出文件创建失败")
		}
		defer outfile.Close()
		//复制文件
		if _, err := io.Copy(outfile, bytes.NewReader(body)); err != nil {
			return errors.New("远程文件复制失败")
		}
		//无异常返回
		fmt.Printf("输出文件:%s\n", outpath)
	}
	return
}

//路由:下载文件
func Download(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	utils.PrintReqLog(r)
	param := Param{}
	json.NewDecoder(r.Body).Decode(&param)
	if fUrl, err := url.Parse(param.ImgSrc); err != nil {
		jsonStr, _ := json.Marshal(map[string]interface{}{"state": 0, "msg": "文件访问路径有误"})
		w.Write(jsonStr)
		return
	} else {
		fmt.Printf("文件下载中:%s\n", param.ImgSrc)
		fPath := path.Join(mkURL, param.ImgURL, param.ImgName)
		if err := downFile(fUrl, fPath); err != nil { //下载图片
			fmt.Print(err)
			jsonStr, _ := json.Marshal(map[string]interface{}{"state": 99, "msg": "文件下载失败"})
			w.Write(jsonStr)
			return
		}
		// else {
		// 	//获取文件后缀
		// 	suffix := strings.ToUpper(path.Ext(fPath))
		// 	//缩略图保存路径
		// 	smallPath := path.Join(smallMkURL, param.ImgURL, param.ImgName)
		// 	//读取需要创建缩略图的后缀
		// 	thumbnailSuffix, sep := utils.GetStringValue("uploadfile", "thumbnailSuffix"), "|"
		// 	if n := strings.Index(thumbnailSuffix, sep+suffix+sep); n >= 0 {
		// 		if err := CreateThumbnail(fPath, smallPath); err != nil {
		// 			jsonStr, _ := json.Marshal(map[string]interface{}{"state": 99, "msg": "缩略图生成失败"})
		// 			w.Write(jsonStr)
		// 			return
		// 		}
		// 	}
		// }
	}
	jsonStr, _ := json.Marshal(map[string]interface{}{"state": 0, "msg": "文件下载成功"})
	w.Write(jsonStr)
}
