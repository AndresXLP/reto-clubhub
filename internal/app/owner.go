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

type Owners interface {
	CreateOwner(ctx context.Context, owner dto.Owner) error
	GetOwnerByID(ctx context.Context, ID int64) (dto.Owner, error)
}

type owners struct {
	repo interfaces.Repository
}

func NewOwnerApp(repo interfaces.Repository) Owners {
	return &owners{
		repo,
	}
}

func (app *owners) CreateOwner(ctx context.Context, owner dto.Owner) error {
	if err := app.repo.CreateOwner(ctx, model.Owners{
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
		Phone:     owner.Phone,
	}, owner.Location); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, entity.Response{
			Message: err.Error(),
		})
	}

	return nil
}

func (app *owners) GetOwnerByID(ctx context.Context, ID int64) (dto.Owner, error) {
	owner, err := app.repo.GetOwnerByID(ctx, ID)
	if err != nil {
		return dto.Owner{}, err
	}

	if owner.ID == 0 {
		return dto.Owner{}, echo.NewHTTPError(http.StatusUnprocessableEntity, entity.Response{
			Message: "Owner not found"})
	}

	return owner, err
}
