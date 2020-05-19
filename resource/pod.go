package resource

import (
	"context"
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	client_go "k8s/client-go"
	"k8s/common/handle"
)

type PodResource struct {
	Resource	*handle.Resource
	Name	string
}

func (p *PodResource)Get() (*v1.Pod, error) {
	return client_go.Clientset.CoreV1().Pods(p.Resource.Namespace).Get(context.Background(),p.Name,metav1.GetOptions{})
}

func (p *PodResource)Delete() (err error) {
	pod,_ := p.Get()
	if err = client_go.Clientset.CoreV1().Pods(p.Resource.Namespace).Delete(context.Background(),p.Name,metav1.DeleteOptions{});err !=nil{
		return
	}
	watcher,err := client_go.Clientset.CoreV1().Pods(p.Resource.Namespace).Watch(context.Background(),metav1.SingleObject(pod.ObjectMeta))
	fmt.Println(watcher)
	return
}

