package service

import (
	"github.com/codingXiang/kuber/client"
	"github.com/codingXiang/kuber/module/deployment"
	"k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentService struct {
	repo k8s.KubernetesClientInterface
}

func NewDeploymentService(repo k8s.KubernetesClientInterface) deployment.Service {
	return &DeploymentService{
		repo: repo,
	}
}

func (s *DeploymentService) List(ns string, opts v1.ListOptions) (*v1beta1.DeploymentList, error) {
	return s.repo.GetDeployment(ns).List(opts)
}

func (s *DeploymentService) Get(ns string, name string, opts v1.GetOptions) (*v1beta1.Deployment, error) {
	return s.repo.GetDeployment(ns).Get(name, opts)
}
