package main

import (
	_ "github.com/betterego/go-better-admin/server/core"
	_ "github.com/betterego/go-better-admin/server/initialize"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,"ok")
	})
	r.Run(":8088")
}