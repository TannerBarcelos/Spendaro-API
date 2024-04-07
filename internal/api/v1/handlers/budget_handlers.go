package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterBudgetRoutes(v1 *echo.Group) {
	budgets := v1.Group("/budgets")
	{
		budgets.GET("", getBudgetsHandler)          // GET /v1/budgets
		budgets.GET("/:id", getBudgetHandler)       // GET /v1/budgets/:id
		budgets.POST("", createBudgetHandler)       // POST /v1/budgets
		budgets.PUT("/:id", updateBudgetHandler)    // PUT /v1/budgets/:id
		budgets.DELETE("/:id", deleteBudgetHandler) // DELETE /v1/budgets/:id
	}
}

func getBudgetsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetBudgets")
}

func createBudgetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "CreateBudget")
}

func getBudgetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "GetBudget")
}

func updateBudgetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "UpdateBudget")
}

func deleteBudgetHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "DeleteBudget")
}
