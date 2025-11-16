package main

import (
	"flag"

	"github.com/gwyy/img-cloud-update/server/global"
	"github.com/gwyy/img-cloud-update/server/initialize"
)

//go:generate go mod tidy
//go:generate go mod download

func main() {
	//获取配置文件
	var configFile string
	flag.StringVar(&configFile, "conf", "data/config.yaml", "config file path")
	flag.Parse()

	// 初始化配置
	initialize.InitConfig(configFile)

	// 初始化日志
	initialize.InitLogger()

	// 初始化数据库
	initialize.InitDB()
	if global.BboltDB != nil {
		defer global.BboltDB.Close()
	}
	// 初始化阿里云oss
	initialize.InitAliyunOss()
	// 初始化服务
	initialize.InitServer()
}
