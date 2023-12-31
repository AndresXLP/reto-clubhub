package handlers

import (
	"net/http"

	"franchises-system/internal/app"
	"franchises-system/internal/domain/dto"
	"franchises-system/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type Owners interface {
	CreateOwner(c echo.Context) error
}

type owners struct {
	app app.Owners
}

func NewOwnersHandler(app app.Owners) Owners {
	return &owners{
		app,
	}
}

// @Tags			Owners
// @Summary		Create owner
// @Description	Create owner
// @Produce		json
// @Param			owner	body		dto.Owner	true	"owner"
// @Success		201		{object}	entity.Response
// @Failure		400		{object}	entity.Response
// @Failure		500		{object}	entity.Response
// @Router			/owners/ [post]
func (hand *owners) CreateOwner(c echo.Context) error {
	request := dto.Owner{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Response{
			Message: err.Error(),
		})
	}

	if err := request.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Response{
			Message: err.Error(),
		})
	}

	ctx := c.Request().Context()
	if err := hand.app.CreateOwner(ctx, request); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.Response{
		Message: "Owner created successfully",
	})
}
