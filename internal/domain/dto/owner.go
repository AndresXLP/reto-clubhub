package dto

type Owner struct {
	ID        int64    `json:"id,omitempty"`
	FirstName string   `json:"first_name" validate:"required" mod:"ucase"`
	LastName  string   `json:"last_name" validate:"required" mod:"ucase"`
	Email     string   `json:"email" validate:"required" mod:"lcase"`
	Phone     string   `json:"phone" validate:"required"`
	Location  Location `json:"location" validate:"required"`
}

func (o *Owner) Validate() error {
	if err := conform.Struct(ctx, o); err != nil {
		return err
	}

	return validate.Struct(o)
}
