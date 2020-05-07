package mysql

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var db  *sql.DB

func init(){
	db,_ = sql.Open("mysql","root:123456@tcp(10.23.5.225:3306)/keight?charset=utf8")
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
}
func DBConn() *sql.DB{
	return db
}
