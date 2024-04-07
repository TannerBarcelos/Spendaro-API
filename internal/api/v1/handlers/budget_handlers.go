package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"spendaro-api/internal/api/v1/services"
	"spendaro-api/internal/api/v1/utils"
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
	data, err := services.GetAllBudgets()

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, http.StatusInternalServerError, "Failed to retrieve all budgets", err))
	}

	return c.JSON(http.StatusOK, utils.PrepareResponse(data, http.StatusOK, "Successfully retrieved all budgets", nil))
}

func createBudgetHandler(c echo.Context) error {
	data, err := services.CreateBudget()

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, http.StatusInternalServerError, "Failed to create budget", err))
	}

	return c.JSON(http.StatusCreated, utils.PrepareResponse(data, http.StatusCreated, "Successfully created budget", nil))
}

func getBudgetHandler(c echo.Context) error {

	budgetId := c.Param("id")

	if budgetId == "" {
		return c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, http.StatusBadRequest, "Invalid budget ID", nil))
	}

	data, err := services.GetBudget(budgetId)

	if err != nil {
		return c.JSON(http.StatusBadGateway, utils.PrepareResponse(nil, http.StatusInternalServerError, "Failed to retrieve budget", err))
	}

	if data == "" {
		return c.JSON(http.StatusNotFound, utils.PrepareResponse(nil, http.StatusNotFound, "Budget not found", nil))
	}

	return c.JSON(http.StatusOK, utils.PrepareResponse(data, http.StatusOK, "Successfully retrieved budget", nil))
}

func updateBudgetHandler(c echo.Context) error {

	budgetId := c.Param("id")

	if budgetId == "" {
		return c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, http.StatusBadRequest, "Invalid budget ID", nil))
	}

	data, err := services.UpdateBudget(budgetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, http.StatusInternalServerError, "Failed to update budget", err))
	}

	if data == "" {
		return c.JSON(http.StatusNotFound, utils.PrepareResponse(nil, http.StatusNotFound, "Budget not found", nil))
	}

	return c.JSON(http.StatusOK, utils.PrepareResponse(data, http.StatusOK, "Successfully updated budget", nil))
}

func deleteBudgetHandler(c echo.Context) error {

	budgetId := c.Param("id")

	if budgetId == "" {
		return c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, http.StatusBadRequest, "Invalid budget ID", nil))
	}

	data, err := services.DeleteBudget(budgetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.PrepareResponse(nil, http.StatusInternalServerError, "Failed to delete budget", err))
	}

	if data == "" {
		return c.JSON(http.StatusNotFound, utils.PrepareResponse(nil, http.StatusNotFound, "Budget not found", nil))
	}

	return c.JSON(http.StatusOK, utils.PrepareResponse(data, http.StatusOK, "Successfully deleted budget", nil))
}
