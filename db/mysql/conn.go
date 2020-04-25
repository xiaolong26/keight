package mysql

import (
	"database/sql"
	"fmt"
	"os"
)

var db  *sql.DB

func init(){
	db,_ = sql.Open("mysql","xue:123456@tcp(127.0.0.1:3306)/keight?charset=utf8")
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
}

func DBConn() *sql.DB{
	return db
}
