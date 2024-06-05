package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHealthRoutes(baseRouter *echo.Group) {
	baseRouter.GET("", healthcheckHandler)

}

func healthcheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Server is healthy!")
}
