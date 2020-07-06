package delivery

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	List(c *gin.Context) error
	Get(c *gin.Context) error
}