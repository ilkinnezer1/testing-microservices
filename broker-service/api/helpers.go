package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

type JsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:",omitempty"`
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadJSON(c echo.Context, data interface{}) error {
	// Limitation of file size
	maxBytes := 1048576 // 1 MB

	body, err := io.ReadAll(io.LimitReader(c.Request().Body, int64(maxBytes)))
	HandleError(err)

	dec := json.NewDecoder(io.NopCloser(bytes.NewBuffer(body)))
	err = dec.Decode(data)

	_, err = dec.Token()
	if err != io.EOF {
		return errors.New("body must have a single JSON value")
	}

	return nil

}

func WriteJSON(c echo.Context, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	HandleError(err)

	// If any headers included as a parameter
	if len(headers) > 0 {
		for key, value := range headers[0] {
			c.Request().Header[key] = value
		}
	}

	c.Request().Header.Set("Content-Type", "application/json")
	c.Response().Writer.WriteHeader(status)

	_, err = c.Response().Writer.Write(out)

	return nil
}

func ErrorJSON(c echo.Context, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return c.JSON(statusCode, payload)
}
