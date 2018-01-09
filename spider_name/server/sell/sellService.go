package sell

import (
	"encoding/json"
	"net/http"

	"../../utils"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

//Add 新增
func Add(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	utils.PrintReqLog(r)

	editinst := Sell{}
	json.NewDecoder(r.Body).Decode(&editinst)
	editinst.ID = bson.NewObjectId()
	editinst.TransactionId = uuid.NewV4().String()
	editinst.Version = 0
	result, query := []Sell{}, map[string]interface{}{"contentUrl": editinst.ContentUrl}
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
