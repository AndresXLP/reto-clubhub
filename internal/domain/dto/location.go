package dto

type Location struct {
	ID      int64  `json:"id,omitempty" swaggerignore:"true"`
	City    string `json:"city" validate:"required" mod:"ucase" example:"Toronto"`
	Country string `json:"country" validate:"required" mod:"ucase" example:"Canada"`
	Address string `json:"address" validate:"required" mod:"ucase" example:"78 Rober ST"`
	ZipCode string `json:"zip_code" validate:"required" mod:"ucase" example:"F9A 92O"`
}
