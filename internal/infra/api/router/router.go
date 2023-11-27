package router

import (
	_ "franchises-system/docs"
	"franchises-system/internal/infra/api/handlers"
	"franchises-system/internal/infra/api/router/groups"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	server    *echo.Echo
	franchise groups.FranchisesGroup
	owner     groups.OwnersGroup
	company   groups.CompaniesGroup
}

func NewRouter(
	server *echo.Echo,
	franchise groups.FranchisesGroup,
	owner groups.OwnersGroup,
	company groups.CompaniesGroup,
) *Router {
	return &Router{
		server,
		franchise,
		owner,
		company,
	}
}

func (rtr *Router) Init() {
	rtr.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status} latency=${latency_human}\n",
	}))
	rtr.server.Use(middleware.Recover())

	base := rtr.server.Group("/api")
	base.GET("/health", handlers.HealthCheck)
	base.GET("/swagger/*", echoSwagger.WrapHandler)

	rtr.franchise.Resource(base)
	rtr.owner.Resource(base)
	rtr.company.Resource(base)
}
