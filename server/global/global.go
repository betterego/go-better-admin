package global

import (
	"github.com/betterego/go-better-admin/server/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	CONFIG config.Server
	VIPER *viper.Viper
)
