package model

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func NewInternalServerError() error {
	return echo.NewHTTPError(http.StatusInternalServerError, ErrorResponse{
		Message: "Internal server error",
	})
}
