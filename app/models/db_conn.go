package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//var Conn _Conn
var Conn *gorm.DB


// Setup inject model with driver conn
func Setup() (err error) {
	dbpath := "/Users/admin/Todo.db"
	Conn,err= gorm.Open("sqlite3", dbpath)
	Conn.DB().SetMaxOpenConns(10)
	return Conn.DB().Ping()
}

//func SetOffsetAndLimit(conn *gorm.DB,userID int64,pageNum int){
//	if pageNum != 0 {
//		conn = conn.Offset((pageNum-1) * GetPageSize(userID)).Limit(GetPageSize(userID))
//	}else{
//		conn = conn.Limit(GetPageSize(userID))
//	}
//
//}