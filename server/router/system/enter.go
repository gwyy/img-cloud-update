package system

import (
	api "github.com/gwyy/img-cloud-update/server/api/v1"
)

type RouterGroup struct {
	BaseRouter
}

var (
	baseApi = api.ApiGroupApp.SystemApiGroup.BaseApi
)
