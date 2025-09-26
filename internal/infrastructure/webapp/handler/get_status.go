package handler

import (
	"net/http"
)

type GetStatusHandler interface {
	HttpHandler
}

type getStatusHandler struct {
	response *responseHandler
}

type status struct {
	Status string `json:"status"`
}

func NewGetStatusHandler() GetStatusHandler {
	response := NewResponseHandler()
	return &getStatusHandler{
		response: response,
	}
}

func (getStatusHandler *getStatusHandler) Handle(w http.ResponseWriter, r *http.Request) {
	status := status{Status: "Healthy"}
	getStatusHandler.response.RequestResponse(w, r, status, http.StatusOK)
}
