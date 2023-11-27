package handlers_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"franchises-system/internal/infra/api/handlers"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	healthJson = `{"status":"OK","message":"Server is running"}`
	ctxTest    = context.Background()
)

type Controller struct {
	Req     *http.Request
	Res     *httptest.ResponseRecorder
	context echo.Context
}

func SetupController(method, endpoint string, body io.Reader) Controller {
	echoInstance := echo.New()
	req := httptest.NewRequest(method, endpoint, body)
	res := httptest.NewRecorder()
	ctx := echoInstance.NewContext(req, res)
	return Controller{
		req,
		res,
		ctx,
	}
}

func TestHeatlhCheck(t *testing.T) {
	controller := SetupController(http.MethodGet, "/api/health", nil)
	if assert.NoError(t, handlers.HealthCheck(controller.context)) {
		assert.Equal(t, http.StatusOK, controller.Res.Code)
		assert.Equal(t, healthJson, strings.TrimSpace(controller.Res.Body.String()))
	}
}
