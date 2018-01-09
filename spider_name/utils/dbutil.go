package utils

import (
	mgo "gopkg.in/mgo.v2"
)

//GetDB 获取数据库对象() Database
func GetDB() *mgo.Database {
	DbURL := GetStringValue("database", "dbURL")
	DbName := GetStringValue("database", "dbName")
	session, err := mgo.Dial(DbURL)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session.DB(DbName)
}

//GetConn 获取数据库通道(tableName) Collection
func GetConn(tableName string) *mgo.Collection {
	db := GetDB()
	return db.C(tableName)
}
