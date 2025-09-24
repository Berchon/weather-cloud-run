package webapp

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Berchon/weather-cloud-run/internal/infra/configs"
	"github.com/Berchon/weather-cloud-run/internal/infra/dependencies"
	"github.com/Berchon/weather-cloud-run/internal/infra/webapp/route"
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
	http.ListenAndServe(port, router)
}
