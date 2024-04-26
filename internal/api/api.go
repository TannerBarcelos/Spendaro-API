package api

import (
	"fmt"
	"os"
	v1 "spendaro-api/internal/api/v1"
	"spendaro-api/pkg/middleware"
	config "spendaro-api/pkg/util"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

type server struct {
	*echo.Echo
}

func RunServer() {
	server := &server{
		echo.New(),
	}
	server.registerMiddlewares()
	server.registerRoutes()
	server.startServer()
}

// startServer starts the server on the port specified in the config file.
func (e *server) startServer() {
	addressFromConfig := config.ReadConfigValue("server.port")

	if err := e.Start(fmt.Sprintf(":%s", addressFromConfig)); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
		os.Exit(1)
	}

}

// registerRoutes registers all the routes for the server and provides a versioned API.
func (e *server) registerRoutes() {
	v1.RegisterRoutes(e.Group("/api"))

	/*
		Register future semver router versions here i.e. v2.RegisterRoutes(root), v3.RegisterRoutes(root), etc.
		This ensures the API is backwards compatible and can be versioned.
	*/

}

// registerMiddlewares registers all the middlewares for the server, such as logging, authentication, etc.
func (e *server) registerMiddlewares() {
	e.Use(middleware.CorsMiddleware())
	e.Use(middleware.RequestLoggerMiddleware())
}
