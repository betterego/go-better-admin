package system

import (
	"github.com/betterego/go-better-admin/server/domain/system/request"
	"github.com/betterego/go-better-admin/server/global"
	"github.com/betterego/go-better-admin/server/util"
	"github.com/gin-gonic/gin"
)

type InitApi struct {
}

func (initApi *InitApi) InitDB(c *gin.Context)  {
	if global.DB != nil {
		util.FailWithMessage("数据库已初始化！",c)
		return
	}
	var config request.DBConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		util.FailWithMessage("参数校验不通过!",c)
		return
	}
	if err := initializeService.InitDb(&config); err != nil {
		util.FailWithMessage("数据库初始化失败！",c)
		return
	}
	if err := initializeService.WriteConfig(&config); err != nil {
		util.FailWithMessage("写入配置文件失败!",c)
		return
	}
	util.Ok(c)
}