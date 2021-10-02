package main

import (
	"encoding/json"
	"net/http"
)

// statusHandler respon status of server
func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := appStatus{
		Status:     "OK",
		Enviroment: app.config.env,
		Version:    version,
	}

	// convert struct to json
	js, err := json.MarshalIndent(currentStatus, "", " ") //2 previs, 3 space
	if err != nil {
		app.logger.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") //set header
	w.WriteHeader(http.StatusOK)                       //send status 200
	w.Write(js)
}
