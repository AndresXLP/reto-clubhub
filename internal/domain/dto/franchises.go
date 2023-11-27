package dto

import (
	"context"

	"github.com/go-playground/mold/v4/modifiers"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
	conform  = modifiers.New()
	ctx      = context.Background()
)

type Franchise struct {
	ID        int64    `json:"ID,omitempty"`
	CompanyID int64    `json:"company_owner,omitempty" validate:"required"`
	Name      string   `json:"name" validate:"required" mod:"ucase"`
	Url       string   `json:"url" validate:"required,url_encoded" mod:"lcase"`
	Location  Location `json:"location" validate:"required"`
}

func (f *Franchise) Validate() error {
	if err := conform.Struct(ctx, f); err != nil {
		return err
	}

	return validate.Struct(f)
}

type Franchises []Franchise

func (f *Franchises) Add(fr ...Franchise) {
	*f = append(*f, fr...)
}

type FranchiseWithCompany struct {
	Company   Company     `json:"company"`
	Franchise []Franchise `json:"franchises"`
}

func (f *FranchiseWithCompany) Add(fr ...Franchise) {
	f.Franchise = append(f.Franchise, fr...)
}
