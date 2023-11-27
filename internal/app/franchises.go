package app

import (
	"context"
	"net/http"

	"franchises-system/internal/domain/dto"
	"franchises-system/internal/domain/entity"
	"franchises-system/internal/domain/ports/postgres/interfaces"
	"franchises-system/internal/infra/adapters/postgres/model"
	"franchises-system/pkg/strings"
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
	WebInfo
}

func NewFranchisesApp(repo interfaces.Repository, companies Companies, webInfo WebInfo) Franchises {
	return &franchises{
		repo:      repo,
		Companies: companies,
		WebInfo:   webInfo,
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

	ID, err := app.repo.CreateFranchise(ctx, franchise, request.Location)
	if err != nil {
		return err
	}

	webInfo := make(chan entity.WebInfo)
	errChan := make(chan error)

	go func() {
		defer close(webInfo)
		defer close(errChan)
		errChan <- app.GetWebInfo(ctx, strings.CleanURL(request.Url), webInfo)
	}()

	info := <-webInfo
	err = app.repo.SetAdditionalInfoFranchise(ctx, model.AdditionalFranchiseInfo{
		FranchiseId:           ID,
		Protocol:              info.Protocol,
		TraceRoutes:           info.TraceRoutes,
		DomainCreatedAt:       info.Domain.CreatedAt,
		DomainExpiredAt:       info.Domain.ExpiredAt,
		DomainRegistrantName:  info.Domain.RegistrantName,
		DomainRegistrantEmail: info.Domain.RegistrantEmail,
	})

	if err != nil {
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
