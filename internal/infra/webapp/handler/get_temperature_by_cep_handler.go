package handler

import (
	"net/http"
)

type GetTemperatureByCepHandler interface {
	HttpHandler
}

type getTemperatureByCepHandler struct {
	// usecase  usecase.GetTemperatureByCepUsecase
	response *responseHandler
}

func NewGetTemperatureByCepHandler() GetTemperatureByCepHandler { //usecase usecase.GetTemperatureByCepUsecase) GetTemperatureByCepHandler {
	response := NewResponseHandler()
	return &getTemperatureByCepHandler{
		// usecase:  usecase,
		response: response,
	}
}

func (getTemperatureByCepHandler *getTemperatureByCepHandler) Handle(w http.ResponseWriter, r *http.Request) {

	getTemperatureByCepHandler.response.RequestResponse(w, r, nil, http.StatusOK)
}
