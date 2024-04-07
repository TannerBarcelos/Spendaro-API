package v1

import (
	"spendaro-api/internal/api/v1/auth"
	"spendaro-api/internal/api/v1/budget"
	"spendaro-api/internal/api/v1/healthcheck"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(routeGroup *echo.Group) {

	router := routeGroup.Group("/v1")

	healthcheck.RegisterHealthRoutes(router)
	auth.RegisterAuthRoutes(router)
	budget.RegisterBudgetRoutes(router)

}
