package common

import (
	"net/http"
	"time"
)

type ResponseData struct {
	Data string 	`json:"data"`
	Code int 		`json:"code"`
	Msg	 interface{} 	`json:"msg"`
}

func HandlerResponse(result interface{},err error)*ResponseData{
	var res *ResponseData
	if err!=nil{
		res = &ResponseData{
			Data: "",
			Code: http.StatusInternalServerError,
			Msg:  result,
		}
	}else  {
		res = &ResponseData{
			Data: "",
			Code: http.StatusOK,
			Msg:  result,
		}
	}
	return res
}
type User struct {
	Id		int		`json:"id"`
	Username string `json:"username"`
	passwd	string	`json:"passwd"`
}

const (
	Dbserver = "xue:123456@tcp(127.0.0.1:3306)/keight?charset=utf8"
	Signing  = "keight"
	ExpiresAt      = 30 * 24 * time.Hour
)

type Action string

const  (
	Get   =  Action("get")
	Delete = Action("delete")
)
