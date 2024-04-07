package v1

import (
	"spendaro-api/internal/api/v1/handlers"

	"github.com/labstack/echo/v4"
)

const APIVersion = "/v1"

func RegisterRoutes(routeGroup *echo.Group) {

	router := routeGroup.Group(APIVersion)

	handlers.RegisterHealthRoutes(router)
	handlers.RegisterAuthRoutes(router)
	handlers.RegisterBudgetRoutes(router)

}
