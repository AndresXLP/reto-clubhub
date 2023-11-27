package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"franchises-system/internal/domain/dto"
	"franchises-system/internal/infra/api/handlers"
	mocks "franchises-system/mocks/app"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var (
	wrongRequestFranshise = dto.Franchise{
		CompanyID: 0,
		Name:      "",
		Url:       "",
		Location: dto.Location{
			City:    "",
			Country: "",
			Address: "",
			ZipCode: "",
		},
	}

	requestFranchise = dto.Franchise{
		Name:      "TEST NAME",
		Url:       "www.test.com",
		CompanyID: 25,
		Location: dto.Location{
			City:    "TEST CITY",
			Country: "TEST COUNTRY",
			Address: "TEST ADDRESS",
			ZipCode: "TEST ZIPCODE",
		},
	}
)

type franchisesTestSuite struct {
	suite.Suite
	app *mocks.Franchises

	underTest handlers.Franchises
}

func TestFranchisesSuite(t *testing.T) {
	suite.Run(t, new(franchisesTestSuite))
}

func (suite *franchisesTestSuite) SetupTest() {
	suite.app = &mocks.Franchises{}
	suite.underTest = handlers.NewFranchisesHandler(suite.app)
}

func (suite *franchisesTestSuite) TestCreateFranchise_WhenBindFail() {
	body, _ := json.Marshal("")
	controller := SetupController(http.MethodPost, "/api/franchises/", bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.CreateFranchise(controller.context))
}

func (suite *franchisesTestSuite) TestCreateFranchise_WhenValidateFail() {
	body, _ := json.Marshal(wrongRequestFranshise)
	controller := SetupController(http.MethodPost, "/api/franchises/", bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.CreateFranchise(controller.context))
}

func (suite *franchisesTestSuite) TestCreateFranchise_WhenCreateFail() {
	body, _ := json.Marshal(requestFranchise)
	controller := SetupController(http.MethodPost, "/api/franchises/", bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.On("CreateFranchise", ctxTest, requestFranchise).
		Return(errExpected)

	suite.Error(suite.underTest.CreateFranchise(controller.context))
}

func (suite *franchisesTestSuite) TestCreateFranchise_WhenCreateSuccess() {
	body, _ := json.Marshal(requestFranchise)
	controller := SetupController(http.MethodPost, "/api/franchises/", bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.On("CreateFranchise", ctxTest, requestFranchise).
		Return(nil)

	suite.NoError(suite.underTest.CreateFranchise(controller.context))
}

func (suite *franchisesTestSuite) TestGetFranchiseByName_WhenNameIsEmpty() {
	controller := SetupController(http.MethodGet, "/api/franchises/details/", nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	controller.context.SetParamNames("name")
	controller.context.SetParamValues("")

	suite.Error(suite.underTest.GetFranchiseByName(controller.context))
}

func (suite *franchisesTestSuite) TestGetFranchiseByName_WhenGetFail() {
	controller := SetupController(http.MethodGet, "/api/franchises/details/test%20name", nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	controller.context.SetParamNames("name")
	controller.context.SetParamValues("test name")

	suite.app.On("GetFranchiseByName", ctxTest, "TEST NAME").
		Return(dto.Franchise{}, errExpected)

	suite.Error(suite.underTest.GetFranchiseByName(controller.context))
}

func (suite *franchisesTestSuite) TestGetFranchiseByName_WhenGetSuccess() {
	controller := SetupController(http.MethodGet, "/api/franchises/details/test%20name", nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	controller.context.SetParamNames("name")
	controller.context.SetParamValues("test name")

	suite.app.On("GetFranchiseByName", ctxTest, "TEST NAME").
		Return(dto.Franchise{}, nil)

	suite.NoError(suite.underTest.GetFranchiseByName(controller.context))
}

func (suite *franchisesTestSuite) TestGetFranchisesByCompanyOwner_WhenCompanyIDIsEmpty() {
	controller := SetupController(http.MethodGet, "/api/franchises/details/company/", nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	controller.context.SetParamNames("company_id")
	controller.context.SetParamValues("")

	suite.Error(suite.underTest.GetFranchisesByCompanyOwner(controller.context))
}

func (suite *franchisesTestSuite) TestGetFranchisesByCompanyOwner_WhenGetFail() {
	controller := SetupController(http.MethodGet, "/api/franchises/details/company/25", nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	controller.context.SetParamNames("company_id")
	controller.context.SetParamValues("25")

	suite.app.On("GetFranchisesByCompanyID", ctxTest, int64(25)).
		Return(dto.FranchiseWithCompany{}, errExpected)

	suite.Error(suite.underTest.GetFranchisesByCompanyOwner(controller.context))
}

func (suite *franchisesTestSuite) TestGetFranchisesByCompanyOwner_WhenGetSuccess() {
	controller := SetupController(http.MethodGet, "/api/franchises/details/company/25", nil)
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	controller.context.SetParamNames("company_id")
	controller.context.SetParamValues("25")

	suite.app.On("GetFranchisesByCompanyID", ctxTest, int64(25)).
		Return(dto.FranchiseWithCompany{}, nil)

	suite.NoError(suite.underTest.GetFranchisesByCompanyOwner(controller.context))
}
