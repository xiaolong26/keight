package  main

import (
	"context"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//v1beta1 "k8s.io/api/apps/v1beta1"
	//dly "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	"fmt"
)

func main() {
	config,err := clientcmd.BuildConfigFromFlags("","./config")

	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//创建命名空间//
	a := createns(clientset,"client-go")
	if a != 0 {
		fmt.Println("The namespace  create  is ok!")
	}else {
		fmt.Println("The namespace  create  is failed!")
	}

}


//创建命令空间函数//
func createns (clientsetx *kubernetes.Clientset, nsname string) int  {
	nc := new(apiv1.Namespace)
	nc.TypeMeta = metav1.TypeMeta{Kind:"Namespace",APIVersion:"v1"}
	nc.ObjectMeta = metav1.ObjectMeta{
		Name: nsname,
	}
	var cont context.Context
	nc.Spec = apiv1.NamespaceSpec{}
	opt := metav1.CreateOptions{
		TypeMeta:     metav1.TypeMeta{"namespace","apps/v1beta2"},
		DryRun:       "dryRun,omitempty",
		FieldManager: "",
	}
	_ ,err := clientsetx.CoreV1().Namespaces().Create(cont,nc,opt)
	if err != nil {
		return  0
	} else {
		return 1
	}
}
