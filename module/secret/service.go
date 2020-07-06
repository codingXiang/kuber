package service

import (
	"k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Service interface {
	List(ns string, opts v12.ListOptions) (*v1.ServiceList, error)
	Get(ns string, name string, opts v12.GetOptions) (*v1.Service, error)
}
