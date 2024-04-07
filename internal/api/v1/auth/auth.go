package auth

import "github.com/labstack/echo/v4"

func RegisterAuthRoutes(v1 *echo.Group) {
	auth := v1.Group("/auth")
	{
		auth.POST("/login", loginHandler)       // POST /v1/auth/login
		auth.POST("/register", registerHandler) // POST /v1/auth/register
	}
}
