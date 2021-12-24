package initialize

import (
	"github.com/betterego/go-better-admin/server/global"
	"github.com/betterego/go-better-admin/server/middleware"
	"github.com/betterego/go-better-admin/server/router"
	"github.com/gin-gonic/gin"
)

func initRouters() {
	r := gin.Default()
	// 跨域
	r.Use(middleware.Cors())
	global.ROUTER = r
	systemRouter := router.RootRouters.SystemRouter
	systemGroup := global.ROUTER.Group("system")
	{
		systemRouter.InitRouter(systemGroup)

	}
}