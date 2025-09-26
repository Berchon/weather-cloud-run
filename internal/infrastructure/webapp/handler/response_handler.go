package handler

import (
	"encoding/json"
	"net/http"
)

type responseHandler struct{}

func NewResponseHandler() *responseHandler {
	return &responseHandler{}
}

func (responseHandler *responseHandler) RequestResponse(w http.ResponseWriter, r *http.Request, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}
