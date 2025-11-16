package system

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/gwyy/img-cloud-update/server/global"
	"github.com/gwyy/img-cloud-update/server/model/api/system/request"
	"github.com/gwyy/img-cloud-update/server/model/common/response"
	"github.com/gwyy/img-cloud-update/server/pkg/aliyunoss"
	"github.com/gwyy/img-cloud-update/server/pkg/bbolt_manager"
)

type BaseService struct {
}

/**
* @description: 创建登录密码
* @param {context.Context} ctx
* @param {request.Login} req
* @return {error}
 */
func (s *BaseService) Install(ctx context.Context, req *request.Login) error {
	sysPassword, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "sys_password")
	if err != nil {
		return err
	}
	if sysPassword != "" {
		return errors.New("已注册！")
	}
	err = bbolt_manager.Save(global.Conf.BboltDB.TableName, "sys_password", req.Password)
	if err != nil {
		return err
	}
	return nil
}

/**
* @description: 登录
* @param {context.Context} ctx
* @param {request.Login} req
* @return {error}
 */
func (s *BaseService) Login(ctx context.Context, req *request.Login) error {
	sysPassword, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "sys_password")
	if err != nil {
		return err
	}
	if req.Password != sysPassword {
		return errors.New("密码错误")
	}
	return nil
}

/**
* @description: 获取阿里云oss配置
* @param {context.Context} ctx
* @return {request.SaveAliyunSecret, error}
 */
func (s *BaseService) GetAliyunSecret(ctx context.Context) (request.SaveAliyunSecret, error) {
	secret, err := aliyunoss.GetAliyunOssConfig()
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
	err = bbolt_manager.Save(global.Conf.BboltDB.TableName, "host", req.Host)
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
	err = bbolt_manager.Save(global.Conf.BboltDB.TableName, "default_path", req.DefaultPath)
	if err != nil {
		return err
	}
	err = bbolt_manager.Save(global.Conf.BboltDB.TableName, "sys_password", req.SysPassword)
	if err != nil {
		return err
	}
	err = bbolt_manager.Save(global.Conf.BboltDB.TableName, "url_mode", req.UrlMode)
	if err != nil {
		return err
	}
	err = bbolt_manager.Save(global.Conf.BboltDB.TableName, "copy_img_format", req.CopyImgFormat)
	if err != nil {
		return err
	}
	// 重新初始化阿里云oss
	client, err := aliyunoss.InitAliyunOss()
	if err != nil {
		global.Log.Error("阿里云oss配置为空或错误", err)
		return err
	}
	global.Oss = client
	global.Log.Info("重新初始化阿里云oss成功")
	return nil
}

/**
* @description: 上传文件
* @param {context.Context} ctx
* @param {*multipart.FileHeader} file
* @return {response.FileDownload, error}
 */
func (s *BaseService) UploadFile(ctx context.Context, path string, file *multipart.FileHeader) (response.FileDownload, error) {
	//设置bucketName
	bucketName := aliyunoss.GetAliyunOssBucket()
	if bucketName == "" {
		return response.FileDownload{}, errors.New("bucketName为空")
	}
	//获取host
	host := aliyunoss.GetAliyunOssHost()
	if host == "" {
		return response.FileDownload{}, errors.New("host为空")
	}

	//设置objectName  判断保存模式
	urlMode, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "url_mode")
	if err != nil {
		return response.FileDownload{}, err
	}
	//如果是默认保存路径
	objectName := ""
	if urlMode == "keep" {
		objectName = path + file.Filename
	} else { //如果是年月/时间戳（例子：202511/1731031200.png）
		//获取中间路径
		timestampStr := fmt.Sprintf("%d", time.Now().Unix())
		midPath := time.Now().Format("200601") + "/" + timestampStr
		fileNameExt := ""
		//获取文件后缀
		ext := filepath.Ext(file.Filename)
		if ext != "" {
			fileNameExt = strings.ToLower(ext)
		}
		//最终保存文件名
		objectName = path + midPath + fileNameExt
	}

	//上传文件
	objectName, err = aliyunoss.UploadFile(ctx, bucketName, objectName, file)
	if err != nil {
		return response.FileDownload{}, err
	}
	//返回文件下载地址
	downloadUrl := host + objectName
	return response.FileDownload{
		Name: objectName,
		Url:  downloadUrl,
	}, nil
}

/**
* @description: 删除文件
* @param {context.Context} ctx
* @param {string} name
* @return {error}
 */
func (s *BaseService) DeleteFile(ctx context.Context, objectName string) error {
	//设置bucketName
	bucketName := aliyunoss.GetAliyunOssBucket()
	if bucketName == "" {
		return errors.New("bucketName为空")
	}
	err := aliyunoss.DeleteFile(ctx, bucketName, objectName)
	if err != nil {
		return err
	}
	return nil
}
