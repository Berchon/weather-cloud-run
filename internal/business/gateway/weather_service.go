//go:generate mockery --dir=. --output=../../infrastructure/service/mock --name=WeatherService --structname=MockWeatherService --outpkg=mock --filename=weather_service.go --disable-version-string
package gateway

import (
	"context"

	"github.com/Berchon/weather-cloud-run/internal/business/model"
)

type WeatherService interface {
	GetWeatherByCity(ctx context.Context, city string) (*float64, *model.CustomError)
}
