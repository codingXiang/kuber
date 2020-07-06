package ingress

import (
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Service interface {
	List(ns string, opts v1.ListOptions) (*v1beta1.IngressList, error)
	Get(ns string, name string, opts v1.GetOptions) (*v1beta1.Ingress, error)
}
