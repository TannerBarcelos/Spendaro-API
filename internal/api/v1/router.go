package v1

import (
	"spendaro-api/internal/api/v1/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(routeGroup *echo.Group) {

	v1Router := routeGroup.Group("/v1")

	handlers.RegisterAuthRoutes(v1Router)
	handlers.RegisterBudgetRoutes(v1Router)

}
