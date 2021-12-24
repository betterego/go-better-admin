package system

import (
	v1 "github.com/betterego/go-better-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type InitializeRouter struct {
}

func (i *InitializeRouter) InitRouter(router *gin.RouterGroup)  {
	initRouter := router.Group("init")
	initApi := v1.RootGroups.SystemGroup.InitApi
	{
		initRouter.POST("initdb",initApi.InitDB)
	}
}