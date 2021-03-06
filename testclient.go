package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//apiv1 "k8s.io/api/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	//      var kubeconfig *string
	// 配置 k8s 集群 kubeconfig 配置文件，默认位置 $HOME/.kube/config
	//      kubeconfig = "/root/.kube/config"
	//      kubeconfig = flag.String("kubeconfig", filepath.Join(home, "go", "src", "go-k8s", "config"), "(optional) absolute path to the kubeconfig file")
	//      flag.Parse()

	config,err := clientcmd.BuildConfigFromFlags("","./config")
	if err != nil {
		fmt.Println(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Connect kubernetes Successful！！！")

	corev1 := clientset.CoreV1()
	cont := context.Background()
	pods,err :=corev1.Pods("default").List(cont,metav1.ListOptions{})
	fmt.Printf("\nThere are %d pods in namespaces default", len(pods.Items))
}
