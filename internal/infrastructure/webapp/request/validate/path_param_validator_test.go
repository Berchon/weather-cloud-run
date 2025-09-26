package validate_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/Berchon/weather-cloud-run/internal/business/model"
	"github.com/Berchon/weather-cloud-run/internal/infrastructure/webapp/request/validate"
)

func TestZipCode(t *testing.T) {
	t.Run("Should return an error When zip code is not present", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, "/zipCode/", nil)
		rctx := chi.NewRouteContext()
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		zipCode, err := validate.ZipCode(r)

		assert.Nil(t, zipCode)
		assert.NotNil(t, err)
		assert.Equal(t, http.StatusBadRequest, err.StatusCode)
		assert.Contains(t, err.Error(), "Zip code path param is not precent")
	})

	t.Run("Should return an error When zip code is invalid", func(t *testing.T) {
		stringZipCode := "abc"
		r := httptest.NewRequest(http.MethodGet, "/zipCode/"+stringZipCode, nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("zipCode", stringZipCode)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		zipCode, err := validate.ZipCode(r)

		assert.Nil(t, zipCode)
		assert.NotNil(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, err.StatusCode)
		assert.Contains(t, err.Error(), "invalid zipcode")
	})

	t.Run("Should return a zip code When zip code is valid", func(t *testing.T) {
		stringZipCode := "12345678"
		r := httptest.NewRequest(http.MethodGet, "/zipCode/"+stringZipCode, nil)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("zipCode", stringZipCode)
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		zipCode, err := validate.ZipCode(r)

		expected, _ := model.BuildZipCode(stringZipCode)
		assert.NotNil(t, zipCode)
		assert.Nil(t, err)
		assert.Equal(t, expected, zipCode)
	})
}
