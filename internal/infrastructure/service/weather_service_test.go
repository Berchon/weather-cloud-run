package service_test

import (
	"context"
	"encoding/json"
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

func configureWeatherEnvironment() {
	_ = configs.LoadConfig(".")
	configs.SetWeatherBaseUrl("https://api.weather.com")
	configs.SetWeatherPath("/v1/current.json")
	configs.SetWeatherAPIKey("test-key")
}

func TestNewWeatherService_DefaultClient(t *testing.T) {
	service := servicepkg.NewWeatherService(nil)
	assert.NotNil(t, service)
}

func TestWeatherService_GetWeatherByCity(t *testing.T) {
	ctx := context.Background()
	configureWeatherEnvironment()

	t.Run("should return error when building endpoint fails", func(t *testing.T) {
		configs.SetWeatherBaseUrl("://invalid-url")
		mockClient := configMock.NewMockHTTPDoer(t)
		svc := servicepkg.NewWeatherService(mockClient)
		_, err := svc.GetWeatherByCity(ctx, "Porto Alegre")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid baseURL")
	})

	t.Run("should return error when sending request fails", func(t *testing.T) {
		configureWeatherEnvironment()
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(nil, errors.New("http failure"))
		svc := servicepkg.NewWeatherService(mockClient)
		_, err := svc.GetWeatherByCity(ctx, "Porto Alegre")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "http failure")
		mockClient.AssertExpectations(t)
	})

	t.Run("should return error when reading response body fails", func(t *testing.T) {
		configureWeatherEnvironment()
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(&http.Response{StatusCode: 200, Body: &config.ErrorReader{}}, nil)
		svc := servicepkg.NewWeatherService(mockClient)
		_, err := svc.GetWeatherByCity(ctx, "Porto Alegre")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error reading response")
		mockClient.AssertExpectations(t)
	})

	t.Run("should return error when weather api returns error json", func(t *testing.T) {
		configureWeatherEnvironment()
		apiErr := `{"error":{"code":1006,"message":"No matching location found."}}`
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(config.NewTestResponse(400, apiErr), nil)
		svc := servicepkg.NewWeatherService(mockClient)
		_, err := svc.GetWeatherByCity(ctx, "InvalidCity")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "No matching location found.")
		mockClient.AssertExpectations(t)
	})

	t.Run("should return error when weather api returns unexpected error", func(t *testing.T) {
		configureWeatherEnvironment()
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(config.NewTestResponse(500, "unexpected error"), nil)
		svc := servicepkg.NewWeatherService(mockClient)
		_, err := svc.GetWeatherByCity(ctx, "Porto Alegre")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected error from weather api")
		mockClient.AssertExpectations(t)
	})

	t.Run("should return error when unmarshalling response body fails", func(t *testing.T) {
		configureWeatherEnvironment()
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(config.NewTestResponse(200, "{invalid"), nil)
		svc := servicepkg.NewWeatherService(mockClient)
		_, err := svc.GetWeatherByCity(ctx, "Porto Alegre")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unmarshalling")
		mockClient.AssertExpectations(t)
	})

	t.Run("should return valid temperatures when weather api returns success", func(t *testing.T) {
		configureWeatherEnvironment()
		weatherResp := map[string]interface{}{
			"location": map[string]interface{}{"name": "Porto Alegre"},
			"current":  map[string]interface{}{"temp_c": 25.0},
		}
		body, _ := json.Marshal(weatherResp)
		mockClient := configMock.NewMockHTTPDoer(t)
		mockClient.On("Do", mock.Anything).Return(config.NewTestResponse(200, string(body)), nil)
		svc := servicepkg.NewWeatherService(mockClient)
		result, err := svc.GetWeatherByCity(ctx, "Porto Alegre")
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.InDelta(t, 25.0, (*result)["temp_C"], 0.01)
		assert.InDelta(t, 77.0, (*result)["temp_F"], 0.01)
		assert.InDelta(t, 298.0, (*result)["temp_K"], 0.01)
		mockClient.AssertExpectations(t)
	})
}
