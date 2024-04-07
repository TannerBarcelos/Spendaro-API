package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHealthRoutes(v1 *echo.Group) {
	healthcheck := v1.Group("/healthcheck")
	{
		healthcheck.GET("", healthcheckHandler) // GET /v1/healthcheck
	}
}

func healthcheckHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Server is healthy!")
}
