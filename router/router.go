package router

import (
	"github.com/gin-gonic/gin"
	"k8s/handle"
)

func SetRouter(r *gin.Engine) *gin.Engine{
	r.GET("/",handle.Test)
	return r
}
