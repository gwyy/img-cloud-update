package system

import (
	"context"

	"github.com/gwyy/img-cloud-update/server/global"
	"github.com/gwyy/img-cloud-update/server/model/api/system/request"
	"github.com/gwyy/img-cloud-update/server/pkg/bbolt_manager"
)

type BaseService struct {
}

func (s *BaseService) GetAliyunSecret(ctx context.Context) (request.SaveAliyunSecret, error) {
	secret := request.SaveAliyunSecret{}
	accessKeyId, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "access_key_id")
	if err != nil {
		return request.SaveAliyunSecret{}, err
	}
	accessKeySecret, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "access_key_secret")
	if err != nil {
		return request.SaveAliyunSecret{}, err
	}
	secret.AccessKeyId = accessKeyId
	secret.AccessKeySecret = accessKeySecret
	return secret, nil
}

func (s *BaseService) SaveAliyunSecret(ctx context.Context, req *request.SaveAliyunSecret) error {
	err := bbolt_manager.Save(global.Conf.BboltDB.TableName, "access_key_id", req.AccessKeyId)
	if err != nil {
		return err
	}
	err = bbolt_manager.Save(global.Conf.BboltDB.TableName, "access_key_secret", req.AccessKeySecret)
	if err != nil {
		return err
	}
	return nil
}
