package handler

import (
	"net/http"

	"github.com/Berchon/weather-cloud-run/internal/business/usecase"
	"github.com/Berchon/weather-cloud-run/internal/infrastructure/webapp/request/validate"
)

type GetTemperatureByZipCodeHandler interface {
	HttpHandler
}

type getTemperatureByZipCodeHandler struct {
	usecase  usecase.GetTemperatureByZipCodeUsecase
	response *responseHandler
}

func NewGetTemperatureByZipCodeHandler(usecase usecase.GetTemperatureByZipCodeUsecase) GetTemperatureByZipCodeHandler {
	response := NewResponseHandler()
	return &getTemperatureByZipCodeHandler{
		usecase:  usecase,
		response: response,
	}
}

func (h *getTemperatureByZipCodeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	zipCode, err := validate.ZipCode(r)
	if err != nil {
		h.response.RequestResponse(w, r, err, err.StatusCode)
		return
	}

	output, err := h.usecase.GetTemperatureByZipCode(r.Context(), *zipCode)
	if err != nil {
		h.response.RequestResponse(w, r, err, err.StatusCode)
		return
	}

	h.response.RequestResponse(w, r, output, http.StatusOK)
}
