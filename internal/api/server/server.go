package server

import (
	"fmt"
	h "spendaro-api/internal/api/v1/healthcheck"

	"github.com/labstack/echo/v4"
)

// server is a struct embedding an echo web server instance
type server struct {
	*echo.Echo
}

// NewEchoServer creates a new echo web server instance and registers routes, middleware, etc. It then starts the server.
func NewEchoServer(addr string) *server {
	e := &server{echo.New()}
	e.registerMiddlewares()
	e.registerRoutes()
	e.startServer(addr)
	return e
}

// startServer starts the server on the given address
func (e *server) startServer(addr string) {
	if err := e.Start(addr); err != nil {
		e.Logger.Fatal(err)
	}
}

// registerRoutes registers all the routes for the server
func (e *server) registerRoutes() {

	apiVersion := "v1" // TODO: Move this to a config file and load it from there into this variable

	routes := e.Group(fmt.Sprintf("/api/%s", apiVersion))

	healthRoutes := routes.Group("/healthcheck")
	healthRoutes.GET("/", h.Healthcheck)

	authRoutes := routes.Group("/auth")
	authRoutes.POST("/login", nil)
	authRoutes.POST("/register", nil)

	budgetRoutes := routes.Group("/budget")
	budgetRoutes.POST("/create", nil)
	budgetRoutes.GET("/get/:id", nil)
	budgetRoutes.PUT("/update/:id", nil)
	budgetRoutes.DELETE("/delete/:id", nil)

}

// registerMiddlewares registers all the middlewares for the server
func (e *server) registerMiddlewares() {
	// Register middlewares here
}
