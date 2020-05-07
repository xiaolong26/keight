package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db  *sql.DB

const (
	username = "xue"
	passwd = "xue"
)
func main(){
	db, _ = sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/keight?charset=utf8")

	db.SetMaxOpenConns(1000)
	stmt, err := db.Prepare("insert into tbl_user(`user_name`,`user_pwd`) values(?,?)")
	if err!=nil{
		fmt.Println("Failed to insert, err:" + err.Error())
	}
	defer stmt.Close()
	_,_ = stmt.Exec(username, passwd)
}