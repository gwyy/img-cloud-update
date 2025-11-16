package initialize

import (
	"github.com/gwyy/img-cloud-update/server/global"
	"github.com/gwyy/img-cloud-update/server/pkg/bbolt_manager"
	"go.etcd.io/bbolt"
)

func InitDB() {
	dbPath := global.Conf.BboltDB.Path
	if dbPath == "" {
		dbPath = "data/config.db"
	}
	db, err := bbolt_manager.OpenDB(dbPath)
	if err != nil {
		global.Log.Error("初始化数据库失败", err)
		return
	}
	//初始化表
	initTable(db)
	// 设置全局数据库
	global.BboltDB = db
	global.Log.Info("初始化数据库成功")
}

func initTable(db *bbolt.DB) {
	tableName := global.Conf.BboltDB.TableName
	if tableName == "" {
		tableName = "config"
	}
	// 创建表
	err := bbolt_manager.CreateTable(db, tableName)
	if err != nil {
		global.Log.Error("初始化表失败", err)
	}
	global.Log.Info("初始化表成功")
}
