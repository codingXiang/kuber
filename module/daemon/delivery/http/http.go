package http

import (
	cx "github.com/codingXiang/cxgateway/delivery"
	"github.com/codingXiang/cxgateway/pkg/e"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/kuber/module/deployment"
	"github.com/codingXiang/kuber/module/deployment/delivery"
	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type deploymentHttpHandler struct {
	gateway cx.HttpHandler
	svc     deployment.Service
}

func NewDeploymentHttpHandler(gateway cx.HttpHandler, svc deployment.Service) delivery.HttpHandler {
	handler := &deploymentHttpHandler{
		gateway: gateway,
		svc:     svc,
	}
	logger.Log.Info("Setup Deployment Handler")
	/*
		v1 版本的 deployment API
	  */
	logger.Log.Debug("use routing `/v1`")
	v1 := gateway.GetApiRoute().Group("/v1")
	api := v1.Group("/deployment")
	{
		api.GET("", e.Wrapper(handler.List))
		api.GET("/:namespace", e.Wrapper(handler.List))
		api.GET("/:namespace/:name", e.Wrapper(handler.Get))

	}
	return handler
}

func (g *deploymentHttpHandler) List(c *gin.Context) error {
	var ns = c.Params.ByName("namespace")
	var labelSelector string = ""
	if data, exist := c.GetQuery("labelSelector"); exist {
		labelSelector = data
	}
	if data, err := g.svc.List(ns, v1.ListOptions{
		LabelSelector: labelSelector,
	}); err == nil {
		c.JSON(e.StatusSuccess("list deployment", data))
		return nil
	} else {
		return e.UnknownError(err.Error())
	}
}

func (g *deploymentHttpHandler) Get(c *gin.Context) error {
	var (
		ns   = c.Params.ByName("namespace")
		name = c.Params.ByName("name")
	)

	if data, err := g.svc.Get(ns, name, v1.GetOptions{}); err == nil {
		c.JSON(e.StatusSuccess("get deployment", data))
		return nil
	} else {
		return e.UnknownError(err.Error())
	}
}
