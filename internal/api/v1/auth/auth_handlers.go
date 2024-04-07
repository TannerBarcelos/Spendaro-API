package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func loginHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Login")
}

func registerHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Register")
}
