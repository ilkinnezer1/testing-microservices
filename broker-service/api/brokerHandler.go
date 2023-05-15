package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type JsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:",omitempty"`
}

func BrokerHandler(c echo.Context) error {
	payload := JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	//output, _ := json.MarshalIndent(payload, "", "\t")
	//fmt.Println(output)
	return c.JSON(http.StatusOK, payload)
}
