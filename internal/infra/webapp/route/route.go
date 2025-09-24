package route

import (
	"github.com/Berchon/weather-cloud-run/internal/infra/configs"
	"github.com/Berchon/weather-cloud-run/internal/infra/dependencies"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func ConfigureApplicationRoutes(handlers *dependencies.Handlers) *chi.Mux {
	port := configs.GetWebServerPort()
	router := chi.NewRouter()
	registerRoutes(port, router, handlers)
	return router
}

func registerRoutes(port string, router *chi.Mux, handlers *dependencies.Handlers) {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/status", handlers.GetStatusHandler.Handle)
	// router.Post("/reload", getAddressByCep.Handle)
	router.Get("/temperature/{cep}", handlers.GetTemperatureByCepHandlerHandler.Handle)

}
