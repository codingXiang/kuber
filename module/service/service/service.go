package service

import (
	"github.com/codingXiang/kuber/client"
	"github.com/codingXiang/kuber/module/ingress"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IngressService struct {
	repo k8s.KubernetesClientInterface
}

func NewIngressService(repo k8s.KubernetesClientInterface) ingress.Service {
	return &IngressService{
		repo: repo,
	}
}

func (s *IngressService) List(ns string, opts v1.ListOptions) (*v1beta1.IngressList, error) {
	return s.repo.GetIngress(ns).List(opts)
}

func (s *IngressService) Get(ns string, name string, opts v1.GetOptions) (*v1beta1.Ingress, error) {
	return s.repo.GetIngress(ns).Get(name, opts)
}
