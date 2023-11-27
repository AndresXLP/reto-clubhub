package groups

import (
	"franchises-system/internal/infra/api/handlers"
	"github.com/labstack/echo/v4"
)

type CompaniesGroup interface {
	Resource(g *echo.Group)
}

type companiesGroup struct {
	hand handlers.Companies
}

func NewCompaniesGroup(hand handlers.Companies) CompaniesGroup {
	return &companiesGroup{
		hand,
	}
}

func (groups companiesGroup) Resource(g *echo.Group) {
	group := g.Group("/companies")
	group.POST("/", groups.hand.CreateCompany)
}
