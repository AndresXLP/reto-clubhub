package model

import "franchises-system/internal/domain/dto"

type Owners struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Phone     string
	AddressID int64
}

type OwnerWithLocation struct {
	ID          int64
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	Address     string
	ZipCode     string
	CityName    string
	CountryName string
}

func (o *OwnerWithLocation) ToDomainDTO() dto.Owner {
	return dto.Owner{
		ID:        o.ID,
		FirstName: o.FirstName,
		LastName:  o.LastName,
		Email:     o.Email,
		Phone:     o.Phone,
		Location: dto.Location{
			City:    o.CityName,
			Country: o.CountryName,
			Address: o.Address,
			ZipCode: o.ZipCode,
		},
	}
}
