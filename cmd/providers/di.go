package providers

import (
	"franchises-system/internal/infra/api/router"
	"franchises-system/internal/infra/resources/postgres"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(postgres.NewPostgresConnection)

	_ = Container.Provide(router.NewRouter)

	return Container
}
