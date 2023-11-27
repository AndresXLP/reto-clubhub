package app

import (
	"context"
	"net/http"

	"franchises-system/internal/domain/dto"
	"franchises-system/internal/domain/entity"
	"franchises-system/internal/domain/ports/postgres/interfaces"
	"franchises-system/internal/infra/adapters/postgres/model"
	"github.com/labstack/echo/v4"
)

type Companies interface {
	CreateCompany(ctx context.Context, company dto.Company) error
	GetCompanyByID(ctx context.Context, ID int64) (dto.Company, error)
}

type companies struct {
	repo interfaces.Repository
	Owners
}

func NewCompaniesApp(repo interfaces.Repository, owners Owners) Companies {
	return &companies{
		repo:   repo,
		Owners: owners,
	}
}

func (app *companies) CreateCompany(ctx context.Context, newCompany dto.Company) error {
	_, err := app.GetOwnerByID(ctx, newCompany.OwnerID)
	if err != nil {
		return err
	}

	if err = app.repo.CreateCompany(ctx, model.Companies{
		Name:      newCompany.Name,
		TaxNumber: newCompany.TaxNumber,
		OwnerID:   newCompany.OwnerID,
	}, newCompany.Location); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{
			Message: err.Error(),
		})
	}

	return nil
}

func (app *companies) GetCompanyByID(ctx context.Context, ID int64) (dto.Company, error) {
	company, err := app.repo.GetCompanyByID(ctx, ID)
	if err != nil {
		return dto.Company{}, echo.NewHTTPError(http.StatusInternalServerError, entity.Response{
			Message: err.Error(),
		})
	}

	if company.ID == 0 {
		return dto.Company{}, echo.NewHTTPError(http.StatusNotFound, entity.Response{
			Message: "Company not Found",
		})
	}

	return company, nil
}
