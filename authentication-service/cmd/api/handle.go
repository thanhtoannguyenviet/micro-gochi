package main

import (
	"errors"
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

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}
	app.writeJSON(w, http.StatusOK, payload)
}
func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	switch requestPayload.Action {
	case "auth":
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {
	//create some json we'll send to the auth microservice
}
