package system

import "github.com/gwyy/img-cloud-update/server/service"

type ApiGroup struct {
	BaseApi
}

var (
	systemService = service.ServiceGroupApp.SystemServiceGroup
)
