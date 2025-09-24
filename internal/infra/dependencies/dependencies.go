package dependencies

import "github.com/Berchon/weather-cloud-run/internal/infra/webapp/handler"

type Handlers struct {
	GetTemperatureByCepHandlerHandler handler.GetTemperatureByCepHandler
	GetStatusHandler                  handler.GetStatusHandler
}

func BuildDependencies() *Handlers {
	// --- Repositories ---
	// userRepo := repository.NewUserRepository()

	// --- Services ---
	// userService := service.NewUserService(userRepo)

	// --- UseCases ---
	// getUserUC := usecase.NewGetUserUsecase(userService)

	// --- Handlers ---
	getTemperatureByCepHandler := handler.NewGetTemperatureByCepHandler()
	getStatusHandler := handler.NewGetStatusHandler()

	return &Handlers{
		GetTemperatureByCepHandlerHandler: getTemperatureByCepHandler,
		GetStatusHandler:                  getStatusHandler,
	}
}
