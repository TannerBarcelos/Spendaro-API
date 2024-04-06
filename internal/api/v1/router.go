package v1

import (
	"spendaro-api/internal/api/v1/healthcheck"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(routeGroup *echo.Group) {

	v1 := routeGroup.Group("/v1")

	// Register routes here
	v1.GET("/healthcheck", healthcheck.Healthcheck)

	authRoutes := v1.Group("/auth")
	authRoutes.POST("/login", nil)
	authRoutes.POST("/register", nil)

	budgetRoutes := v1.Group("/budgets")
	budgetRoutes.POST("/create", nil)
	budgetRoutes.GET("/:id", nil)
	budgetRoutes.PUT("/:id", nil)
	budgetRoutes.DELETE("/:id", nil)
}
