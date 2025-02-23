package global

import (
	"github.com/gwyy/img-cloud-update/server/config"
	"github.com/spf13/viper"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

var (
	Conf    config.Config
	Vp      *viper.Viper
	Log     *zap.SugaredLogger
	BboltDB *bbolt.DB
)
