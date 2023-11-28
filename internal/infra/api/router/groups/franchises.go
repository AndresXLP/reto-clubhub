package groups

import (
	"franchises-system/internal/infra/api/handlers"
	"github.com/labstack/echo/v4"
)

type FranchisesGroup interface {
	Resource(g *echo.Group)
}

type franchisesGroup struct {
	franchisesHand handlers.Franchises
}

func NewFranchisesGroup(franchisesHand handlers.Franchises) FranchisesGroup {
	return &franchisesGroup{
		franchisesHand,
	}
}

func (groups franchisesGroup) Resource(g *echo.Group) {
	group := g.Group("/franchises")
	group.POST("/", groups.franchisesHand.CreateFranchise)
	group.GET("/details/:name", groups.franchisesHand.GetFranchiseByName)
	group.GET("/company/:company_id", groups.franchisesHand.GetFranchisesByCompanyOwner)
}
