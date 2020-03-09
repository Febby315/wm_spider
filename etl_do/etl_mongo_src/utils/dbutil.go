package utils

import (
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

// GetDB return a mongo session
func GetDB() *mgo.Database {
	DbURL := GetStringValue("database", "dbURL")
	DbName := GetStringValue("database", "dbName")
	// Connect to our local mongo
	s, err := mgo.Dial(DbURL)
	if err != nil {
		panic(err)
	}
	db := s.DB(DbName)
	return db
}

// GetConn return a mongo session
func GetConn(tableName string) *mgo.Collection {

	c := GetDB().C(tableName)
	return c
}

//允许跨域及访问日志
func EnableXDA(w http.ResponseWriter, r *http.Request) http.ResponseWriter {
	log.Println(r.Method, r.URL)
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")
	return w
}
