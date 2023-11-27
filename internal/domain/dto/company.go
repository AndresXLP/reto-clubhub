package dto

type Company struct {
	ID        int64    `json:"ID,omitempty" swaggerignore:"true"`
	Name      string   `json:"name" validate:"required" mod:"ucase" example:"My entreprise holding"`
	TaxNumber string   `json:"tax_number" validate:"required" mod:"ucase" example:"DD79654121"`
	OwnerID   int64    `json:"owner_id,omitempty" validate:"required" example:"1"`
	Location  Location `json:"location" validate:"required"`
}

func (c *Company) Validate() error {
	return validate.Struct(c)
}
