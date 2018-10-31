package utils

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

//GetDB 获取数据库对象() Database
func GetDB() *mgo.Database {
	dbURL, dbName := GetStringValue("database", "dbURL"), GetStringValue("database", "dbName")
	if session, err := mgo.Dial(dbURL); err != nil {
		log.Panicln("链接数据库失败", err.Error())
	} else {
		session.SetMode(mgo.Monotonic, true)
		return session.DB(dbName)
	}
	return nil
}

//GetConn 获取数据库表通道(tableName) Collection
func GetConn(tableName string) *mgo.Collection {
	return GetDB().C(tableName)
}
