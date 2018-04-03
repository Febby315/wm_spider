package updownload

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"../../utils"
	"github.com/julienschmidt/httprouter"
)

var mkURL = utils.GetStringValue("uploadfile", "savepath")

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
		log.Printf("输出文件:%s\n", outpath)
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
		log.Printf("文件下载中:%s\n", param.ImgSrc)
		fPath := path.Join(mkURL, param.ImgURL, param.ImgName)
		if err := downFile(fUrl, fPath); err != nil { //下载图片
			log.Println(err.Error())
			jsonStr, _ := json.Marshal(map[string]interface{}{"state": 99, "msg": err.Error(), "src": param.ImgSrc})
			w.Write(jsonStr)
			return
		}
	}
	jsonStr, _ := json.Marshal(map[string]interface{}{"state": 0, "msg": "文件下载成功", "src": param.ImgSrc})
	w.Write(jsonStr)
}
