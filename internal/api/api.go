package api

import (
	"fmt"
	"os"
	"spendaro-api/internal/api/middleware"
	v1 "spendaro-api/internal/api/v1"
	config "spendaro-api/pkg/util"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// server is a struct embedding an echo web server instance which will embed all properties and methods of the echo web server into the server struct to be used in the server instance.
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

// startServer starts the server on the port specified in the config file.
func (e *server) startServer() {
	addressFromConfig := config.GetConfigString("server.port")

	log.Log().Msgf("server started on port %s", addressFromConfig)
	err := e.Start(fmt.Sprintf(":%s", addressFromConfig))

	if err != nil {
		log.Error().Err(err).Msgf("Error starting server on port %s", addressFromConfig)
		os.Exit(1)
	}

}

// registerRoutes registers all the routes for the server and provides a versioned API.
func (e *server) registerRoutes() {
	root := e.Group("/api")
	v1.RegisterRoutes(root)

	/*
		Register future semver router versions here i.e. v2.RegisterRoutes(root), v3.RegisterRoutes(root), etc.
		This ensures the API is backwards compatible and can be versioned.
	*/

}

// registerMiddlewares registers all the middlewares for the server, such as logging, authentication, etc.
func (e *server) registerMiddlewares() {
	e.Use(middleware.Cors())
}
