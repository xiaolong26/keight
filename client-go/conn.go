package client_go

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)


var clientset *kubernetes.Clientset

func init(){

	config,err := clientcmd.BuildConfigFromFlags("","./config")
	if err != nil {
		fmt.Println(err.Error())
	}
	clientset,err = kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func ClientConn() *kubernetes.Clientset{
	return clientset
}