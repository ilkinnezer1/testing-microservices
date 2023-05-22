package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func Heartbeat(endpoint string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if r := c.Request(); r.Method == "GET" && strings.EqualFold(r.URL.Path, endpoint) {
				return c.String(http.StatusOK, ".")
			}
			return next(c)
		}
	}
}
