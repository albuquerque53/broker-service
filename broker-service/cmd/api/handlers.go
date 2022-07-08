package main

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

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
