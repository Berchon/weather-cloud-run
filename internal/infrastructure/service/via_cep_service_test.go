package service_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/Berchon/weather-cloud-run/internal/infrastructure/configs"
	servicepkg "github.com/Berchon/weather-cloud-run/internal/infrastructure/service"
	configMock "github.com/Berchon/weather-cloud-run/internal/infrastructure/service/config/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func configureEnvironment() {
	_ = configs.LoadConfig(".")
	configs.SetViaCepBaseUrl("https://viacep.com.br")
	configs.SetViaCepPath("/ws/%s/json")
}
func TestViaCepService_GetAddressByZipCode(t *testing.T) {
	ctx := context.Background()
	_ = configs.LoadConfig(".")
	configs.SetViaCepBaseUrl("https://viacep.com.br")
	configs.SetViaCepPath("/ws/%s/json")

	t.Run("Should return error when creating request", func(t *testing.T) {
		_ = configs.LoadConfig(".")
		configs.SetViaCepBaseUrl("http://[::1]:namedport") // invalid URL
		configs.SetViaCepPath("/ws/%s/json")
		mockClient := configMock.NewMockHTTPDoer(t)

		svc := servicepkg.NewViaCepService(mockClient)

		_, err := svc.GetAddressByZipCode(ctx, "12345-678")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error creating request")
	})

	t.Run("Should return error when sending request", func(t *testing.T) {
		configureEnvironment()
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(nil, errors.New("http failure"))

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "http failure")
		mockClient.AssertExpectations(t)
	})

	t.Run("Should return a 422 when via cep return status 400", func(t *testing.T) {
		configureEnvironment()
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(newResponse(http.StatusBadRequest, ""), nil)

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		assert.Equal(t, 422, err.StatusCode)
		mockClient.AssertExpectations(t)
	})

	t.Run("Sould return a 500 when via cep return status different from 200 or 400", func(t *testing.T) {
		configureEnvironment()
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(newResponse(http.StatusInternalServerError, ""), nil)

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		assert.Equal(t, 500, err.StatusCode)
		mockClient.AssertExpectations(t)
	})

	t.Run("Should return a 500 when error reading response", func(t *testing.T) {
		configureEnvironment()
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(&http.Response{
			StatusCode: 200,
			Body:       &errorReader{},
		}, nil)

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error reading response")
		mockClient.AssertExpectations(t)
	})

	t.Run("Should return a 500 when error unmarshalling response body", func(t *testing.T) {
		configureEnvironment()
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(newResponse(200, "{invalid"), nil)

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unmarshalling")
		mockClient.AssertExpectations(t)
	})

	t.Run("Should return a 404 when via cep return error", func(t *testing.T) {
		configureEnvironment()
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(newResponse(200, `{"erro":"true"}`), nil)

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		assert.Equal(t, 404, err.StatusCode)
		mockClient.AssertExpectations(t)
	})

	t.Run("Should return a 500 when via cep return empty city", func(t *testing.T) {
		configureEnvironment()
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(newResponse(200, `{"cep":"12345-678","localidade":""}`), nil)

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		assert.Equal(t, 500, err.StatusCode)
		mockClient.AssertExpectations(t)
	})

	t.Run("Should return a valid city with success", func(t *testing.T) {
		configureEnvironment()
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(newResponse(200, `{"cep":"12345-678","localidade":"Porto Alegre"}`), nil)

		svc := servicepkg.NewViaCepService(mockClient)
		city, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Nil(t, err)
		assert.NotNil(t, city)
		assert.Equal(t, "Porto Alegre", *city)
		mockClient.AssertExpectations(t)
	})
}

func newResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

// Struct usada para simular erro na leitura do body
type errorReader struct{}

func (e *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("erro simulado de leitura")
}
func (e *errorReader) Close() error { return nil }
