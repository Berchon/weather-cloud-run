package gateway

import (
	"context"

	"github.com/Berchon/weather-cloud-run/internal/business/model"
)

type WeatherService interface {
	GetWeatherByCity(ctx context.Context, city string) (*map[string]float64, *model.CustomError)
}
