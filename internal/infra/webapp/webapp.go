package webapp

import (
	"log"

	"github.com/Berchon/weather-cloud-run/internal/infra/configs"
	"github.com/Berchon/weather-cloud-run/internal/infra/dependencies"
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

	_ = dependencies.BuildDependencies()

}
