package router

import "github.com/betterego/go-better-admin/server/router/system"

type RootRouter struct {
	system.SystemRouter
}

var RootRouters =new(RootRouter)