package initialize

import (
	"github.com/betterego/go-better-admin/server/global"
)

func init() {
	global.VIPER = Viper()
	global.LOG = Zap()
	global.DB = initMysql()
	initRouters()
}