package main

import "net/http"

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}
	var reqPayload mailMessage
	err := app.readJSON(w, r, &reqPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	msg := Message{
		From:    reqPayload.From,
		To:      reqPayload.To,
		Subject: reqPayload.Subject,
		Data:    reqPayload.Message,
	}
	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	payload := jsonResponse{
		Error:   false,
		Message: "send to" + reqPayload.To,
	}
	app.writeJSON(w, http.StatusAccepted, payload)
}
