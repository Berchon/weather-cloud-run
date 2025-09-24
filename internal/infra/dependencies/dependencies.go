package dependencies

import "github.com/Berchon/weather-cloud-run/internal/infra/webapp/handler"

type Handlers struct {
	GetTemperatureByCepHandlerHandler handler.GetTemperatureByCepHandler
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

	return &Handlers{
		GetTemperatureByCepHandlerHandler: getTemperatureByCepHandler,
	}
}
