package interfaces

import (
	"context"

	"franchises-system/internal/domain/dto"
	"franchises-system/internal/infra/adapters/postgres/model"
)

type Repository interface {
	CreateFranchise(ctx context.Context, newFranchise model.Franchises, locations dto.Location) error
	GetFranchiseByID(ctx context.Context, ID int64) (dto.Franchise, error)
	GetFranchisesByCompanyID(ctx context.Context, ID int64) (dto.Franchises, error)
	GetFranchiseByName(ctx context.Context, name string) (dto.Franchise, error)
	CreateOwner(ctx context.Context, newOwner model.Owners, location dto.Location) error
	GetOwnerByID(ctx context.Context, ID int64) (dto.Owner, error)
	CreateCompany(ctx context.Context, companies model.Companies, location dto.Location) error
	GetCompanyByID(ctx context.Context, ID int64) (dto.Company, error)
}
