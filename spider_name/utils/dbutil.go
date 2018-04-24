package utils

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

//GetDB 获取数据库对象() Database
func GetDB() *mgo.Database {
	DbURL, DbName := GetStringValue("database", "dbURL"), GetStringValue("database", "dbName")
	session, err := mgo.Dial(DbURL)
	if err != nil {
		log.Panicln("链接数据库失败", err.Error())
	}
	session.SetMode(mgo.Monotonic, true)
	return session.DB(DbName)
}

//GetConn 获取数据库通道(tableName) Collection
func GetConn(tableName string) *mgo.Collection {
	return GetDB().C(tableName)
}
