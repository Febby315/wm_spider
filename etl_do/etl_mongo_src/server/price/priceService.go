package price

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../../utils"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// 列表查询
func List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w = utils.EnableXDA(w, r) //跨域请求及日志
	pageSize, _ := strconv.Atoi(utils.GetStringValue("database", "pageSize"))
	qparam, _param := QueryParam{PageSize: pageSize, CurrentPage: 1}, bson.M{} //初始化查询参数
	json.NewDecoder(r.Body).Decode(&qparam)
	if qparam.Conf["_id"] != nil && qparam.Conf["_id"].(string) != "" { //有_id则取大于_id记录
		_param["_id"] = bson.M{"$gt": bson.ObjectIdHex(qparam.Conf["_id"].(string))}
	}
	if qparam.Conf["info_flag"] != nil { //有标记按照条件查询
		if strings.ToLower(qparam.Conf["info_flag"].(string)) != "_all" {
			infoflags := strings.Split(qparam.Conf["info_flag"].(string), ",")
			_param["info_flag"] = bson.M{"$in": infoflags}
		}
	}
	log.Println("查询参数", _param)
	count, resultdata := 0, []Price{}
	count, _ = conn.Find(_param).Count()
	if err := conn.Find(_param).Sort("_id").Limit(qparam.PageSize).All(&resultdata); err != nil {
		log.Println("数据查询失败", err.Error())
	}
	result := bson.M{"count": count, "data": resultdata}
	uj, _ := json.Marshal(result)
	w.Write(uj)

}
