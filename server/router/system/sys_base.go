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
		BaseRouter.POST("login", baseApi.Login)
		BaseRouter.GET("getAliyunImgData", baseApi.GetAliyunImgData)
		BaseRouter.POST("saveAliyunSecret", baseApi.SaveAliyunSecret)
		BaseRouter.POST("install", baseApi.Install)
		BaseRouter.GET("isInstall", baseApi.IsInstall)
		BaseRouter.GET("getAliyunSecret", baseApi.GetAliyunSecret)
		BaseRouter.POST("uploadImage", baseApi.UploadImage)
		BaseRouter.POST("deleteImage", baseApi.DeleteImage)
	}
}
