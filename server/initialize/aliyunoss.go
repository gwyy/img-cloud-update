package initialize

import (
	"github.com/gwyy/img-cloud-update/server/global"
	"github.com/gwyy/img-cloud-update/server/model/api/system/request"
	"github.com/gwyy/img-cloud-update/server/pkg/aliyunoss"
)

/**
* @description: 获取阿里云oss配置
* @param {context.Context} ctx
* @return {request.SaveAliyunSecret, error}
 */
func GetAliyunOssConfig() (request.SaveAliyunSecret, error) {
	secret, err := aliyunoss.GetAliyunOssConfig()
	if err != nil {
		return request.SaveAliyunSecret{}, err
	}
	return secret, nil
}

/**
* @description: 初始化阿里云oss
* @param {context.Context} ctx
* @return {error}
 */
func InitAliyunOss() {
	client, err := aliyunoss.InitAliyunOss()
	if err != nil {
		global.Log.Error("阿里云oss配置为空或错误", err)
		return
	}
	global.Oss = client
	global.Log.Info("阿里云oss初始化成功")
}
