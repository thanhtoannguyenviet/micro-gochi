package main

import (
	"net/http"
)

// Broker is a test handler, just to make sure we can hit the broker from a web client
func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
