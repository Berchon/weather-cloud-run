package webapp

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Berchon/weather-cloud-run/internal/infrastructure/configs"
	"github.com/Berchon/weather-cloud-run/internal/infrastructure/dependencies"
	"github.com/Berchon/weather-cloud-run/internal/infrastructure/webapp/route"
)

type WebApp interface {
	Start()
}

type webApp struct{}

func New() WebApp {
	return &webApp{}
}

func (webApp *webApp) Start() {
	err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading configs:", err)
		panic(err)
	}

	dependencies := dependencies.BuildDependencies()

	router := route.ConfigureApplicationRoutes(dependencies)

	port := fmt.Sprintf(":%s", configs.GetWebServerPort())
	log.Printf("Starting server on port %s\n", port)
	http.ListenAndServe(port, router)
}
