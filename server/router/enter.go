package router

import "github.com/gwyy/img-cloud-update/server/router/system"

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System system.RouterGroup
}
