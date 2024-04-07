package budget

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

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
