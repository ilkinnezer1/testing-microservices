package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func BrokerHandler(c echo.Context) error {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	return WriteJSON(c, http.StatusOK, payload)
}
