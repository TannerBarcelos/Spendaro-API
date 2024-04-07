package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"spendaro-api/internal/api/v1/services"
	httputils "spendaro-api/internal/api/v1/utils"
)

func RegisterBudgetRoutes(v1 *echo.Group) {
	budgets := v1.Group("/budgets") // budgets resource ; /v1/budgets
	{
		budgets.GET("", getBudgetsHandler)                                         // GET /v1/budgets
		budgets.GET("/:id", getBudgetHandler)                                      // GET /v1/budgets/:id
		budgets.POST("/create-budget", createBudgetHandler)                        // POST /v1/budgets
		budgets.POST("/category/create-category", createBudgetCategoryHandler)     // POST /v1/budgets/category
		budgets.POST("/category/:id/create-item", createBudgetCategoryItemHandler) // POST /v1/budgets/item
		budgets.PUT("/update/:id", updateBudgetHandler)                            // PUT /v1/budgets/:id
		budgets.DELETE("/delete/:id", deleteBudgetHandler)                         // DELETE /v1/budgets/:id
	}
}

func getBudgetsHandler(c echo.Context) error {
	data, err := services.GetAllBudgets()

	if err != nil {
		return c.JSON(http.StatusBadRequest, httputils.PrepareResponse(nil, http.StatusInternalServerError, "Failed to retrieve all budgets", err))
	}

	return c.JSON(http.StatusOK, httputils.PrepareResponse(data, http.StatusOK, "Successfully retrieved all budgets", nil))
}

func createBudgetHandler(c echo.Context) error {
	data, err := services.CreateBudget()

	if err != nil {
		return c.JSON(http.StatusBadRequest, httputils.PrepareResponse(nil, http.StatusInternalServerError, "Failed to create budget", err))
	}

	return c.JSON(http.StatusCreated, httputils.PrepareResponse(data, http.StatusCreated, "Successfully created budget", nil))
}

func createBudgetCategoryHandler(c echo.Context) error {
	data, err := services.CreateBudgetCategory()

	if err != nil {
		return c.JSON(http.StatusBadRequest, httputils.PrepareResponse(nil, http.StatusInternalServerError, "Failed to create budget category", err))
	}

	return c.JSON(http.StatusCreated, httputils.PrepareResponse(data, http.StatusCreated, "Successfully created budget category", nil))
}

func createBudgetCategoryItemHandler(c echo.Context) error {

	categoryId := c.Param("id")

	if categoryId == "" {
		return c.JSON(http.StatusBadRequest, httputils.PrepareResponse(nil, http.StatusBadRequest, "Invalid category ID", nil))
	}

	data, err := services.CreateBudgetCategoryItem(categoryId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, httputils.PrepareResponse(nil, http.StatusInternalServerError, "Failed to create budget category item", err))
	}

	if data == "" {
		return c.JSON(http.StatusNotFound, httputils.PrepareResponse(nil, http.StatusNotFound, "Category not found", nil))
	}

	return c.JSON(http.StatusCreated, httputils.PrepareResponse(data, http.StatusCreated, "Successfully created budget category item", nil))
}

func getBudgetHandler(c echo.Context) error {

	budgetId := c.Param("id")

	if budgetId == "" {
		return c.JSON(http.StatusBadRequest, httputils.PrepareResponse(nil, http.StatusBadRequest, "Invalid budget ID", nil))
	}

	data, err := services.GetBudget(budgetId)

	if err != nil {
		return c.JSON(http.StatusBadGateway, httputils.PrepareResponse(nil, http.StatusInternalServerError, "Failed to retrieve budget", err))
	}

	if data == "" {
		return c.JSON(http.StatusNotFound, httputils.PrepareResponse(nil, http.StatusNotFound, "Budget not found", nil))
	}

	return c.JSON(http.StatusOK, httputils.PrepareResponse(data, http.StatusOK, "Successfully retrieved budget", nil))
}

func updateBudgetHandler(c echo.Context) error {

	budgetId := c.Param("id")

	if budgetId == "" {
		return c.JSON(http.StatusBadRequest, httputils.PrepareResponse(nil, http.StatusBadRequest, "Invalid budget ID", nil))
	}

	data, err := services.UpdateBudget(budgetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, httputils.PrepareResponse(nil, http.StatusInternalServerError, "Failed to update budget", err))
	}

	if data == "" {
		return c.JSON(http.StatusNotFound, httputils.PrepareResponse(nil, http.StatusNotFound, "Budget not found", nil))
	}

	return c.JSON(http.StatusOK, httputils.PrepareResponse(data, http.StatusOK, "Successfully updated budget", nil))
}

func deleteBudgetHandler(c echo.Context) error {

	budgetId := c.Param("id")

	if budgetId == "" {
		return c.JSON(http.StatusBadRequest, httputils.PrepareResponse(nil, http.StatusBadRequest, "Invalid budget ID", nil))
	}

	data, err := services.DeleteBudget(budgetId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, httputils.PrepareResponse(nil, http.StatusInternalServerError, "Failed to delete budget", err))
	}

	if data == "" {
		return c.JSON(http.StatusNotFound, httputils.PrepareResponse(nil, http.StatusNotFound, "Budget not found", nil))
	}

	return c.JSON(http.StatusOK, httputils.PrepareResponse(data, http.StatusOK, "Successfully deleted budget", nil))
}
