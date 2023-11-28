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
	wrongRequestOwner = dto.Owner{
		FirstName: "",
		LastName:  "",
		Email:     "",
		Phone:     "",
		Location: dto.Location{
			City:    "",
			Country: "",
			Address: "",
			ZipCode: "",
		},
	}

	requestOwner = dto.Owner{
		FirstName: "TEST NAME",
		LastName:  "TEST LAST NAME",
		Email:     "email@test.com",
		Phone:     "3000000000",
		Location: dto.Location{
			City:    "TEST CITY",
			Country: "TEST COUNTRY",
			Address: "TEST ADDRESS",
			ZipCode: "TEST ZIP CODE",
		},
	}
)

type ownersTestSuite struct {
	suite.Suite
	app *mocks.Owners

	underTest handlers.Owners
}

func TestOwnersSuite(t *testing.T) {
	suite.Run(t, new(ownersTestSuite))
}

func (suite *ownersTestSuite) SetupTest() {
	suite.app = &mocks.Owners{}
	suite.underTest = handlers.NewOwnersHandler(suite.app)
}

func (suite *ownersTestSuite) TestCreateOwner_WhenBindFail() {
	body, _ := json.Marshal("")
	controller := SetupController(http.MethodPost, "/api/owners/", bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.CreateOwner(controller.context))
}

func (suite *ownersTestSuite) TestCreatOwner_WhenValidateFail() {
	body, _ := json.Marshal(wrongRequestOwner)
	controller := SetupController(http.MethodPost, "/api/owners/", bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.Error(suite.underTest.CreateOwner(controller.context))
}

func (suite *ownersTestSuite) TestCreateOwner_WhenCreateFail() {
	body, _ := json.Marshal(requestOwner)
	controller := SetupController(http.MethodPost, "/api/owners/", bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.On("CreateOwner", ctxTest, requestOwner).
		Return(errExpected)

	suite.Error(suite.underTest.CreateOwner(controller.context))
}

func (suite *ownersTestSuite) TestCreateOwner_WhenCreateSuccess() {
	body, _ := json.Marshal(requestOwner)
	controller := SetupController(http.MethodPost, "/api/owners/", bytes.NewBuffer(body))
	controller.Req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	suite.app.On("CreateOwner", ctxTest, requestOwner).
		Return(nil)

	suite.NoError(suite.underTest.CreateOwner(controller.context))
}
