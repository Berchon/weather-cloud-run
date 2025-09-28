package usecase

import (
	"context"
	"math"
	"net/http"
	"testing"

	"github.com/Berchon/weather-cloud-run/internal/business/model"
	serviceMock "github.com/Berchon/weather-cloud-run/internal/infrastructure/service/mock"

	"github.com/stretchr/testify/assert"
)

func TestRoundToDecimalPlaces(t *testing.T) {
	t.Run("should round to 1 decimal place", func(t *testing.T) {
		assert.Equal(t, 2.3, roundToDecimalPlaces(2.34, 1))
		assert.Equal(t, 2.4, roundToDecimalPlaces(2.36, 1))
	})

	t.Run("should round to 2 decimal places", func(t *testing.T) {
		assert.Equal(t, 2.35, roundToDecimalPlaces(2.345, 2))
		assert.Equal(t, 2.35, roundToDecimalPlaces(2.346, 2))
	})

	t.Run("should round negative numbers", func(t *testing.T) {
		assert.Equal(t, -2.3, roundToDecimalPlaces(-2.34, 1))
		assert.Equal(t, -2.4, roundToDecimalPlaces(-2.36, 1))
	})

	t.Run("should handle zero decimal places", func(t *testing.T) {
		assert.Equal(t, 2.0, roundToDecimalPlaces(2.4, 0))
		assert.Equal(t, 3.0, roundToDecimalPlaces(2.5, 0))
	})

	t.Run("should handle large numbers", func(t *testing.T) {
		assert.Equal(t, 123456789.12, roundToDecimalPlaces(123456789.1234, 2))
	})

	t.Run("should handle very small numbers", func(t *testing.T) {
		assert.Equal(t, 0.0, roundToDecimalPlaces(0.00001, 2))
	})

	t.Run("should handle very large decimalPlaces", func(t *testing.T) {
		assert.Equal(t, 1.2345, roundToDecimalPlaces(1.2345, 100))
		assert.Equal(t, 1.0, roundToDecimalPlaces(1.2345, 0))
	})
}

func TestConvertCelsiusToFahrenheit(t *testing.T) {
	t.Run("should convert celsius to fahrenheit", func(t *testing.T) {
		assert.InDelta(t, 32.0, convertCelsiusToFahrenheit(0), 0.0001)
		assert.InDelta(t, 212.0, convertCelsiusToFahrenheit(100), 0.0001)
		assert.InDelta(t, -40.0, convertCelsiusToFahrenheit(-40), 0.0001)
	})

	t.Run("should handle decimal celsius", func(t *testing.T) {
		assert.InDelta(t, 98.6, convertCelsiusToFahrenheit(37), 0.0001)
	})
}

func TestConvertCelsiusToKelvin(t *testing.T) {
	t.Run("should convert celsius to kelvin", func(t *testing.T) {
		assert.InDelta(t, 273.0, convertCelsiusToKelvin(0), 0.0001)
		assert.InDelta(t, 373.0, convertCelsiusToKelvin(100), 0.0001)
		assert.InDelta(t, 233.0, convertCelsiusToKelvin(-40), 0.0001)
	})

	t.Run("should handle decimal celsius", func(t *testing.T) {
		assert.InDelta(t, 310.5, convertCelsiusToKelvin(37.5), 0.0001)
	})

	t.Run("should handle extreme values", func(t *testing.T) {
		assert.InDelta(t, math.MaxFloat64+273, convertCelsiusToKelvin(math.MaxFloat64), 1e292)
		assert.InDelta(t, math.SmallestNonzeroFloat64+273, convertCelsiusToKelvin(math.SmallestNonzeroFloat64), 1e-292)
	})
}

func TestNewGetTemperatureByZipCodeUsecase(t *testing.T) {
	t.Run("should create usecase with valid dependencies", func(t *testing.T) {
		viaCep := serviceMock.NewMockViaCepService(t)
		weather := serviceMock.NewMockWeatherService(t)
		uc := NewGetTemperatureByZipCodeUsecase(viaCep, weather)
		assert.NotNil(t, uc)
	})

	t.Run("should accept nil dependencies (not recommended, but possible)", func(t *testing.T) {
		uc := NewGetTemperatureByZipCodeUsecase(nil, nil)
		assert.NotNil(t, uc)
	})
}

func TestGetTemperatureByZipCode(t *testing.T) {
	zip := model.ZipCode("12345678")
	ctx := context.Background()
	viaCep := serviceMock.NewMockViaCepService(t)
	weather := serviceMock.NewMockWeatherService(t)

	t.Run("should return error if viaCepService returns error", func(t *testing.T) {
		viaCep.On("GetAddressByZipCode", ctx, zip).Return(nil, model.NewCustomError(http.StatusInternalServerError, "via cep error")).Once()
		uc := NewGetTemperatureByZipCodeUsecase(viaCep, weather)
		result, err := uc.GetTemperatureByZipCode(ctx, zip)
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "via cep error")
	})

	t.Run("should return error if city is empty", func(t *testing.T) {
		empty := ""
		viaCep.On("GetAddressByZipCode", ctx, zip).Return(&empty, nil).Once()
		uc := NewGetTemperatureByZipCodeUsecase(viaCep, weather)
		result, err := uc.GetTemperatureByZipCode(ctx, zip)
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "city field is empty")
	})

	t.Run("should return error if weatherService returns error", func(t *testing.T) {
		city := "Porto Alegre"
		viaCep.On("GetAddressByZipCode", ctx, zip).Return(&city, nil).Once()
		weather.On("GetWeatherByCity", ctx, city).Return(nil, model.NewCustomError(http.StatusInternalServerError, "weather error")).Once()
		uc := NewGetTemperatureByZipCodeUsecase(viaCep, weather)
		result, err := uc.GetTemperatureByZipCode(ctx, zip)
		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "weather error")
	})

	t.Run("should return temperatures when all services succeed", func(t *testing.T) {
		city := "Porto Alegre"
		viaCep.On("GetAddressByZipCode", ctx, zip).Return(&city, nil).Once()
		temp := 25.0
		weather.On("GetWeatherByCity", ctx, city).Return(&temp, nil).Once()
		uc := NewGetTemperatureByZipCodeUsecase(viaCep, weather)
		result, err := uc.GetTemperatureByZipCode(ctx, zip)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.InDelta(t, 25.0, (*result)["temp_C"], 0.01)
		assert.InDelta(t, 77.0, (*result)["temp_F"], 0.01)
		assert.InDelta(t, 298.0, (*result)["temp_K"], 0.01)
	})

	t.Run("should handle negative temperature", func(t *testing.T) {
		city := "Porto Alegre"
		temp := -10.0
		viaCep.On("GetAddressByZipCode", ctx, zip).Return(&city, nil).Once()
		weather.On("GetWeatherByCity", ctx, city).Return(&temp, nil).Once()
		uc := NewGetTemperatureByZipCodeUsecase(viaCep, weather)
		result, err := uc.GetTemperatureByZipCode(ctx, zip)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.InDelta(t, -10.0, (*result)["temp_C"], 0.01)
		assert.InDelta(t, 14.0, (*result)["temp_F"], 0.01)
		assert.InDelta(t, 263.0, (*result)["temp_K"], 0.01)
	})
}
