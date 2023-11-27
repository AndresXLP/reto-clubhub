package dto

type Location struct {
	ID      int64  `json:"id,omitempty"`
	City    string `json:"city" validate:"required" mod:"ucase"`
	Country string `json:"country" validate:"required" mod:"ucase"`
	Address string `json:"address" validate:"required" mod:"ucase"`
	ZipCode string `json:"zip_code" validate:"required" mod:"ucase"`
}
