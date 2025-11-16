package initialize

import (
	"fmt"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/gwyy/img-cloud-update/server/global"
	"github.com/gwyy/img-cloud-update/server/router"
)

func InitServer() {
	//初始化gin
	Router := gin.New()
	Router.Use(gin.Recovery())
	//设置gin模式
	gin.SetMode(getGinMode())
	if global.Conf.Mode == "dev" || global.Conf.Mode == "test" {
		Router.Use(gin.Logger())
	}
	global.Log.Info("gin启动模式为：", gin.Mode())

	//初始化路由
	Router.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"message": "start success"}) })
	systemRouter := router.RouterGroupApp.System
	RouterGroup := Router.Group(global.Conf.Gin.Prefix)
	systemRouter.InitBaseRouter(RouterGroup)

	// 获取端口
	address := fmt.Sprintf(":%d", global.Conf.Gin.Port)
	// 启动服务器
	s := endless.NewServer(address, Router)
	s.ReadHeaderTimeout = 10 * time.Minute
	s.WriteTimeout = 10 * time.Minute
	s.MaxHeaderBytes = 1 << 20
	global.Log.Info("server run success on ", address)

	fmt.Printf(`
	img-cloud-update 启动成功
	当前版本:v%s
	默认前端文件运行地址:http://127.0.0.1%s
`, global.Conf.Version, address)
	global.Log.Error(s.ListenAndServe().Error())
}

// 获取gin模式
func getGinMode() string {
	if global.Conf.Mode == "dev" {
		return gin.DebugMode
	} else if global.Conf.Mode == "test" {
		return gin.TestMode
	} else {
		return gin.ReleaseMode
	}
}
