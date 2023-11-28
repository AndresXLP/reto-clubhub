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

type Franchises interface {
	CreateFranchise(ctx context.Context, request dto.Franchise) error
	GetFranchiseByName(ctx context.Context, name string) (dto.Franchise, error)
	GetFranchisesByCompanyID(ctx context.Context, companyID int64) (dto.FranchiseWithCompany, error)
}

type franchises struct {
	repo interfaces.Repository
	Companies
}

func NewFranchisesApp(repo interfaces.Repository, companies Companies) Franchises {
	return &franchises{
		repo:      repo,
		Companies: companies,
	}
}

func (app *franchises) CreateFranchise(ctx context.Context, request dto.Franchise) error {
	_, err := app.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return err
	}

	franchise := model.Franchises{
		CompanyID: request.CompanyID,
		Name:      request.Name,
		Url:       request.Url,
	}

	if err = app.repo.CreateFranchise(ctx, franchise, request.Location); err != nil {
		return err
	}
	return nil
}

func (app *franchises) GetFranchiseByName(ctx context.Context, name string) (dto.Franchise, error) {
	franchise, err := app.repo.GetFranchiseByName(ctx, name)
	if err != nil {
		return dto.Franchise{}, err
	}

	if franchise.ID == 0 {
		return dto.Franchise{}, echo.NewHTTPError(http.StatusNotFound, entity.Response{
			Message: "Franchise not found",
		})
	}

	return franchise, nil
}

func (app *franchises) GetFranchisesByCompanyID(ctx context.Context, companyID int64) (dto.FranchiseWithCompany, error) {
	company, err := app.GetCompanyByID(ctx, companyID)
	if err != nil {
		return dto.FranchiseWithCompany{}, err
	}

	franchise, err := app.repo.GetFranchisesByCompanyID(ctx, companyID)
	if err != nil {
		return dto.FranchiseWithCompany{}, err
	}

	return dto.FranchiseWithCompany{
		Company:   company,
		Franchise: franchise,
	}, nil
}
