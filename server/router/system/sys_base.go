package system

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (a *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("hello", baseApi.Hello)
		BaseRouter.GET("health", baseApi.HealthCheck)
		BaseRouter.POST("saveAliyunSecret", baseApi.SaveAliyunSecret)
		BaseRouter.GET("getAliyunSecret", baseApi.GetAliyunSecret)
	}
}
