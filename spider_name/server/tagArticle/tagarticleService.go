package tagArticle

import (
	"encoding/json"
	"net/http"

	"../../utils"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

//Add 新增
func Add(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w = utils.EnableXDA(w, r) //跨域请求及日志
	editinst := TagArticle{}
	json.NewDecoder(r.Body).Decode(&editinst)
	editinst.ID = bson.NewObjectId()
	editinst.Version = 0
	result, query := []TagArticle{}, map[string]interface{}{"articleURL": editinst.ArticleURL}
	if err := conn.Find(&query).All(&result); err != nil || len(result) > 0 {
		jsonStr, _ := json.Marshal(map[string]interface{}{"state": "fail", "msg": "新增记录失败:记录已存在或验证失败"})
		w.Write(jsonStr)
		return
	}
	if err := conn.Insert(editinst); err != nil {
		jsonStr, _ := json.Marshal(map[string]interface{}{"state": "fail", "msg": "新增记录失败:请联系管理员"})
		w.Write(jsonStr)
		return
	}
	jsonStr, _ := json.Marshal(map[string]interface{}{"state": "success", "msg": "新增记录成功", "data": editinst})
	w.Write(jsonStr)
	return
}
