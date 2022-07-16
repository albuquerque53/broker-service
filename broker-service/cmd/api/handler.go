package main

import (
	"encoding/json"
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := JSONResponse{
		Error:   false,
		Message: "Hello, World!",
	}

	json, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(json))
}
