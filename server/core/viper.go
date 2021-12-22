package core

import (
	"fmt"
	"github.com/betterego/go-better-admin/server/config"
	"github.com/betterego/go-better-admin/server/global"
	"github.com/spf13/viper"
)

func init() {
	global.VIPER = Viper()
}

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigName(config.ConfigFileName)
	v.SetConfigType(config.ConfigFileType)
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("配置文件错误: %s \n", err))
	}
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}