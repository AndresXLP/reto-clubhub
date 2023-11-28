package app_test

import (
	"testing"
	"time"

	"franchises-system/internal/app"
	"franchises-system/internal/domain/dto"
	"franchises-system/internal/domain/entity"
	"franchises-system/internal/infra/adapters/postgres/model"
	"franchises-system/internal/utils/strings"
	mocks2 "franchises-system/mocks/app"
	mocks "franchises-system/mocks/domain/ports/postgres/interfaces"
	"github.com/stretchr/testify/suite"
)

var (
	request = dto.Franchise{
		CompanyID: 1,
		Name:      "TEST FRANCHISE",
		Url:       "www.test.com",
		Location: dto.Location{
			City:    "TEST CITY",
			Country: "TEST COUNTRY",
			Address: "TEST ADDRESS",
			ZipCode: "TEST ZIP CODE",
		},
	}

	webInfoEntity = entity.WebInfo{
		FranchiseID: 1,
		UrlImage:    "",
		Protocol:    "http",
		TraceRoutes: []string{"1", "2"},
		Domain: entity.Domain{
			CreatedAt:       time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
			ExpiredAt:       time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC),
			RegistrantName:  "TEST REGISTRANT NAME",
			RegistrantEmail: "TEST REGISTRANT EMAIL",
		},
	}
)

type franchisesTestSuite struct {
	suite.Suite
	repo *mocks.Repository
	*mocks2.Companies
	*mocks2.WebInfo

	underTest app.Franchises
}

func TestFranchisesSuite(t *testing.T) {
	suite.Run(t, new(franchisesTestSuite))
}

func (suite *franchisesTestSuite) SetupTest() {
	suite.repo = &mocks.Repository{}
	suite.Companies = &mocks2.Companies{}
	suite.WebInfo = &mocks2.WebInfo{}
	suite.underTest = app.NewFranchisesApp(suite.repo, suite.Companies, suite.WebInfo)
}

func (suite *franchisesTestSuite) TestCreateFranchise_WhenCompanyNotFound() {
	suite.Companies.On("GetCompanyByID", ctxTest, request.CompanyID).
		Return(dto.Company{}, errExpected)

	suite.Error(suite.underTest.CreateFranchise(ctxTest, request))
}

func (suite *franchisesTestSuite) TestCreateFranchise_WhenCreateFail() {
	suite.Companies.On("GetCompanyByID", ctxTest, request.CompanyID).
		Return(dto.Company{}, nil)

	franchise := model.Franchises{
		CompanyID: request.CompanyID,
		Name:      request.Name,
		Url:       request.Url,
	}

	suite.repo.On("CreateFranchise", ctxTest, franchise, request.Location).
		Return(int64(0), errExpected)

	suite.Error(suite.underTest.CreateFranchise(ctxTest, request))
}

func (suite *franchisesTestSuite) TestCreateFranchise_WhenGetWebInfoFail() {
	suite.Companies.On("GetCompanyByID", ctxTest, request.CompanyID).
		Return(dto.Company{}, nil)

	franchise := model.Franchises{
		CompanyID: request.CompanyID,
		Name:      request.Name,
		Url:       request.Url,
	}

	suite.repo.On("CreateFranchise", ctxTest, franchise, request.Location).
		Return(int64(1), nil)

	suite.WebInfo.On("GetWebInfo", ctxTest, strings.CleanURL(request.Url)).
		Return(entity.WebInfo{}, errExpected)

	suite.Error(suite.underTest.CreateFranchise(ctxTest, request))
}

func (suite *franchisesTestSuite) TestCreateFranchise_WhenSetAdditionalInfoFail() {
	suite.Companies.On("GetCompanyByID", ctxTest, request.CompanyID).
		Return(dto.Company{}, nil)

	franchise := model.Franchises{
		CompanyID: request.CompanyID,
		Name:      request.Name,
		Url:       request.Url,
	}

	suite.repo.On("CreateFranchise", ctxTest, franchise, request.Location).
		Return(int64(1), nil)

	suite.WebInfo.On("GetWebInfo", ctxTest, strings.CleanURL(request.Url)).
		Return(webInfoEntity, nil)

	additionalInfo := model.AdditionalFranchiseInfo{
		FranchiseId:           1,
		UrlImage:              "",
		Protocol:              webInfoEntity.Protocol,
		TraceRoutes:           webInfoEntity.TraceRoutes,
		DomainCreatedAt:       webInfoEntity.Domain.CreatedAt,
		DomainExpiredAt:       webInfoEntity.Domain.ExpiredAt,
		DomainRegistrantName:  webInfoEntity.Domain.RegistrantName,
		DomainRegistrantEmail: webInfoEntity.Domain.RegistrantEmail,
	}
	suite.repo.On("SetAdditionalInfoFranchise", ctxTest, additionalInfo).
		Return(errExpected)

	suite.Error(suite.underTest.CreateFranchise(ctxTest, request))
}

func (suite *franchisesTestSuite) TestCreateFranchise_WhenSuccess() {
	suite.Companies.On("GetCompanyByID", ctxTest, request.CompanyID).
		Return(dto.Company{}, nil)

	franchise := model.Franchises{
		CompanyID: request.CompanyID,
		Name:      request.Name,
		Url:       request.Url,
	}

	suite.repo.On("CreateFranchise", ctxTest, franchise, request.Location).
		Return(int64(1), nil)

	suite.WebInfo.On("GetWebInfo", ctxTest, strings.CleanURL(request.Url)).
		Return(webInfoEntity, nil)

	additionalInfo := model.AdditionalFranchiseInfo{
		FranchiseId:           1,
		UrlImage:              "",
		Protocol:              webInfoEntity.Protocol,
		TraceRoutes:           webInfoEntity.TraceRoutes,
		DomainCreatedAt:       webInfoEntity.Domain.CreatedAt,
		DomainExpiredAt:       webInfoEntity.Domain.ExpiredAt,
		DomainRegistrantName:  webInfoEntity.Domain.RegistrantName,
		DomainRegistrantEmail: webInfoEntity.Domain.RegistrantEmail,
	}
	suite.repo.On("SetAdditionalInfoFranchise", ctxTest, additionalInfo).
		Return(nil)

	suite.NoError(suite.underTest.CreateFranchise(ctxTest, request))
}

func (suite *franchisesTestSuite) TestGetFranchiseByName_WhenFail() {
	suite.repo.On("GetFranchiseByName", ctxTest, request.Name).
		Return(dto.Franchise{}, errExpected)

	_, err := suite.underTest.GetFranchiseByName(ctxTest, request.Name)
	suite.Error(err)
}

func (suite *franchisesTestSuite) TestGetFranchiseByName_WhenNotFound() {
	suite.repo.On("GetFranchiseByName", ctxTest, request.Name).
		Return(dto.Franchise{}, nil)

	_, err := suite.underTest.GetFranchiseByName(ctxTest, request.Name)
	suite.Error(err)
}

func (suite *franchisesTestSuite) TestGetFranchiseByName_WhenSuccess() {
	suite.repo.On("GetFranchiseByName", ctxTest, request.Name).
		Return(dto.Franchise{ID: 1}, nil)

	_, err := suite.underTest.GetFranchiseByName(ctxTest, request.Name)
	suite.NoError(err)
}

func (suite *franchisesTestSuite) TestGetFranchisesByCompanyID_WhenCompanyFail() {
	suite.Companies.On("GetCompanyByID", ctxTest, request.CompanyID).
		Return(dto.Company{}, errExpected)

	_, err := suite.underTest.GetFranchisesByCompanyID(ctxTest, request.CompanyID)
	suite.Error(err)
}

func (suite *franchisesTestSuite) TestGetFranchisesByCompanyID_WhenFail() {
	suite.Companies.On("GetCompanyByID", ctxTest, request.CompanyID).
		Return(dto.Company{}, nil)

	suite.repo.On("GetFranchisesByCompanyID", ctxTest, request.CompanyID).
		Return(dto.Franchises{}, errExpected)

	_, err := suite.underTest.GetFranchisesByCompanyID(ctxTest, request.CompanyID)
	suite.Error(err)
}

func (suite *franchisesTestSuite) TestGetFranchisesByCompanyID_WhenSuccess() {
	suite.Companies.On("GetCompanyByID", ctxTest, request.CompanyID).
		Return(dto.Company{}, nil)

	suite.repo.On("GetFranchisesByCompanyID", ctxTest, request.CompanyID).
		Return(dto.Franchises{}, nil)

	_, err := suite.underTest.GetFranchisesByCompanyID(ctxTest, request.CompanyID)
	suite.NoError(err)
}
