package dto

type Owner struct {
	ID        int64    `json:"id,omitempty"`
	FirstName string   `json:"first_name" validate:"required"`
	LastName  string   `json:"last_name" validate:"required"`
	Email     string   `json:"email" validate:"required"`
	Phone     string   `json:"phone" validate:"required"`
	Location  Location `json:"location" validate:"required"`
}

func (o *Owner) Validate() error {
	return validate.Struct(o)
}
