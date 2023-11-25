package router

import (
	"franchises-system/internal/infra/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	server *echo.Echo
}

func NewRouter(
	server *echo.Echo,
) *Router {
	return &Router{
		server,
	}
}

func (rtr *Router) Init() {
	rtr.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))
	rtr.server.Use(middleware.Recover())

	base := rtr.server.Group("/api")
	base.GET("/health", handlers.HealthCheck)
}
