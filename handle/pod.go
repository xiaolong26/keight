package handle

import (
	"context"
	"github.com/gin-gonic/gin"
	client_go "k8s/client-go"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Getpods(c *gin.Context){
	clientset := client_go.ClientConn()
	corev1 := clientset.CoreV1()
	cont := context.Background()
	pods,err :=corev1.Pods("default").List(cont,metav1.ListOptions{})
	if err!= nil{
		c.JSON(400,"pod错误")
	}
	c.JSON(200,pods.Items)
}