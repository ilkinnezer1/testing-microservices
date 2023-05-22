package routes

import (
	"authentication/api"
	"authentication/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthRoutes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://*", "http://*"},
		AllowHeaders:     []string{echo.HeaderAccept, echo.HeaderContentType, echo.HeaderXCSRFToken, echo.HeaderAuthorization},
		AllowMethods:     []string{echo.PUT, echo.POST, echo.GET, echo.DELETE, echo.PATCH},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	e.Use(handlers.Heartbeat("/ping"))
	e.POST("/auth", api.Authenticate)

	return e
}
