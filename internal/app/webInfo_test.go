package app_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"franchises-system/internal/app"
	"franchises-system/internal/constants"
	"franchises-system/internal/domain/entity"
	mocks2 "franchises-system/mocks/domain/ports/postgres/interfaces"
	mocks "franchises-system/mocks/utils/http"
	"github.com/stretchr/testify/suite"
)

var respHttp = entity.SSLLabsResponse{
	Host:            "testurl.com",
	Port:            443,
	Protocol:        "https",
	IsPublic:        false,
	Status:          "READY",
	StartTime:       0,
	TestTime:        0,
	EngineVersion:   "1.0.0",
	CriteriaVersion: "1.0.0",
	Endpoints: entity.Endpoints{
		entity.Endpoint{
			IPAddress:         "127.0.0.1",
			ServerName:        "testurl.com",
			StatusMessage:     "Ready",
			Grade:             "A+",
			GradeTrustIgnored: "A+",
			HasWarnings:       false,
			IsExceptional:     false,
			Progress:          0,
			Duration:          0,
			Delegation:        0,
		},
	},
}

type webInfoTestSuite struct {
	suite.Suite
	http *mocks.HttpClient
	repo *mocks2.Repository

	underTest app.WebInfo
}

func TestWebInfoSuite(t *testing.T) {
	suite.Run(t, new(webInfoTestSuite))
}

func (suite *webInfoTestSuite) SetupTest() {
	suite.http = &mocks.HttpClient{}
	suite.repo = &mocks2.Repository{}
	suite.underTest = app.NewWebInfoApp(suite.http, suite.repo)
}

func (suite *webInfoTestSuite) TestGetWebInfo_WhenGetFail() {
	httpResp := http.Response{}
	suite.http.On("Get", ctxTest, fmt.Sprintf(constants.Ssllabs, "testurl.com")).
		Return(&httpResp, errExpected)

	_, err := suite.underTest.GetWebInfo(ctxTest, "testurl.com")
	suite.Error(err)
}

func (suite *webInfoTestSuite) TestGetWebInfo_WhenGetSuccess() {
	parsedJson, _ := json.Marshal(respHttp)
	httpResp := http.Response{
		Body:       io.NopCloser(bytes.NewBuffer(parsedJson)),
		StatusCode: 200,
	}
	suite.http.On("Get", ctxTest, fmt.Sprintf(constants.Ssllabs, "testurl.com")).
		Return(&httpResp, nil)

	_, err := suite.underTest.GetWebInfo(ctxTest, "testurl.com")
	suite.NoError(err)
}
