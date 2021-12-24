package core

import (
	"fmt"
	"github.com/betterego/go-better-admin/server/config"
	"github.com/betterego/go-better-admin/server/global"
	"net/http"
	"time"
)

func Run()  {
	err:=initServer().ListenAndServe().Error()
	fmt.Println(err)
}

func initServer() *http.Server {
	port := global.SYSTEM.Server.Port
	if port == 0 {
		port = config.DefaulPort
	}
	address := fmt.Sprintf(":%d", port)
	return &http.Server{
		Addr:           address,
		Handler:        global.ROUTER,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}