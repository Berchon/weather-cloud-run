package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Berchon/weather-cloud-run/internal/business/model"
	usecaseMock "github.com/Berchon/weather-cloud-run/internal/business/usecase/mock"
	"github.com/Berchon/weather-cloud-run/internal/infra/webapp/handler"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTestServer(t *testing.T, mockUsecase *usecaseMock.MockGetTemperatureByCepUsecase) *httptest.Server {
	h := handler.NewGetTemperatureByCepHandler(mockUsecase)
	r := chi.NewRouter()
	r.Get("/temperature/{cep}", h.Handle)
	server := httptest.NewServer(r)
	t.Cleanup(func() { server.Close() })
	return server
}

// helper para fazer GET e ler body
func getResponse(t *testing.T, url string) (int, string) {
	resp, err := http.Get(url)
	assert.NoError(t, err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(body)
}

func TestGetTemperatureByCepHandler_HTTP(t *testing.T) {
	mockUsecase := usecaseMock.NewMockGetTemperatureByCepUsecase(t)
	server := setupTestServer(t, mockUsecase)
	t.Run("should return 400 when CEP validation fails", func(t *testing.T) {

		status, body := getResponse(t, server.URL+"/temperature/12")
		assert.Equal(t, http.StatusBadRequest, status)
		assert.JSONEq(t, body, `{"status_code":400,"message":"CEP [12] is not valid"}`)
	})

	t.Run("should return 422 when usecase returns error", func(t *testing.T) {
		cep := "12345678"
		mockUsecase.On("GetTemperatureByCep", mock.Anything, model.CEP(cep)).
			Return(nil, model.NewCustomError(http.StatusUnprocessableEntity, "something went wrong")).
			Once()

		status, body := getResponse(t, server.URL+"/temperature/"+cep)
		assert.Equal(t, http.StatusUnprocessableEntity, status)
		assert.Contains(t, body, `{"status_code":422,"message":"something went wrong"}`)
	})

	t.Run("should return 200 when usecase succeeds", func(t *testing.T) {
		cep := "12345678"
		temp := "25.3°C"
		mockUsecase.On("GetTemperatureByCep", mock.Anything, model.CEP(cep)).
			Return(&temp, nil).
			Once()

		status, body := getResponse(t, server.URL+"/temperature/"+cep)
		assert.Equal(t, http.StatusOK, status)
		assert.JSONEq(t, body, `"25.3°C"`)
	})
}
