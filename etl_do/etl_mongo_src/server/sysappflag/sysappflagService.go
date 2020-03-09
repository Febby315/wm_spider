package sysappflag

import (
	"encoding/json"
	"log"
	"net/http"

	"../../utils"
	"github.com/julienschmidt/httprouter"
)

// 查询一条记录
func FindOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w = utils.EnableXDA(w, r) //跨域请求及日志
	qparam, result := QueryParam{}, Sysappflag{}
	json.NewDecoder(r.Body).Decode(&qparam)
	if err := conn.Find(&qparam.Conf).One(&result); err != nil {
		log.Println("数据查询失败", err.Error())
	}
	uj, _ := json.Marshal(result)
	w.Write(uj)
}
