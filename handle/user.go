package handle

import (
	"github.com/gin-gonic/gin"
	"k8s/common"
	"k8s/db"
)

func Userlogup(c *gin.Context) {
	name := c.PostForm("username")
	passwd := c.PostForm("passwd")

	var res common.ResponseData
	res.Msg = db.UserSignup(name,passwd)
	c.JSON(200,res)
}