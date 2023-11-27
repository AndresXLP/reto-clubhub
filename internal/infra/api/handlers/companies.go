package handlers

import (
	"net/http"

	"franchises-system/internal/app"
	"franchises-system/internal/domain/dto"
	"franchises-system/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type Companies interface {
	CreateCompany(c echo.Context) error
}

type company struct {
	app app.Companies
}

func NewCompaniesHandler(app app.Companies) Companies {
	return &company{
		app,
	}
}

func (hand *company) CreateCompany(c echo.Context) error {
	request := dto.Company{}
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
	if err := hand.app.CreateCompany(ctx, request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.Response{
		Message: "Company created successfully",
	})
}
