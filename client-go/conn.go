package client_go

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)


var kubeconfig *string

func init(){

	config,err  := clientcmd.BuildConfigFromFlags("10.23.5.224","./config")
	if err!=nil{
		fmt.Println("config failed")
	}
	clientconn,err := kubernetes.NewForConfig(config)
	if err!=nil{
		fmt.Println("clientconn failed")
	}

}
