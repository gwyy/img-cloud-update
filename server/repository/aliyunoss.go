package repository

import (
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/gwyy/img-cloud-update/server/global"
	"github.com/gwyy/img-cloud-update/server/model/api/system/request"
	"github.com/gwyy/img-cloud-update/server/pkg/bbolt_manager"
)

/**
* @description: 获取阿里云oss配置
* @param {context.Context} ctx
* @return {request.SaveAliyunSecret, error}
 */
func GetAliyunOssConfig() (request.SaveAliyunSecret, error) {
	secret := request.SaveAliyunSecret{}
	region, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "region")
	if err != nil {
		return request.SaveAliyunSecret{}, err
	}
	bucket, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "bucket")
	if err != nil {
		return request.SaveAliyunSecret{}, err
	}
	accessKeyId, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "access_key_id")
	if err != nil {
		return request.SaveAliyunSecret{}, err
	}
	accessKeySecret, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "access_key_secret")
	if err != nil {
		return request.SaveAliyunSecret{}, err
	}
	secret.Region = region
	secret.Bucket = bucket
	secret.AccessKeyId = accessKeyId
	secret.AccessKeySecret = accessKeySecret
	return secret, nil
}

func InitAliyunOss() (*oss.Client, error) {
	//判断是否有accesskeyId 或者 secret
	secret, err := GetAliyunOssConfig()
	if err != nil || secret.AccessKeyId == "" || secret.AccessKeySecret == "" || secret.Region == "" || secret.Bucket == "" {
		return nil, err
	}

	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider(secret.AccessKeyId, secret.AccessKeySecret)).
		WithRegion(secret.Region)

	// 创建OSS客户端
	return oss.NewClient(cfg), nil
}
