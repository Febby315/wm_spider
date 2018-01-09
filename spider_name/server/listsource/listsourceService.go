package listsource

import (
	"encoding/json"
	"net/http"

	"../../utils"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

//List 列表
func List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	utils.PrintReqLog(r)

	qparam := QueryParam{PageSize: utils.GetIntValue("database", "pageSize"), CurrentPage: 1} //获取默认分页参数
	resultdata := []ListSource{}                                                              //结果集
	json.NewDecoder(r.Body).Decode(&qparam)                                                   //接收来自请求的分页参数
	total, _ := conn.Find(&qparam.Conf).Count()
	sort, page, size := utils.GetStringValue("database", "sortCol"), qparam.CurrentPage-1, qparam.PageSize
	if err := conn.Find(&qparam.Conf).Sort(sort).Skip(page * size).Limit(size).All(&resultdata); err != nil {
		panic(err)
	}
	result := map[string]interface{}{"total": total, "page": qparam.CurrentPage, "size": qparam.PageSize, "data": resultdata}
	jsonStr, _ := json.Marshal(result)
	w.Write(jsonStr)
	return
}

//Add 新增
func Add(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	utils.PrintReqLog(r)

	editinst := ListSource{}
	json.NewDecoder(r.Body).Decode(&editinst)
	editinst.ID = bson.NewObjectId()
	editinst.Version = 0
	//拥有没有处理或者没有处理完的记录暂不新增
	result, query := ListSource{}, bson.M{"status": bson.M{"$in": []int{1, 0}}, "list_url": editinst.ListURL}
	if err := conn.Find(&query).One(&result); err == nil || result.ListURL != "" {
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

//Edit 修改
func Edit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	utils.PrintReqLog(r)

	editinst := ListSource{}
	json.NewDecoder(r.Body).Decode(&editinst)
	param := bson.M{"_id": editinst.ID, "version": editinst.Version}
	editinst.Version++
	if err := conn.Update(param, bson.M{"$set": editinst}); err != nil {
		jsonStr, _ := json.Marshal(map[string]interface{}{"state": "fail", "msg": "修改记录失败:未找到符合条件的记录,请检查id和version"})
		w.Write(jsonStr)
		return
	}
	// if err := conn.FindId(editinst.ID).One(&editinst); err != nil {
	// 	jsonStr, _ := json.Marshal(map[string]interface{}{"state": "fail", "msg": "服务器错误:请联系管理员"})
	// 	w.Write(jsonStr)
	// 	return
	// }
	jsonStr, _ := json.Marshal(map[string]interface{}{"state": "success", "msg": "修改记录成功", "data": editinst})
	//jsonStr, _ := json.Marshal(editinst)//直接返回记录
	w.Write(jsonStr)
	return
}
