package handler

import (
	"broker/internal/application/tool"
	"net/http"
)

func Broker(w http.ResponseWriter, r *http.Request) {
	payload := tool.JSONResponse{
		Error:   false,
		Message: "Hello, World!",
	}

	_ = tool.WriteJSON(w, http.StatusOK, payload)
}
