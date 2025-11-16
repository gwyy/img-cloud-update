package system

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gwyy/img-cloud-update/server/global"
	"github.com/gwyy/img-cloud-update/server/model/api/system/request"
	"github.com/gwyy/img-cloud-update/server/model/common/response"
	"github.com/gwyy/img-cloud-update/server/pkg/bbolt_manager"
)

type BaseApi struct{}

func (a *BaseApi) Hello(c *gin.Context) {
	response.OkWithMessageAndData("hello", "v1.api.system.hello!", c)
}

func (a *BaseApi) HealthCheck(c *gin.Context) {
	response.OkWithMessage("ok", c)
}

// @Tags Base
// @Summary 登录
// @Accept json
// @Produce json
// @Param request body request.Login true "登录"
// @Success 200 {object} response.Response
// @Router /system/base/login [post]
func (a *BaseApi) Login(c *gin.Context) {
	var login request.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		response.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	err := systemService.Login(c.Request.Context(), &login)
	if err != nil {
		response.FailWithMessage("登录失败："+err.Error(), c)
		return
	}
	response.OkWithMessage("登录成功！", c)
}

// @Tags Base
// @Summary 判断是否安装
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /system/base/isInstall [get]
func (a *BaseApi) IsInstall(c *gin.Context) {
	sysPassword, err := bbolt_manager.Get(global.Conf.BboltDB.TableName, "sys_password")
	if err != nil {
		response.FailWithMessage("获取失败："+err.Error(), c)
		return
	}
	if sysPassword != "" {
		response.OkWithData(map[string]string{
			"isInstall": "true",
		}, c)
	} else {
		response.OkWithData(map[string]string{
			"isInstall": "false",
		}, c)
	}
}

// @Tags Base
// @Summary 初始化
// @Accept json
// @Produce json
// @Param request body request.Login true "初始化"
// @Success 200 {object} response.Response
// @Router /system/base/install [post]
func (a *BaseApi) Install(c *gin.Context) {
	var Login request.Login
	if err := c.ShouldBindJSON(&Login); err != nil {
		response.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	err := systemService.Install(c.Request.Context(), &Login)
	if err != nil {
		response.FailWithMessage("初始化失败："+err.Error(), c)
		return
	}
	response.OkWithMessage("初始化成功！", c)
}

// @Tags Base
// @Summary 保存阿里云密钥
// @Accept json
// @Produce json
// @Param request body request.SaveAliyunSecret true "阿里云密钥"
// @Success 200 {object} response.Response
// @Router /system/base/saveAliyunSecret [post]
func (a *BaseApi) SaveAliyunSecret(c *gin.Context) {
	var saveAliyunSecret request.SaveAliyunSecret
	if err := c.ShouldBindJSON(&saveAliyunSecret); err != nil {
		response.FailWithMessage("参数错误："+err.Error(), c)
		return
	}
	//host兼容处理，如果host不以/结尾，则加上/
	if !strings.HasSuffix(saveAliyunSecret.Host, "/") {
		saveAliyunSecret.Host += "/"
	}
	//默认路径兼容处理，如果默认路径不以/结尾，则加上/
	if saveAliyunSecret.DefaultPath != "" && !strings.HasSuffix(saveAliyunSecret.DefaultPath, "/") {
		saveAliyunSecret.DefaultPath += "/"
	}
	err := systemService.SaveAliyunSecret(c.Request.Context(), &saveAliyunSecret)
	if err != nil {
		response.FailWithMessage("更新失败："+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功！", c)
}

/**
* @description: 获取阿里云图片数据
* @param {context.Context} ctx
* @return {map[string]string, error}
 */
func (a *BaseApi) GetAliyunImgData(c *gin.Context) {
	secret, err := systemService.GetAliyunSecret(c.Request.Context())
	if err != nil {
		response.FailWithMessage("获取失败："+err.Error(), c)
		return
	}
	response.OkWithData(map[string]string{
		"path":          secret.DefaultPath,
		"copyImgFormat": secret.CopyImgFormat,
		"host":          secret.Host,
	}, c)
}

// @Tags Base
// @Summary 获取阿里云密钥
// @Accept json
// @Produce json
// @Success 200 {object} request.SaveAliyunSecret
// @Router /system/base/getAliyunSecret [get]
func (a *BaseApi) GetAliyunSecret(c *gin.Context) {
	secret, err := systemService.GetAliyunSecret(c.Request.Context())
	if err != nil {
		response.FailWithMessage("获取失败："+err.Error(), c)
		return
	}
	response.OkWithData(secret, c)
}

// @Tags Base
// @Summary 上传文件
// @Accept json
// @Produce json
// @Param file formData file true "文件"
// @Success 200 {object} response.Response
func (a *BaseApi) UploadImage(c *gin.Context) {
	//获取上传文件
	_, header, err := c.Request.FormFile("file")
	path := c.PostForm("path")
	global.Log.Info("path", path)
	global.Log.Info("header", header.Filename)
	if err != nil {
		response.FailWithMessage("获取上传文件失败："+err.Error(), c)
		return
	}
	file, err := systemService.UploadFile(c.Request.Context(), path, header)
	if err != nil {
		response.FailWithMessage("上传失败："+err.Error(), c)
		return
	}
	response.OkWithData(file, c)
}

// @Tags Base
// @Summary 删除图片
// @Accept json
// @Produce json
// @Param name formData string true "图片名称"
// @Success 200 {object} response.Response
// @Router /system/base/deleteImage [post]
func (a *BaseApi) DeleteImage(c *gin.Context) {
	var req struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("请求数据错误："+err.Error(), c)
		return
	}
	global.Log.Info("发送来的name", req.Name)
	err := systemService.DeleteFile(c.Request.Context(), req.Name)
	if err != nil {
		response.FailWithMessage("删除失败："+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功！", c)
}
