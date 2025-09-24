package validate_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/Berchon/weather-cloud-run/internal/business/model"
	"github.com/Berchon/weather-cloud-run/internal/infra/webapp/request/validate"
)

func TestCep(t *testing.T) {
	t.Run("Should return an error When zip code is not present", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/cep/", nil)
		rctx := chi.NewRouteContext()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		cep, err := validate.Cep(r)

		assert.Nil(t, cep)
		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.StatusCode)
		assert.Contains(t, err.Error(), "CEP path param is not precent")
	})

	t.Run("Should return an error When zip code is invalid", func(t *testing.T) {
		stringCep := "abc"
		r := httptest.NewRequest(http.MethodGet, "/cep/"+stringCep, nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("cep", stringCep)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		cep, err := validate.Cep(r)

		assert.Nil(t, cep)
		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.StatusCode)
		assert.Contains(t, err.Error(), "Erro 400: CEP [abc] is not valid")
	})

	t.Run("Should return a zip code When zip code is valid", func(t *testing.T) {
		stringCep := "12345678"
		r := httptest.NewRequest(http.MethodGet, "/cep/"+stringCep, nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("cep", stringCep)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		cep, err := validate.Cep(r)

		expected, _ := model.BuildCEP(stringCep)
		assert.NotNil(t, cep)
		assert.Nil(t, err)
		assert.Equal(t, expected, cep)
	})
}
