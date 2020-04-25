package main

import (
	"github.com/gin-gonic/gin"
	"k8s/router"
	)
func main(){
	r := gin.Default()
	router := router.SetRouter(r)
	router.Run("127.0.0.1:8080")
}
