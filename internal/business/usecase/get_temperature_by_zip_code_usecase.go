//go:generate mockery --dir=. --output=./mock --name=GetTemperatureByZipCodeUsecase --structname=MockGetTemperatureByZipCodeUsecase --outpkg=mock --filename=get_temperature_by_zip_code_usecase.go --disable-version-string
package usecase

import (
	"context"
	"fmt"
	"math"
	"net/http"

	"github.com/Berchon/weather-cloud-run/internal/business/gateway"
	"github.com/Berchon/weather-cloud-run/internal/business/model"
)

type GetTemperatureByZipCodeUsecase interface {
	GetTemperatureByZipCode(ctx context.Context, zipCode model.ZipCode) (*map[string]float64, *model.CustomError)
}

type getTemperatureByZipCodeUsecase struct {
	ViaCepService  gateway.ViaCepService
	WeatherService gateway.WeatherService
}

func NewGetTemperatureByZipCodeUsecase(viaCepService gateway.ViaCepService, weatherService gateway.WeatherService) GetTemperatureByZipCodeUsecase {
	return &getTemperatureByZipCodeUsecase{
		ViaCepService:  viaCepService,
		WeatherService: weatherService,
	}
}

func (uc *getTemperatureByZipCodeUsecase) GetTemperatureByZipCode(ctx context.Context, zipCode model.ZipCode) (*map[string]float64, *model.CustomError) {
	city, err := uc.ViaCepService.GetAddressByZipCode(ctx, zipCode)
	fmt.Println("Error: ", err)
	if err != nil {
		return nil, err
	}

	if *city == "" {
		return nil, model.NewCustomError(http.StatusInternalServerError, "city field is empty in response from via cep service")
	}

	celcius, err := uc.WeatherService.GetWeatherByCity(ctx, *city)
	if err != nil {
		return nil, err
	}

	result := map[string]float64{
		"temp_C": roundToDecimalPlaces(*celcius, 1),
		"temp_F": roundToDecimalPlaces(convertCelsiusToFahrenheit(*celcius), 1),
		"temp_K": roundToDecimalPlaces(convertCelsiusToKelvin(*celcius), 1),
	}

	return &result, nil
}

func roundToDecimalPlaces(value float64, decimalPlaces uint) float64 {
	factor := math.Pow(10, float64(decimalPlaces))
	return math.Round(value*factor) / factor
}

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func convertCelsiusToKelvin(celsius float64) float64 {
	return celsius + 273 // 273.15
}
