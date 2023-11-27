package router_test

import (
	"net/http"
	"testing"

	"franchises-system/internal/infra/api/router"
	"franchises-system/internal/infra/api/router/groups"
	mocks "franchises-system/mocks/infra/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

var (
	paths = []string{
		"/api/health",
		"/api/swagger/*",
		"/api/franchises/",
		"/api/franchises/details/:name",
		"/api/franchises/company/:company_id",
		"/api/owners/",
		"/api/companies/",
	}
	methods = []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodPatch,
		http.MethodHead,
		http.MethodOptions,
		http.MethodConnect,
	}
)

type routerTestSuite struct {
	suite.Suite
	server    *echo.Echo
	franchise groups.FranchisesGroup
	owner     groups.OwnersGroup
	company   groups.CompaniesGroup

	underTest *router.Router
}

func TestRouterSuite(t *testing.T) {
	suite.Run(t, new(routerTestSuite))
}

func (suite *routerTestSuite) SetupSuite() {
	suite.server = echo.New()
	suite.franchise = groups.NewFranchisesGroup(&mocks.Franchises{})
	suite.owner = groups.NewOwnersGroup(&mocks.Owners{})
	suite.company = groups.NewCompaniesGroup(&mocks.Companies{})
	suite.underTest = router.NewRouter(
		suite.server,
		suite.franchise,
		suite.owner,
		suite.company,
	)
}

func (suite *routerTestSuite) TestInit() {
	suite.underTest.Init()

	for _, route := range suite.server.Routes() {
		suite.Contains(paths, route.Path)
		suite.Contains(methods, route.Method)
	}
}
