package initialize

import (
	"fmt"
	"github.com/betterego/go-better-admin/server/global"
	"github.com/betterego/go-better-admin/server/util"
	"gorm.io/gorm"

)

func initMysql() *gorm.DB {
	mysql := global.SYSTEM.Mysql
	db,err := util.LinkDB(mysql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return db
}

