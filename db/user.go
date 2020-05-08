package db

import (
	"fmt"
	mydb "k8s/db/mysql"
)

func UserSignup(username string,passwd string)bool{
	db := mydb.DBConn()
	stmt, err := db.Prepare(
		"insert into tbl_user(`user_name`,`user_pwd`) values(?,?)")
	if err!=nil{
		fmt.Println("failed to insert, err:" + err.Error())
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(username,passwd)
	if err != nil {
		fmt.Println("Failed to insert, err:" + err.Error())
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}

func UserSignin(username string,passwd string)bool{
	db := mydb.DBConn()
	stmt,err := db.Prepare("select `user_pwd` from tbl_user where user_name=? limit 1")
	if err!= nil{
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()
	rows,err :=stmt.Query(username)
	if err!=nil{
		fmt.Println(err.Error())
	}else if rows==nil{
		fmt.Println("query is nil"+username)
	}
	parse := mydb.ParseRows(rows)
	if len(parse)>0 && string(parse[0]["user_pwd"].([]byte)) == passwd{
		return true
	}
	return false
}