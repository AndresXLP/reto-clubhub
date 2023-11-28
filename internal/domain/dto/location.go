package dto

type Location struct {
	ID      int64  `json:"id,omitempty"`
	City    string `json:"city" validate:"required"`
	Country string `json:"country" validate:"required"`
	Address string `json:"address" validate:"required"`
	ZipCode string `json:"zip_code" validate:"required"`
}
