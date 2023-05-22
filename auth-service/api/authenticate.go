package api

import (
	"authentication/data"
	"authentication/handlers"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Authenticate(c echo.Context) error {
	var reqPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := handlers.ReadJSON(c, &reqPayload)

	if err != nil {
		err := handlers.ErrorJSON(c, err, http.StatusBadRequest)
		if err != nil {
			return err
		}
		return err
	}

	// Validate user against the database
	user, err := data.GetByEmail(reqPayload.Email)
	valid, err := user.PasswordMatches(reqPayload.Password)

	if err != nil || !valid {
		err := handlers.ErrorJSON(c, errors.New("invalid credentials"), http.StatusBadRequest)
		if err != nil {
			return err
		}
		return err
	}

	payload := handlers.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	return c.JSON(http.StatusAccepted, payload)
}
