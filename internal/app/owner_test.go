package app_test

import (
	"testing"

	"franchises-system/internal/app"
	"franchises-system/internal/domain/dto"
	"franchises-system/internal/infra/adapters/postgres/model"
	mocks "franchises-system/mocks/domain/ports/postgres/interfaces"
	"github.com/stretchr/testify/suite"
)

var requestOwner = dto.Owner{
	FirstName: "TEST NAME",
	LastName:  "TEST LAST NAME",
	Email:     "TEST EMAIL",
	Phone:     "TEST PHONE",
	Location: dto.Location{
		City:    "TEST CITY",
		Country: "TEST COUNTRY",
		Address: "TEST ADDRESS",
		ZipCode: "TEST ZIP CODE",
	},
}

type ownerTestSuite struct {
	suite.Suite
	repo *mocks.Repository

	underTest app.Owners
}

func TestOwnerSuite(t *testing.T) {
	suite.Run(t, new(ownerTestSuite))
}

func (suite *ownerTestSuite) SetupTest() {
	suite.repo = &mocks.Repository{}
	suite.underTest = app.NewOwnerApp(suite.repo)
}

func (suite *ownerTestSuite) TestCreateOwner_WhenFail() {
	suite.repo.On("CreateOwner", ctxTest, model.Owners{
		FirstName: requestOwner.FirstName,
		LastName:  requestOwner.LastName,
		Email:     requestOwner.Email,
		Phone:     requestOwner.Phone,
	}, dto.Location{
		City:    request.Location.City,
		Country: request.Location.Country,
		Address: request.Location.Address,
		ZipCode: request.Location.ZipCode,
	}).Return(errExpected)

	suite.Error(suite.underTest.CreateOwner(ctxTest, requestOwner))
}

func (suite *ownerTestSuite) TestCreateOwner_WhenSuccess() {
	suite.repo.On("CreateOwner", ctxTest, model.Owners{
		FirstName: requestOwner.FirstName,
		LastName:  requestOwner.LastName,
		Email:     requestOwner.Email,
		Phone:     requestOwner.Phone,
	}, dto.Location{
		City:    request.Location.City,
		Country: request.Location.Country,
		Address: request.Location.Address,
		ZipCode: request.Location.ZipCode,
	}).Return(nil)

	suite.NoError(suite.underTest.CreateOwner(ctxTest, requestOwner))
}

func (suite *ownerTestSuite) TestGetOwnerByID_WhenFail() {
	suite.repo.On("GetOwnerByID", ctxTest, int64(1)).
		Return(dto.Owner{}, errExpected)

	_, err := suite.underTest.GetOwnerByID(ctxTest, int64(1))
	suite.Error(err)
}

func (suite *ownerTestSuite) TestGetOwnerByID_WhenNotFound() {
	suite.repo.On("GetOwnerByID", ctxTest, int64(1)).
		Return(dto.Owner{}, nil)

	_, err := suite.underTest.GetOwnerByID(ctxTest, int64(1))
	suite.Error(err)
}

func (suite *ownerTestSuite) TestGetOwnerByID_WhenSuccess() {
	suite.repo.On("GetOwnerByID", ctxTest, int64(1)).
		Return(dto.Owner{ID: 1}, nil)

	_, err := suite.underTest.GetOwnerByID(ctxTest, int64(1))
	suite.NoError(err)
}
