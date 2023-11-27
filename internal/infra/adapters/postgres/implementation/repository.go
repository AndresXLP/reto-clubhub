package implementation

import (
	"context"

	"franchises-system/internal/domain/dto"
	"franchises-system/internal/domain/ports/postgres/interfaces"
	"franchises-system/internal/infra/adapters/postgres/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.Repository {
	return &repository{
		db,
	}
}

//Franchises Repository

func (repo repository) CreateFranchise(ctx context.Context, newFranchise model.Franchises, location dto.Location) error {
	country := model.Countries{Name: location.Country}
	city := model.Cities{Name: location.City}
	address := model.Addresses{
		Address: location.Address,
		ZipCode: location.ZipCode,
	}

	if err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := repo.getLocation(tx, model.Locations{
			Countries: &country,
			Cities:    &city,
			Addresses: &address,
		}); err != nil {
			return err
		}

		newFranchise.AddressID = address.ID
		if err := tx.Create(&newFranchise).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo repository) GetFranchiseByID(ctx context.Context, ID int64) (dto.Franchise, error) {
	franchise := model.FranchiseWithLocation{}
	if err := repo.db.WithContext(ctx).
		Table("franchises").
		Select("franchises.id, franchises.name, franchises.tax_number, addresses.address, "+
			"addresses.zip_code, cities.name as city_name, countries.name as country_name").
		Joins("INNER JOIN addresses ON addresses.id = franchises.address_id").
		Joins("INNER JOIN cities ON cities.id = addresses.city_id").
		Joins("INNER JOIN countries ON countries.id = cities.country_id").
		Where("franchises.id = ?", ID).
		Scan(&franchise).Error; err != nil {
		return dto.Franchise{}, err
	}

	return franchise.ToDomainDTO(), nil
}

func (repo repository) GetFranchisesByCompanyID(ctx context.Context, ID int64) (dto.Franchises, error) {
	franchises := model.FranchisesWithLocations{}
	if err := repo.db.WithContext(ctx).
		Table("franchises").
		Select("franchises.id, franchises.name, franchises.url, addresses.address, "+
			"addresses.zip_code, cities.name as city_name, countries.name as country_name").
		Joins("INNER JOIN addresses ON addresses.id = franchises.address_id").
		Joins("INNER JOIN cities ON cities.id = addresses.city_id").
		Joins("INNER JOIN countries ON countries.id = cities.country_id").
		Where("franchises.company_id = ?", ID).
		Scan(&franchises).Error; err != nil {
		return dto.Franchises{}, err
	}

	return franchises.ToDomainDTO(), nil
}

func (repo repository) GetFranchiseByName(ctx context.Context, name string) (dto.Franchise, error) {
	franchise := model.FranchiseWithLocation{}
	if err := repo.db.WithContext(ctx).
		Table("franchises").
		Select("franchises.id, franchises.name, franchises.url, addresses.address,"+
			"addresses.zip_code, cities.name as city_name, countries.name as country_name").
		Joins("INNER JOIN addresses ON addresses.id = franchises.address_id").
		Joins("INNER JOIN cities ON cities.id = addresses.city_id").
		Joins("INNER JOIN countries ON countries.id = cities.country_id").
		Where("franchises.name = ?", name).
		Scan(&franchise).Error; err != nil {
		return dto.Franchise{}, err
	}

	return franchise.ToDomainDTO(), nil
}

//Company Repository

func (repo repository) CreateCompany(ctx context.Context, companies model.Companies, location dto.Location) error {
	country := model.Countries{Name: location.Country}
	city := model.Cities{Name: location.City}
	address := model.Addresses{
		Address: location.Address,
		ZipCode: location.ZipCode,
	}

	if err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := repo.getLocation(tx, model.Locations{
			Countries: &country,
			Cities:    &city,
			Addresses: &address,
		}); err != nil {
			return err
		}

		companies.AddressID = address.ID
		if err := tx.Create(&companies).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo repository) GetCompanyByID(ctx context.Context, ID int64) (dto.Company, error) {
	company := model.CompanyWithLocations{}
	if err := repo.db.WithContext(ctx).
		Table("companies").
		Select("companies.*, addresses.address, "+
			"addresses.zip_code, cities.name as city_name, countries.name as country_name").
		Joins("INNER JOIN addresses ON addresses.id = companies.address_id").
		Joins("INNER JOIN cities ON cities.id = addresses.city_id").
		Joins("INNER JOIN countries ON countries.id = cities.country_id").
		Where("companies.id = ?", ID).
		Scan(&company).Error; err != nil {
		return dto.Company{}, err
	}

	return company.ToDomainDTO(), nil
}

//Owner Repository

func (repo repository) CreateOwner(ctx context.Context, newOwner model.Owners, location dto.Location) error {
	country := model.Countries{Name: location.Country}
	city := model.Cities{Name: location.City}
	address := model.Addresses{
		Address: location.Address,
		ZipCode: location.ZipCode,
	}

	if err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := repo.getLocation(tx, model.Locations{
			Countries: &country,
			Cities:    &city,
			Addresses: &address,
		}); err != nil {
			return err
		}

		newOwner.AddressID = address.ID
		if err := tx.Create(&newOwner).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo repository) GetOwnerByID(ctx context.Context, ID int64) (dto.Owner, error) {
	owner := model.OwnerWithLocation{}
	if err := repo.db.WithContext(ctx).
		Table("owners").
		Select("owners.id, owners.first_name, owners.last_name, owners.email, owners.phone, addresses.address, "+
			"addresses.zip_code, cities.name as city_name, countries.name as country_name").
		Joins("INNER JOIN addresses ON addresses.id = owners.address_id").
		Joins("INNER JOIN cities ON cities.id = addresses.city_id").
		Joins("INNER JOIN countries ON countries.id = cities.country_id").
		Where("owners.id = ?", ID).
		Scan(&owner).Error; err != nil {
		return dto.Owner{}, err
	}

	return owner.ToDomainDTO(), nil
}

// Auxiliary Methods Repository

func (repo repository) getLocation(tx *gorm.DB, location model.Locations) error {
	if err := repo.getCountry(tx, location.Countries); err != nil {
		return err
	}

	location.Cities.CountryID = location.Countries.ID
	if err := repo.getCity(tx, location.Cities); err != nil {
		return err
	}

	location.Addresses.CityID = location.Cities.ID
	if err := repo.getAddress(tx, location.Addresses); err != nil {
		return err
	}

	return nil
}

func (repo repository) getCountry(tx *gorm.DB, country *model.Countries) error {
	if err := tx.FirstOrCreate(country, "name = ?",
		country.Name).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo repository) getCity(tx *gorm.DB, city *model.Cities) error {
	if err := tx.FirstOrCreate(city, "name = ? AND country_id = ?",
		&city.Name, &city.CountryID).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo repository) getAddress(tx *gorm.DB, location *model.Addresses) error {
	if err := tx.FirstOrCreate(location, "address = ? AND zip_code = ? AND city_id = ?",
		&location.Address, &location.ZipCode, &location.CityID).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
