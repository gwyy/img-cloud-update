package v1

import (
	"github.com/gwyy/img-cloud-update/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}
