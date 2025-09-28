package dependencies

import (
	"time"

	"github.com/Berchon/weather-cloud-run/internal/business/usecase"
	"github.com/Berchon/weather-cloud-run/internal/infrastructure/service"
	"github.com/Berchon/weather-cloud-run/internal/infrastructure/service/config"
	"github.com/Berchon/weather-cloud-run/internal/infrastructure/webapp/handler"
)

type Handlers struct {
	GetTemperatureByZipCodeHandler handler.GetTemperatureByZipCodeHandler
	GetStatusHandler               handler.GetStatusHandler
}

func BuildDependencies() *Handlers {
	// --- Clients ---
	httpClient := config.NewHTTPClient(3 * time.Second)

	// --- Repositories ---

	// --- Services ---
	viaCepService := service.NewViaCepService(httpClient)
	weatherService := service.NewWeatherService(httpClient)

	// --- UseCases ---
	getTemperatureByZipCodeUsecase := usecase.NewGetTemperatureByZipCodeUsecase(viaCepService, weatherService)

	// --- Handlers ---
	getTemperatureByZipCodeHandler := handler.NewGetTemperatureByZipCodeHandler(getTemperatureByZipCodeUsecase)
	getStatusHandler := handler.NewGetStatusHandler()

	return &Handlers{
		GetTemperatureByZipCodeHandler: getTemperatureByZipCodeHandler,
		GetStatusHandler:               getStatusHandler,
	}
}
