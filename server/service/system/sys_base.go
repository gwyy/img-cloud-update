package system

import (
	"context"

	"github.com/gwyy/img-cloud-update/server/global"
	"github.com/gwyy/img-cloud-update/server/model/api/system/request"
	"github.com/gwyy/img-cloud-update/server/pkg/bbolt_manager"
	"github.com/gwyy/img-cloud-update/server/repository"
)

type BaseService struct {
}

/**
* @description: 获取阿里云oss配置
* @param {context.Context} ctx
* @return {request.SaveAliyunSecret, error}
 */
func (s *BaseService) GetAliyunSecret(ctx context.Context) (request.SaveAliyunSecret, error) {
	secret, err := repository.GetAliyunOssConfig()
	if err != nil {
		return request.SaveAliyunSecret{}, err
	}
	return secret, nil
}

/**
* @description: 保存阿里云oss配置
* @param {context.Context} ctx
* @param {request.SaveAliyunSecret} req
* @return {error}
 */
func (s *BaseService) SaveAliyunSecret(ctx context.Context, req *request.SaveAliyunSecret) error {
	err := bbolt_manager.Save(global.Conf.BboltDB.TableName, "region", req.Region)
	if err != nil {
		return err
	}
	err = bbolt_manager.Save(global.Conf.BboltDB.TableName, "bucket", req.Bucket)
	if err != nil {
		return err
	}
	err = bbolt_manager.Save(global.Conf.BboltDB.TableName, "access_key_id", req.AccessKeyId)
	if err != nil {
		return err
	}
	err = bbolt_manager.Save(global.Conf.BboltDB.TableName, "access_key_secret", req.AccessKeySecret)
	if err != nil {
		return err
	}
	// 重新初始化阿里云oss
	client, err := repository.InitAliyunOss()
	if err != nil {
		global.Log.Error("阿里云oss配置为空或错误", err)
		return err
	}
	global.Oss = client
	global.Log.Info("重新初始化阿里云oss成功")
	return nil
}
