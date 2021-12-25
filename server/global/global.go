package global

import (
	"github.com/betterego/go-better-admin/server/config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	SYSTEM config.System
	VIPER *viper.Viper
	ROUTER *gin.Engine
	LOG *zap.Logger
)
