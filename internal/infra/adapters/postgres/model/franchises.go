package model

import "franchises-system/internal/domain/dto"

type Franchises struct {
	ID        int64
	CompanyID int64
	Name      string
	Url       string
	AddressID int64
}

type FranchiseWithLocation struct {
	ID          int64
	Name        string
	Url         string
	Address     string
	ZipCode     string
	CityName    string
	CountryName string
}

func (f *FranchiseWithLocation) ToDomainDTO() dto.Franchise {
	return dto.Franchise{
		ID:   f.ID,
		Name: f.Name,
		Url:  f.Url,
		Location: dto.Location{
			City:    f.CityName,
			Country: f.CountryName,
			Address: f.Address,
			ZipCode: f.ZipCode,
		},
	}
}

type FranchisesWithLocations []FranchiseWithLocation

func (f *FranchisesWithLocations) ToDomainDTO() dto.Franchises {
	var franchises dto.Franchises
	for _, franchise := range *f {
		franchises.Add(franchise.ToDomainDTO())
	}
	return franchises
}
