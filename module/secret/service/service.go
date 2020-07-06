package service

import (
	"github.com/codingXiang/kuber/client"
	"github.com/codingXiang/kuber/module/service"
	"k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SvcService struct {
	repo k8s.KubernetesClientInterface
}

func NewSvcService(repo k8s.KubernetesClientInterface) service.Service {
	return &SvcService{
		repo: repo,
	}
}

func (s *SvcService) List(ns string, opts v12.ListOptions) (*v1.ServiceList, error) {
	return s.repo.GetService(ns).List(opts)
}

func (s *SvcService) Get(ns string, name string, opts v12.GetOptions) (*v1.Service, error) {
	return s.repo.GetService(ns).Get(name, opts)
}
