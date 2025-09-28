//go:generate mockery --dir=. --output=./mock --name=GetTemperatureByZipCodeUsecase --structname=MockGetTemperatureByZipCodeUsecase --outpkg=mock --filename=get_temperature_by_zip_code_usecase.go --disable-version-string
package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Berchon/weather-cloud-run/internal/business/gateway"
	"github.com/Berchon/weather-cloud-run/internal/business/model"
)

type GetTemperatureByZipCodeUsecase interface {
	GetTemperatureByZipCode(ctx context.Context, zipCode model.ZipCode) (*map[string]float64, *model.CustomError)
}

type getTemperatureByZipCodeUsecase struct {
	ViaCepService  gateway.ViaCepService
	WeatherService WeatherService
}

// Novo contrato p/ WeatherService
type WeatherService interface {
	GetWeatherByCity(ctx context.Context, city string) (*map[string]float64, *model.CustomError)
}

func NewGetTemperatureByZipCodeUsecase(viaCepService gateway.ViaCepService, weatherService WeatherService) GetTemperatureByZipCodeUsecase {
	return &getTemperatureByZipCodeUsecase{
		ViaCepService:  viaCepService,
		WeatherService: weatherService,
	}
}

func (uc *getTemperatureByZipCodeUsecase) GetTemperatureByZipCode(ctx context.Context, zipCode model.ZipCode) (*map[string]float64, *model.CustomError) {
	city, err := uc.ViaCepService.GetAddressByZipCode(ctx, zipCode)
	if err != nil {
		return nil, err
	}

	if *city == "" {
		return nil, model.NewCustomError(http.StatusInternalServerError, "city field is empty in response from via cep service")
	}

	// var city *string = new(string)
	// *city = "Porto Alegre"
	fmt.Println(*city)

	temps, err := uc.WeatherService.GetWeatherByCity(ctx, *city)
	if err != nil {
		return nil, err
	}

	return temps, nil
}
