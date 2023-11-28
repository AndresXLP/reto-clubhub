package dto

type Company struct {
	ID        int64    `json:"ID,omitempty"`
	Name      string   `json:"name" validate:"required"`
	TaxNumber string   `json:"taxNumber" validate:"required"`
	OwnerID   int64    `json:"owner_id" validate:"required"`
	Location  Location `json:"location" validate:"required"`
}

func (c *Company) Validate() error {
	return validate.Struct(c)
}
