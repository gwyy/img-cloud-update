package system

import (
	"github.com/gin-gonic/gin"
	"github.com/gwyy/img-cloud-update/server/model/api/system/request"
	"github.com/gwyy/img-cloud-update/server/model/common/response"
)

type BaseApi struct{}

func (a *BaseApi) Hello(c *gin.Context) {
	response.OkWithMessageAndData("hello", "v1.api.system.hello!", c)
}

func (a *BaseApi) HealthCheck(c *gin.Context) {
	response.OkWithMessage("ok", c)
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
	err := systemService.SaveAliyunSecret(c.Request.Context(), &saveAliyunSecret)
	if err != nil {
		response.FailWithMessage("更新失败："+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功！", c)
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
