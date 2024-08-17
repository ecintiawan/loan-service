package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Health struct{}

func NewHealth() *Health {
	return &Health{}
}

func (h *Health) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
