package handle

import (
	"github.com/gin-gonic/gin"
	"k8s/common"
	"k8s/common/handle"
	"k8s/resource"
)


func Getpods(c *gin.Context){
	res := HandlePods(c,common.Get)
	c.JSON(res.Code,res)
}

func Deletepods(c *gin.Context){
	res := HandlePods(c,common.Delete)
	c.JSON(res.Code,res)
}

func HandlePods(c *gin.Context,action common.Action)*common.ResponseData{
	g := handle.GenerateCommonParams(c)
	r := resource.PodResource{
		Resource: g,
		Name : c.Request.FormValue("podname"),
	}
	res := &common.ResponseData{}
	switch action {
	case common.Get:
		msg,err := r.Get()
		res = common.HandlerResponse(msg,err)
	case common.Delete:
		msg  := "delete success"
		err := r.Delete()
		if err!=nil{
			msg = "delete failed"
		}
		res = common.HandlerResponse(msg,err)
	}
	return res
}

func Test(c *gin.Context){
	test := handle.GenerateCommonParams(c)
	c.JSON(200,test)
}