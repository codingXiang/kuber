package http

import (
	cx "github.com/codingXiang/cxgateway/delivery"
	"github.com/codingXiang/cxgateway/pkg/e"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/kuber/module/deployment/delivery"
	"github.com/codingXiang/kuber/module/service"
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ServiceHttpHandler struct {
	gateway cx.HttpHandler
	svc     service.Service
}

func NewServiceHttpHandler(gateway cx.HttpHandler, svc service.Service) delivery.HttpHandler {
	handler := &ServiceHttpHandler{
		gateway: gateway,
		svc:     svc,
	}
	logger.Log.Info("Setup Service Handler")
	/*
		v1 版本的 service API
	  */
	logger.Log.Debug("use routing `/v1`")
	v1 := gateway.GetApiRoute().Group("/v1")
	deployment := v1.Group("/service")
	{
		deployment.GET("", e.Wrapper(handler.List))
		deployment.GET("/:namespace", e.Wrapper(handler.List))
		deployment.GET("/:namespace/:name", e.Wrapper(handler.Get))
	}
	return handler
}

func (g *ServiceHttpHandler) List(c *gin.Context) error {
	var ns = c.Params.ByName("namespace")

	if data, err := g.svc.List(ns, v1.ListOptions{}); err == nil {
		c.JSON(e.StatusSuccess("list service", data))
		return nil
	} else {
		return e.UnknownError(err.Error())
	}
}

func (g *ServiceHttpHandler) Get(c *gin.Context) error {
	var (
		ns   = c.Params.ByName("namespace")
		name = c.Params.ByName("name")
	)

	if data, err := g.svc.Get(ns, name, v1.GetOptions{}); err == nil {
		c.JSON(e.StatusSuccess("get service", data))
		return nil
	} else {
		return e.UnknownError(err.Error())
	}
}
