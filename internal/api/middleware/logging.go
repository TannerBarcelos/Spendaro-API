package middleware

import (
	"spendaro-api/pkg/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RequestLoggerMiddleware is a middleware function that logs all incoming requests to the server. It logs the URI, status, method, user-agent, and headers of the request.
func RequestLoggerMiddleware() echo.MiddlewareFunc {
	logger := util.GetLogger()
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			uri := c.Request().RequestURI
			status := c.Response().Status
			method := c.Request().Method
			userAgent := c.Request().UserAgent()
			headers := c.Request().Header
			protocol := c.Request().Proto
			requestIp := c.RealIP()
			logger.Info().
				Str("uri", uri).
				Int("status", status).
				Str("method", method).
				Str("user-agent", userAgent).
				Interface("headers", headers).
				Str("protocol", protocol).
				Str("request-ip", requestIp).
				Msg("Request received")
			return nil
		},
	})
}
