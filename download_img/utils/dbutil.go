package utils

import (
	mgo "gopkg.in/mgo.v2"
)

var dbURL = GetStringValue("database", "dbURL")
var dbName = GetStringValue("database", "dbName")

// GetDB return a mongo database
func GetDB() *mgo.Database {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		panic(err)
	}
	return session.DB(dbName)
}

// GetConn return a mongo collection
func GetConn(tableName string) *mgo.Collection {
	return GetDB().C(tableName)
}
