package routes

import (
	"broker/api"
	"broker/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BaseRoutes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "https://*"},
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderContentType, echo.HeaderXCSRFToken, echo.HeaderAuthorization},
		AllowMethods:     []string{echo.PUT, echo.POST, echo.GET, echo.DELETE, echo.PATCH},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Check Connection
	e.Use(handlers.Heartbeat("ping"))
	e.GET("/broker", api.BrokerHandler)
	return e
}
