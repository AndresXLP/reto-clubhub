package dto

import "github.com/go-playground/validator/v10"

var (
	validate = validator.New()
)

type Franchise struct {
	ID        int64    `json:"ID,omitempty"`
	CompanyID int64    `json:"company_owner" validate:"required"`
	Name      string   `json:"name" validate:"required"`
	Url       string   `json:"url" validate:"required,url_encoded"`
	Location  Location `json:"location" validate:"required"`
}

func (f *Franchise) Validate() error {
	return validate.Struct(f)
}

type Franchises []Franchise

func (f *Franchises) Add(fr Franchise) {
	*f = append(*f, fr)
}
