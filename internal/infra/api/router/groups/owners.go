package groups

import (
	"franchises-system/internal/infra/api/handlers"
	"github.com/labstack/echo/v4"
)

type OwnersGroup interface {
	Resource(g *echo.Group)
}

type ownersGroup struct {
	hand handlers.Owners
}

func NewOwnersGroup(hand handlers.Owners) OwnersGroup {
	return &ownersGroup{
		hand,
	}
}

func (groups ownersGroup) Resource(g *echo.Group) {
	group := g.Group("/owners")
	group.POST("/", groups.hand.CreateOwner)
}
