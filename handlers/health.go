package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthResponse struct {
	Message string `json:"message"`
}

// Health - Health Check Handler
func Health(c echo.Context) error {
	resp := HealthResponse{
		Message: "Healthy!",
	}
	return c.JSON(http.StatusOK, resp)
}
