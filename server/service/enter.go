package service

import "github.com/gwyy/img-cloud-update/server/service/system"

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup system.SystemServiceGroup
}
