package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"franchises-system/internal/domain/dto"
	"franchises-system/internal/infra/api/handlers"
	mocks "franchises-system/mocks/app"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var (
	wrongRequestCompany = dto.Company{
		Name:      "",
		TaxNumber: "",
		OwnerID:   0,
		Location: dto.Location{
			City:    "",
			Country: "",
			Address: "",
			ZipCode: "",
		},
	}

	requestCompany = dto.Company{
		Name:      "TEST NAME",
		TaxNumber: "TEST TAXNUMBER",
		OwnerID:   25,
		Location: dto.Location{
			City:    "TEST CITY",
			Country: "TEST COUNTRY",
			Address: "TEST ADDRESS",
			ZipCode: "TEST ZIPCODE",
		},
	}

	errExpected = errors.New("error occurred")
)

type companiesTestSuite struct {
	suite.Suite
	app *mocks.Companies

	underTest handlers.Companies
}

func TestCompaniesSuite(t *testing.T) {
	suite.Run(t, new(companiesTestSuite))
}

func (suite *companiesTestSuite) SetupTest() {
	suite.app = &mocks.Companies{}
	suite.underTest = handlers.NewCompaniesHandler(suite.app)
}

func (suite *companiesTestSuite) TestCreateCompany_WhenBindFail() {
	body, _ := json.Marshal("")
	controller := SetupController(http.MethodPost, "/api/companies/", bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.CreateCompany(controller.context))
}

func (suite *companiesTestSuite) TestCreateCompany_WhenValidateFail() {
	body, _ := json.Marshal(wrongRequestCompany)
	controller := SetupController(http.MethodPost, "/api/companies/", bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.CreateCompany(controller.context))
}

func (suite *companiesTestSuite) TestCreateCompany_WhenCreateFail() {
	body, _ := json.Marshal(requestCompany)
	controller := SetupController(http.MethodPost, "/api/companies/", bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.On("CreateCompany", ctxTest, requestCompany).
		Return(errExpected)

	suite.Error(suite.underTest.CreateCompany(controller.context))
}

func (suite *companiesTestSuite) TestCreateCompany_WhenCreateSuccess() {
	body, _ := json.Marshal(requestCompany)
	controller := SetupController(http.MethodPost, "/api/companies/", bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.On("CreateCompany", ctxTest, requestCompany).
		Return(nil)

	suite.NoError(suite.underTest.CreateCompany(controller.context))
}
