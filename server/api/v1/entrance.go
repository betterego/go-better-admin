package v1

import "github.com/betterego/go-better-admin/server/api/v1/system"

type RootGroup struct {
	system.SystemGroup
}
var RootGroups = new(RootGroup)