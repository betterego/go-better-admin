package initialize

import (
	"fmt"
	"github.com/betterego/go-better-admin/server/global"
	"github.com/betterego/go-better-admin/server/server/system"
	"gorm.io/gorm"
)

func initMysql() *gorm.DB {
	mysql := global.SYSTEM.Mysql
	db,err := system.LinkDB(mysql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return db
}

