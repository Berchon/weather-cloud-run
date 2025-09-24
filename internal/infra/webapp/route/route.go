package route

import (
	"net/http"

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

	router.Get("/status", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	// router.Get("/reload", getAddressByCep.Handle)
	router.Get("/temperature/{cep}", handlers.GetTemperatureByCepHandlerHandler.Handle)

}
