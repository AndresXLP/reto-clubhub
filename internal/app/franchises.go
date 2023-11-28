package app

import (
	"context"

	"franchises-system/internal/domain/dto"
	"franchises-system/internal/domain/ports/postgres/interfaces"
	"franchises-system/internal/infra/adapters/postgres/model"
)

type Franchises interface {
	CreateFranchise(ctx context.Context, request dto.Franchise) error
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
