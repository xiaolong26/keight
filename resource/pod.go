package resource

import (
	"context"
	"k8s.io/api/core/v1"
	client_go "k8s/client-go"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodResource struct {
	Name	string
	Namespace	string
}

func (p *PodResource)Get() (*v1.Pod, error) {
	return client_go.Clientset.CoreV1().Pods(p.Namespace).Get(context.Background(),p.Name,metav1.GetOptions{})
}
