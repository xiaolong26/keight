package common

type ResponseData struct {
	Data string 	`json:"data"`
	Code int 		`json:"code"`
	Msg	 interface{} 	`json:"msg"`
}

type User struct {
	Id		int		`json:"id"`
	Username string `json:"username"`
	passwd	string	`json:"passwd"`
}

const (
	Dbserver = "xue:123456@tcp(127.0.0.1:3306)/keight?charset=utf8"
)
