package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheckResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HealthCheck godoc
//
//	@Tags			Health
//	@Summary		Check if service is active
//	@Description	health service
//	@Produce		json
//	@Success		200	{object}	HealthCheckResponse
//	@Router			/health [get]
func HealthCheck(c echo.Context) error {
	response := HealthCheckResponse{
		Status:  "OK",
		Message: "Server is running",
	}

	return c.JSON(http.StatusOK, response)
}
