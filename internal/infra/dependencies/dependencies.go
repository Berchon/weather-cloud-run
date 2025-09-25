package dependencies

import (
	"github.com/Berchon/weather-cloud-run/internal/business/usecase"
	"github.com/Berchon/weather-cloud-run/internal/infra/webapp/handler"
)

type Handlers struct {
	GetTemperatureByCepHandler handler.GetTemperatureByCepHandler
	GetStatusHandler           handler.GetStatusHandler
}

func BuildDependencies() *Handlers {
	// --- Repositories ---
	// userRepo := repository.NewUserRepository()

	// --- Services ---
	// userService := service.NewUserService(userRepo)

	// --- UseCases ---
	getTemperatureByCepUsecase := usecase.NewGetTemperatureByCepUsecase()

	// --- Handlers ---
	getTemperatureByCepHandler := handler.NewGetTemperatureByCepHandler(getTemperatureByCepUsecase)
	getStatusHandler := handler.NewGetStatusHandler()

	return &Handlers{
		GetTemperatureByCepHandler: getTemperatureByCepHandler,
		GetStatusHandler:           getStatusHandler,
	}
}
