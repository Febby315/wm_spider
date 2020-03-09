package sysdatapoint

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"../../utils"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// 查询一条记录
func FindOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w = utils.EnableXDA(w, r) //跨域请求及日志
	qparam, resultdata := QueryParam{}, Sysdatapoint{}
	json.NewDecoder(r.Body).Decode(&qparam)
	if err := conn.Find(&qparam.Conf).One(&resultdata); err != nil {
		log.Println("数据查询失败", err.Error())
	}
	uj, _ := json.Marshal(resultdata)
	w.Write(uj)
}

//编辑&新增
func Edit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w = utils.EnableXDA(w, r) //跨域请求及日志
	editinst := Sysdatapoint{}
	json.NewDecoder(r.Body).Decode(&editinst)
	param := map[string]interface{}{
		"appid":     editinst.Appid,
		"passwd":    editinst.Passwd,
		"tablename": editinst.Tablename,
	}
	resultdata := Sysdatapoint{}
	log.Println("sysdatapoint编辑参数", editinst)
	if err := conn.Find(param).One(&resultdata); err != nil {
		editinst.ID = bson.NewObjectId()
		editinst.CreatedDate = time.Now().Format("2006-01-02 15:04:05")
		editinst.LastUpdateDate = time.Now().Format("2006-01-02 15:04:05")
		conn.Insert(editinst)
		uj, _ := json.Marshal(editinst)
		w.Write(uj)
	} else {
		_param := bson.M{"_id": resultdata.ID, "version": resultdata.Version}
		editinst.ID = resultdata.ID
		editinst.Version = resultdata.Version + 1
		editinst.LastUpdateDate = time.Now().Format("2006-01-02 15:04:05")
		if err := conn.Update(_param, bson.M{"$set": editinst}); err != nil {
			w.WriteHeader(404)
			return
		}
		uj, _ := json.Marshal(editinst)
		w.Write(uj)
	}
}
