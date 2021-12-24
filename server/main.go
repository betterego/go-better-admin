package main

import (
	"github.com/betterego/go-better-admin/server/core"
	"github.com/betterego/go-better-admin/server/global"
	_ "github.com/betterego/go-better-admin/server/initialize"
)

func main() {
	if global.DB != nil {
		db,_ := global.DB.DB()
		defer db.Close()
	}
	core.Run()
}