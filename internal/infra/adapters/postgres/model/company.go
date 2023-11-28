package model

import "franchises-system/internal/domain/dto"

type Companies struct {
	ID        int64
	Name      string
	TaxNumber string
	OwnerID   int64
	AddressID int64
}

type CompanyWithLocations struct {
	ID          int64
	Name        string
	TaxNumber   string
	OwnerID     int64
	AddressID   int64
	Address     string
	ZipCode     string
	CityName    string
	CountryName string
}

func (c *CompanyWithLocations) ToDomainDTO() dto.Company {
	return dto.Company{
		ID:        c.ID,
		Name:      c.Name,
		TaxNumber: c.TaxNumber,
		Location: dto.Location{
			City:    c.CityName,
			Country: c.CountryName,
			Address: c.Address,
			ZipCode: c.ZipCode,
		},
	}
}
