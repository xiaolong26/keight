package router

import (
	"github.com/gin-gonic/gin"
	"k8s/handle"
	"k8s/jwt"
)

func SetRouter(r *gin.Engine) *gin.Engine{

	//r.Static("/static","./static")
	//r.GET("/pod",handle.Getpods)
	r.POST("/user/logup",handle.Userlogup)
	r.GET("/user/logup",handle.Userdologup)
	r.GET("/user/login",handle.Userdologin)
	r.POST("/usr/login",handle.Userlogin)
	r.POST("/usr/auth",handle.CheckToken)
	r.GET("/test",handle.Test)
	auth := r.Group("/a",jwt.JWY())
	{
		auth.Static("/static","./static")
		//auth.GET("/podlist",handle.Getpods)
		auth.GET("/pod",handle.Getpods)
		auth.GET("/pod/delete",handle.Deletepods)
	}
	return r
}
