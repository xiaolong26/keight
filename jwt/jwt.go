package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"k8s/common"
)




/*var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's malformed token")
	TokenInvalid     = errors.New("couldn't handle this token")
)*/


type Cliams struct{
	Username string `json:"username"`
	Passwd	 string `json:"passwd"`
	jwt.StandardClaims
}
var jwtSecret=[]byte(common.Signing)


func Createtoken(cliam Cliams)(string,error){
	tokenClaims:=jwt.NewWithClaims(jwt.SigningMethodHS256,cliam)
	token,err:=tokenClaims.SignedString(jwtSecret)
	return token,err
}

func ParseToken(token string)(*Cliams,error)  {
	tokenClaims, err := jwt.ParseWithClaims(token, &Cliams{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims!=nil{
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims,ok:=tokenClaims.Claims.(*Cliams);ok&&tokenClaims.Valid{
			return claims,nil
		}
	}
	return nil,err
}