package handlers

import (
	"net/http"

	"franchises-system/internal/app"
	"franchises-system/internal/domain/dto"
	"franchises-system/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type Franchises interface {
	Create(c echo.Context) error
}

type franchises struct {
	app app.Franchises
}

func NewFranchisesHandler(app app.Franchises) Franchises {
	return &franchises{
		app,
	}
}

func (hand *franchises) Create(c echo.Context) error {
	ctx := c.Request().Context()

	request := dto.Franchise{}
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Response{
			Message: err.Error(),
		})
	}

	if err := request.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{
			Message: err.Error()})
	}

	if err := hand.app.CreateFranchise(ctx, request); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, entity.Response{
		Message: "Franchise created successfully",
	})
}
