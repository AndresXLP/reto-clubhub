package providers

import (
	"franchises-system/internal/app"
	"franchises-system/internal/infra/adapters/postgres/implementation"
	"franchises-system/internal/infra/api/handlers"
	"franchises-system/internal/infra/api/router"
	"franchises-system/internal/infra/api/router/groups"
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

	_ = Container.Provide(implementation.NewRepository)

	_ = Container.Provide(router.NewRouter)

	_ = Container.Provide(groups.NewFranchisesGroup)
	_ = Container.Provide(groups.NewOwnersGroup)
	_ = Container.Provide(groups.NewCompaniesGroup)

	_ = Container.Provide(handlers.NewFranchisesHandler)
	_ = Container.Provide(handlers.NewOwnersHandler)
	_ = Container.Provide(handlers.NewCompaniesHandler)

	_ = Container.Provide(app.NewFranchisesApp)
	_ = Container.Provide(app.NewOwnerApp)
	_ = Container.Provide(app.NewCompaniesApp)

	return Container
}
