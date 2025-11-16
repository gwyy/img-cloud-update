package aliyunoss

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/gwyy/img-cloud-update/server/global"
	"github.com/gwyy/img-cloud-update/server/model/api/system/request"
	"github.com/gwyy/img-cloud-update/server/pkg/bbolt_manager"
)

func GetAliyunOssBucket() string {
	bucket, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "bucket")
	if err != nil {
		return ""
	}
	return bucket
}

func GetAliyunOssHost() string {
	host, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "host")
	if err != nil {
		return ""
	}
	return host
}

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
	host, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "host")
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
	defaultPath, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "default_path")
	if err != nil {
		return request.SaveAliyunSecret{}, err
	}
	sysPassword, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "sys_password")
	if err != nil {
		return request.SaveAliyunSecret{}, err
	}
	urlMode, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "url_mode")
	if err != nil {
		return request.SaveAliyunSecret{}, err
	}
	copyImgFormat, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "copy_img_format")
	if err != nil {
		return request.SaveAliyunSecret{}, err
	}

	secret.Region = region
	secret.Bucket = bucket
	secret.Host = host
	secret.AccessKeyId = accessKeyId
	secret.AccessKeySecret = accessKeySecret
	secret.DefaultPath = defaultPath
	secret.SysPassword = sysPassword
	secret.UrlMode = urlMode
	secret.CopyImgFormat = copyImgFormat

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

/**
* @description: 上传文件
* @param {context.Context} ctx
* @param {*multipart.FileHeader} file
* @return {response.FileDownload, error}
 */
func UploadFile(ctx context.Context, bucketName string, objectName string, file *multipart.FileHeader) (string, error) {
	if global.Oss == nil {
		global.Log.Error("阿里云oss未初始化")
		return "", errors.New("阿里云oss未初始化")
	}
	if bucketName == "" {
		global.Log.Error("bucketName为空")
		return "", errors.New("bucketName为空")
	}
	if objectName == "" {
		global.Log.Error("objectName为空")
		return "", errors.New("objectName为空")
	}
	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		return "", errors.New("读取不到文件内容！")
	}
	defer f.Close() // 创建文件 defer 关闭
	// 创建上传对象的请求
	putRequest := &oss.PutObjectRequest{
		Bucket:       oss.Ptr(bucketName),      // 存储空间名称
		Key:          oss.Ptr(objectName),      // 对象名称
		StorageClass: oss.StorageClassStandard, // 指定对象的存储类型为标准存储
		Acl:          oss.ObjectACLPrivate,     // 指定对象的访问权限为私有访问
		Body:         f,
	}

	// 执行上传对象的请求
	_, err := global.Oss.PutObject(ctx, putRequest)
	if err != nil {
		global.Log.Errorf("上传文件失败 %v", err)
		return "", err
	}
	return objectName, nil
}

func DeleteFile(ctx context.Context, bucketName string, objectName string) error {
	if global.Oss == nil {
		global.Log.Error("阿里云oss未初始化")
		return errors.New("阿里云oss未初始化")
	}
	// 创建删除对象的请求
	request := &oss.DeleteObjectRequest{
		Bucket: oss.Ptr(bucketName), // 存储空间名称
		Key:    oss.Ptr(objectName), // 对象名称
	}
	global.Log.Info("删除文件", bucketName)
	global.Log.Info("删除文件", objectName)
	// 执行删除对象的操作并处理结果
	_, err := global.Oss.DeleteObject(ctx, request)
	if err != nil {
		global.Log.Errorf("删除文件失败 %v", err)
	}
	return nil
}
