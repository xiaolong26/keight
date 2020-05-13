package handle

import (
	"github.com/gin-gonic/gin"
	"k8s/common"
	"k8s/resource"
)


func Getpods(c *gin.Context){
	res := HandlePods(c,common.Get)
	c.JSON(res.Code,res)
}

func HandlePods(c *gin.Context,action common.Action)common.ResponseData{
	r := resource.PodResource{
		Name : c.Request.FormValue("podname"),
		Namespace:c.Request.FormValue("namespace"),
	}
	res := common.ResponseData{}
	switch action {
	case common.Get:
		msg,err := r.Get()
		common.HandlerResponse(msg,err)
	}
	c.JSON(res.Code,res)
	return res
}