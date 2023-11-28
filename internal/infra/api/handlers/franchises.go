package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"franchises-system/internal/app"
	"franchises-system/internal/domain/dto"
	"franchises-system/internal/domain/entity"
	"github.com/labstack/echo/v4"
)

type Franchises interface {
	CreateFranchise(c echo.Context) error
	GetFranchiseByName(c echo.Context) error
	GetFranchisesByCompanyOwner(c echo.Context) error
}

type franchises struct {
	app app.Franchises
}

func NewFranchisesHandler(app app.Franchises) Franchises {
	return &franchises{
		app,
	}
}

// @Tags			Franchises
// @Summary		Create a franchise
// @Description	Create a franchise
// @Produce		json
// @Param			franchise	body		dto.Franchise	true	"Franchise"
// @Success		201			{object}	entity.Response
// @Failure		400			{object}	entity.Response
// @Router			/franchises/ [post]
func (hand *franchises) CreateFranchise(c echo.Context) error {
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
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity.Response{
		Message: "Franchise created successfully",
	})
}

// @Tags			Franchises
// @Summary		Get a franchise by name
// @Description	Get a franchise by name
// @Produce		json
// @Param			name	path		string	true	"Franchise name"
// @Success		200		{object}	entity.Response{data=dto.Franchise}
// @Failure		400		{object}	entity.Response
// @Router			/franchises/details/{name} [get]
func (hand *franchises) GetFranchiseByName(c echo.Context) error {
	request := strings.ToUpper(c.Param("name"))
	if request == "" {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Response{
			Message: "Name is required"})
	}

	ctx := c.Request().Context()
	franchise, err := hand.app.GetFranchiseByName(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{
		Message: "Franchise found successfully",
		Data:    franchise,
	})
}

// @Tags			Franchises
// @Summary		Get franchises by company owner
// @Description	Get franchises by company owner
// @Produce		json
// @Param			company_id	path		int	true	"Company ID"
// @Success		200			{object}	entity.Response{data=dto.FranchiseWithCompany}
// @Failure		400			{object}	entity.Response
// @Router			/franchises/company/{company_id} [get]
func (hand *franchises) GetFranchisesByCompanyOwner(c echo.Context) error {
	companyID, err := strconv.ParseInt(c.Param("company_id"), 10, 64)
	if err != nil || companyID == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, entity.Response{
			Message: "Company ID is required"})
	}

	ctx := c.Request().Context()
	franchisesWithCompany, err := hand.app.GetFranchisesByCompanyID(ctx, companyID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, entity.Response{
		Message: "Franchises found successfully",
		Data:    franchisesWithCompany,
	})
}
