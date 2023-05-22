package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SubmissionHandler(c echo.Context) error {
	var requestPayload RequestPayload

	err := ReadJSON(c, &requestPayload)

	if err != nil {
		err := ErrorJSON(c, err, http.StatusInternalServerError)
		if err != nil {
			return err
		}
	}

	switch requestPayload.Action {
	case "auth":
		Authenticate(c, requestPayload.Auth)
	default:
		ErrorJSON(c, errors.New("unknown Action"))

	}

	return c.JSON(http.StatusOK, requestPayload.Auth)
}

func Authenticate(c echo.Context, a AuthPayload) error {
	// create json to send the auth microservice
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	// call the service
	request, err := http.NewRequest("POST", "http://auth-service/authenticate", bytes.NewBuffer(jsonData))

	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(request)
	defer res.Body.Close()
	// get back correct status code

	if res.StatusCode == http.StatusUnauthorized {
		ErrorJSON(c, errors.New("invalid credentials"))
		return err
	} else if res.StatusCode == http.StatusAccepted {
		ErrorJSON(c, errors.New("error calling auth service"))
		return err
	}

	var jsonFromAuthService JsonResponse

	err = json.NewDecoder(res.Body).Decode(&jsonFromAuthService)

	if jsonFromAuthService.Error {
		ErrorJSON(c, err, http.StatusUnauthorized)
	}

	var payload JsonResponse
	payload.Error = false
	payload.Message = "Authenticated"
	payload.Data = jsonFromAuthService.Data

	return c.JSON(http.StatusAccepted, payload)
}