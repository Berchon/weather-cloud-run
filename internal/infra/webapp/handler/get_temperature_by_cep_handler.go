package handler

import (
	"net/http"

	"github.com/Berchon/weather-cloud-run/internal/business/usecase"
	"github.com/Berchon/weather-cloud-run/internal/infra/webapp/request/validate"
)

type GetTemperatureByCepHandler interface {
	HttpHandler
}

type getTemperatureByCepHandler struct {
	usecase  usecase.GetTemperatureByCepUsecase
	response *responseHandler
}

func NewGetTemperatureByCepHandler(usecase usecase.GetTemperatureByCepUsecase) GetTemperatureByCepHandler {
	response := NewResponseHandler()
	return &getTemperatureByCepHandler{
		usecase:  usecase,
		response: response,
	}
}

func (h *getTemperatureByCepHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cep, err := validate.Cep(r)
	if err != nil {
		h.response.RequestResponse(w, r, err, err.StatusCode)
		return
	}

	output, err := h.usecase.GetTemperatureByCep(r.Context(), *cep)
	if err != nil {
		h.response.RequestResponse(w, r, err, http.StatusUnprocessableEntity)
		return
	}

	h.response.RequestResponse(w, r, output, http.StatusOK)
}
