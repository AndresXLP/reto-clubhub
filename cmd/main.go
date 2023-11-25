package main

import (
	"fmt"

	"franchises-system/cmd/providers"
	"franchises-system/config"
	"franchises-system/internal/infra/api/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

var (
	serverPort = config.Environments().Server.Port
)

func main() {
	container := providers.BuildContainer()

	if err := container.Invoke(func(server *echo.Echo, router *router.Router) {
		router.Init()
		server.Logger.Fatal(server.Start(fmt.Sprintf(":%d", serverPort)))
	}); err != nil {
		log.Panic(err)
	}
}
