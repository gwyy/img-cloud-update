package initialize

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/gwyy/img-cloud-update/server/global"
	"github.com/spf13/viper"
)

func InitConfig(file string) {
	v := viper.New()
	v.SetConfigFile(file)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Conf); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.Conf); err != nil {
		panic(err)
	}

	global.Vp = v // 将viper实例赋值给全局变量
}
