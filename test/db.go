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
	stmt, err := db.Prepare("select * from tbl_user where user_name=? limit 1")
	if err!=nil{
		fmt.Println("Failed to insert, err:" + err.Error())
	}
	defer stmt.Close()
	rows,err :=stmt.Query(username)
	columns,_ := rows.Columns()
	scanArgs := make([]interface{},len(columns))
	values := make([]interface{},len(columns))
	for j := range values{
		scanArgs[j] = &values[j]
	}
	record := make(map[string]interface{})
	records := make([]map[string]interface{}, 0)
	for rows.Next() {
		//将行数据保存到record字典
		err := rows.Scan(scanArgs...)
		if err!=nil{
			fmt.Println(err.Error())
		}
		for i, col := range values {
			if col != nil {
				record[columns[i]] = col
			}
		}
		records = append(records, record)
	}
	fmt.Println(values)
}