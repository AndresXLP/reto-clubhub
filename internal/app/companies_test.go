package app_test

import (
	"context"
	"errors"
	"testing"

	"franchises-system/internal/app"
	"franchises-system/internal/domain/dto"
	"franchises-system/internal/infra/adapters/postgres/model"
	mocks2 "franchises-system/mocks/app"
	mocks "franchises-system/mocks/domain/ports/postgres/interfaces"
	"github.com/stretchr/testify/suite"
)

var (
	ctxTest     = context.Background()
	errExpected = errors.New("expected error")

	newCompany = dto.Company{
		Name:      "TEST NAME",
		TaxNumber: "TEST TAX NUMBER",
		OwnerID:   1,
		Location: dto.Location{
			City:    "TEST CITY",
			Country: "TEST COUNTRY",
			Address: "TEST ADDRESS",
			ZipCode: "TEST ZIP CODE",
		},
	}
)

type companiesTestSuite struct {
	suite.Suite
	repo  *mocks.Repository
	owner *mocks2.Owners

	underTest app.Companies
}

func TestCompaniesSuite(t *testing.T) {
	suite.Run(t, new(companiesTestSuite))
}

func (suite *companiesTestSuite) SetupTest() {
	suite.repo = &mocks.Repository{}
	suite.owner = &mocks2.Owners{}
	suite.underTest = app.NewCompaniesApp(suite.repo, suite.owner)
}

func (suite *companiesTestSuite) TestCreateCompany_WhenOwnerNotFound() {
	suite.owner.On("GetOwnerByID", ctxTest, int64(1)).
		Return(dto.Owner{}, errExpected)

	suite.Error(suite.underTest.CreateCompany(ctxTest, newCompany))
}

func (suite *companiesTestSuite) TestCreateCompany_WhenRepoError() {
	suite.owner.On("GetOwnerByID", ctxTest, int64(1)).
		Return(dto.Owner{}, nil)

	suite.repo.On("CreateCompany", ctxTest, model.Companies{
		Name:      newCompany.Name,
		TaxNumber: newCompany.TaxNumber,
		OwnerID:   newCompany.OwnerID,
	}, dto.Location{
		City:    newCompany.Location.City,
		Country: newCompany.Location.Country,
		Address: newCompany.Location.Address,
		ZipCode: newCompany.Location.ZipCode,
	}).
		Return(errExpected)

	suite.Error(suite.underTest.CreateCompany(ctxTest, newCompany))
}

func (suite *companiesTestSuite) TestCreateCompany_WhenSuccess() {
	suite.owner.On("GetOwnerByID", ctxTest, int64(1)).
		Return(dto.Owner{}, nil)

	suite.repo.On("CreateCompany", ctxTest, model.Companies{
		Name:      newCompany.Name,
		TaxNumber: newCompany.TaxNumber,
		OwnerID:   newCompany.OwnerID,
	}, dto.Location{
		City:    newCompany.Location.City,
		Country: newCompany.Location.Country,
		Address: newCompany.Location.Address,
		ZipCode: newCompany.Location.ZipCode,
	}).
		Return(nil)

	suite.NoError(suite.underTest.CreateCompany(ctxTest, newCompany))
}

func (suite *companiesTestSuite) TestGetCompanyByID_WhenRepoError() {
	suite.repo.On("GetCompanyByID", ctxTest, int64(1)).
		Return(dto.Company{}, errExpected)

	_, err := suite.underTest.GetCompanyByID(ctxTest, int64(1))
	suite.Error(err)
}

func (suite *companiesTestSuite) TestGetCompanyByID_WhenCompanyNotFound() {
	suite.repo.On("GetCompanyByID", ctxTest, int64(1)).
		Return(dto.Company{}, nil)

	_, err := suite.underTest.GetCompanyByID(ctxTest, int64(1))
	suite.Error(err)
}

func (suite *companiesTestSuite) TestGetCompanyByID_WhenSuccess() {
	suite.repo.On("GetCompanyByID", ctxTest, int64(1)).
		Return(dto.Company{ID: 1}, nil)

	_, err := suite.underTest.GetCompanyByID(ctxTest, int64(1))
	suite.NoError(err)
}
