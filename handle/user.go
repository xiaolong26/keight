package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s/common"
	"k8s/db"
	"k8s/jwt"
	"net/http"
	jwtAuth "github.com/dgrijalva/jwt-go"
	"time"
)

func Userlogup(c *gin.Context) {
	name := c.Request.FormValue("username")
	passwd := c.Request.FormValue("passwd")

	var res common.ResponseData
	if len(name)<3 || len(passwd) <5{
		c.JSON(http.StatusOK,gin.H{
			"msg":len(passwd)+len(name),
		})
		return
	}

	res.Msg = db.UserSignup(name,passwd)
	c.JSON(200,res)
}

func Userdologup(c *gin.Context){
	c.Redirect(http.StatusFound,"http://"+c.Request.Host+"/static/view/signup.html")
}

func Userdologin(c *gin.Context){
	c.Redirect(http.StatusFound,"http://"+c.Request.Host+"/static/view/sigin.html")
}

func Userlogin(c *gin.Context){
	name := c.Request.FormValue("username")
	passwd := c.Request.FormValue("passwd")
	pwdcheck := db.UserSignin(name,passwd)
	if !pwdcheck{
		res := common.ResponseData{
			Data: "",
			Code: 200,
			Msg:  pwdcheck,
		}
		c.JSON(200,res)
		return
	}
	cliams := jwt.Cliams{
		Username:       name,
		Passwd:         passwd,
		StandardClaims: jwtAuth.StandardClaims{
			ExpiresAt: time.Now().Add(common.ExpiresAt).Unix(),
		},
	}
	token,err := jwt.Createtoken(cliams)
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	res := common.ResponseData{
		Data: "返回token",
		Code: 0,
		Msg:  token,
	}
	c.JSON(200,res)
}

func CheckToken(c *gin.Context){
	token := c.Request.FormValue("token")
	claim,err := jwt.ParseToken(token)
	if err!=nil{
		fmt.Println(err)
	}
	name := claim.Username
	passwd := claim.Passwd
	pwdcheck := db.UserSignin(name,passwd)
	if !pwdcheck{
		c.JSON(200,false)
		return
	}
	c.JSON(200,claim)
}