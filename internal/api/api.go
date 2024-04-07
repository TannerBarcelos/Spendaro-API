package api

import (
	"fmt"
	v1 "spendaro-api/internal/api/v1"
	config "spendaro-api/pkg/util"

	"github.com/labstack/echo/v4"
)

// server is a struct embedding an echo web server instance
type server struct {
	*echo.Echo
}

// NewEchoServer creates a new echo web server instance and registers routes, middleware, etc. It then starts the server.
func NewEchoServer() {
	server := &server{echo.New()}
	server.registerMiddlewares()
	server.registerRoutes()
	server.startServer()
}

func (e *server) startServer() {
	addressFromConfig := config.GetConfigString("server.port")
	fmt.Println("Starting server on port", addressFromConfig)
	if err := e.Start(fmt.Sprintf(":%s", addressFromConfig)); err != nil {
		e.Logger.Fatal(err)
	}
}

// registerRoutes registers all the routes for the server
func (e *server) registerRoutes() {
	v1.RegisterRoutes(e.Group("/api"))
}

// registerMiddlewares registers all the middlewares for the server
func (e *server) registerMiddlewares() {
	// Register middlewares here
}
