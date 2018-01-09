package downimg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../../utils"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
)

// List retrives an individual user resource
func List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.URL.Path)
	qparam := QueryParam{PageSize: utils.GetIntValue("database", "pageSize"), CurrentPage: 1} //make(map[string]interface{})

	json.NewDecoder(r.Body).Decode(&qparam)

	resultdata := []Downimg{}
	countNum, err := conn.Find(&qparam.Conf).Count()
	if err != nil {
		panic(err)
	}
	sortCol := utils.GetStringValue("database", "sortCol")

	err = conn.Find(&qparam.Conf).Sort(sortCol).Skip((qparam.CurrentPage - 1) * qparam.PageSize).Limit(qparam.PageSize).All(&resultdata)
	if err != nil {
		panic(err)
	}
	result := bson.M{"data": resultdata, "count": countNum}
	uj, _ := json.Marshal(result)
	w.Write(uj)
}

//Add 新增
func Add(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.URL.Path)
	editinst := Downimg{}
	json.NewDecoder(r.Body).Decode(&editinst)
	editinst.ID = bson.NewObjectId()
	editinst.Version = 0
	conn.Insert(editinst)
	uj, _ := json.Marshal(editinst)
	w.Write(uj)
}

//Edit 编辑
func Edit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.URL.Path)
	editinst := Downimg{}
	json.NewDecoder(r.Body).Decode(&editinst)

	_param := bson.M{"_id": editinst.ID, "version": editinst.Version}
	editinst.Version = editinst.Version + 1
	update := bson.M{"$set": editinst}

	if err := conn.Update(_param, update); err != nil {
		w.WriteHeader(404)
		return
	}

	if err := conn.FindId(editinst.ID).One(&editinst); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, _ := json.Marshal(editinst)
	w.Write(uj)
}

//FindOne 编辑
func FindOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.URL.Path)
	qparam := QueryParam{} //make(map[string]interface{})

	json.NewDecoder(r.Body).Decode(&qparam)
	result := Downimg{}
	conn.Find(&qparam.Conf).One(&result)

	uj, _ := json.Marshal(result)

	w.Write(uj)
}
