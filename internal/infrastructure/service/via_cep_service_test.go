package service_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/Berchon/weather-cloud-run/internal/infrastructure/configs"
	servicepkg "github.com/Berchon/weather-cloud-run/internal/infrastructure/service"
	config "github.com/Berchon/weather-cloud-run/internal/infrastructure/service/config"
	configMock "github.com/Berchon/weather-cloud-run/internal/infrastructure/service/config/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func configureEnvironment() {
	_ = configs.LoadConfig(".")
	configs.SetViaCepBaseUrl("https://viacep.com.br")
	configs.SetViaCepPath("/ws/%s/json")
}

func TestNewViaCepService_DefaultClient(t *testing.T) {
	service := servicepkg.NewViaCepService(nil)
	assert.NotNil(t, service)
}
func TestViaCepService_GetAddressByZipCode(t *testing.T) {

	ctx := context.Background()
	configureEnvironment()

	t.Run("should return error when creating request with invalid URL", func(t *testing.T) {
		originalBaseUrl := configs.GetViaCepBaseUrl()
		configs.SetViaCepBaseUrl("http://[::1]:namedport") // invalid URL
		defer configs.SetViaCepBaseUrl(originalBaseUrl)

		mockClient := configMock.NewMockHTTPDoer(t)
		svc := servicepkg.NewViaCepService(mockClient)

		_, err := svc.GetAddressByZipCode(ctx, "12345-678")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid baseURL")
	})

	t.Run("should return error when sending request fails", func(t *testing.T) {
		configs.SetViaCepBaseUrl("https://viacep.com.br")
		configs.SetViaCepPath("/ws/%s/json")

		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(nil, errors.New("http failure"))

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "http failure")
		mockClient.AssertExpectations(t)
	})

	t.Run("should return 422 when via cep returns status 400", func(t *testing.T) {
		configs.SetViaCepBaseUrl("https://viacep.com.br")
		configs.SetViaCepPath("/ws/%s/json")

		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(config.NewTestResponse(http.StatusBadRequest, ""), nil)

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		if assert.NotNil(t, err) {
			assert.Equal(t, 422, err.StatusCode)
		}
		mockClient.AssertExpectations(t)
	})

	t.Run("should return 500 when via cep returns status different from 200 or 400", func(t *testing.T) {
		configs.SetViaCepBaseUrl("https://viacep.com.br")
		configs.SetViaCepPath("/ws/%s/json")

		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(config.NewTestResponse(http.StatusInternalServerError, ""), nil)

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		if assert.NotNil(t, err) {
			assert.Equal(t, 500, err.StatusCode)
		}
		mockClient.AssertExpectations(t)
	})

	t.Run("should return 500 when reading response body fails", func(t *testing.T) {
		configs.SetViaCepBaseUrl("https://viacep.com.br")
		configs.SetViaCepPath("/ws/%s/json")

		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(&http.Response{
			StatusCode: 200,
			Body:       &config.ErrorReader{},
		}, nil)

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error reading response")
		mockClient.AssertExpectations(t)
	})

	t.Run("should return 500 when unmarshalling response body fails", func(t *testing.T) {
		configs.SetViaCepBaseUrl("https://viacep.com.br")
		configs.SetViaCepPath("/ws/%s/json")

		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(config.NewTestResponse(200, "{invalid"), nil)

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unmarshalling")
		mockClient.AssertExpectations(t)
	})

	t.Run("should return 404 when via cep returns error field", func(t *testing.T) {
		configs.SetViaCepBaseUrl("https://viacep.com.br")
		configs.SetViaCepPath("/ws/%s/json")

		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(config.NewTestResponse(200, `{"erro":"true"}`), nil)

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		if assert.NotNil(t, err) {
			assert.Equal(t, 404, err.StatusCode)
		}
		mockClient.AssertExpectations(t)
	})

	t.Run("should return 500 when city is empty in response", func(t *testing.T) {
		configs.SetViaCepBaseUrl("https://viacep.com.br")
		configs.SetViaCepPath("/ws/%s/json")

		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(config.NewTestResponse(200, `{"cep":"12345-678","localidade":""}`), nil)

		svc := servicepkg.NewViaCepService(mockClient)
		_, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Error(t, err)
		if assert.NotNil(t, err) {
			assert.Equal(t, 500, err.StatusCode)
		}
		mockClient.AssertExpectations(t)
	})

	t.Run("should return valid city when via cep returns success", func(t *testing.T) {
		configs.SetViaCepBaseUrl("https://viacep.com.br")
		configs.SetViaCepPath("/ws/%s/json")

		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(config.NewTestResponse(200, `{"cep":"12345-678","localidade":"Porto Alegre"}`), nil)

		svc := servicepkg.NewViaCepService(mockClient)
		city, err := svc.GetAddressByZipCode(ctx, "12345-678")

		assert.Nil(t, err)
		assert.NotNil(t, city)
		assert.Equal(t, "Porto Alegre", *city)
		mockClient.AssertExpectations(t)
	})
}
