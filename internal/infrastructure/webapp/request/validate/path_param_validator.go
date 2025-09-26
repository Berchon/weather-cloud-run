package validate

import (
	"net/http"

	"github.com/Berchon/weather-cloud-run/internal/business/model"
	"github.com/go-chi/chi/v5"
)

const (
	zipCodePathParam = "zipCode"
)

func ZipCode(r *http.Request) (*model.ZipCode, *model.CustomError) {
	zipCodeParam := chi.URLParam(r, zipCodePathParam)
	if len(zipCodeParam) == 0 {
		return nil, model.NewCustomError(http.StatusBadRequest, "Zip code path param is not precent")
	}
	zipCode, err := model.BuildZipCode(zipCodeParam)
	return zipCode, err
}
