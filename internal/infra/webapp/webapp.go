package webapp

import (
	"fmt"
	"log"

	"github.com/Berchon/weather-cloud-run/internal/infra/configs"
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
	fmt.Println(configs.GetWebServerPort())
	configs.RefreshConfig()

	fmt.Println(configs.GetWebServerPort())
	fmt.Println(configs.GetViaCepAPIUrl())
	fmt.Println(configs.GetWeatherAPIUrl())
	fmt.Println(configs.GetWeatherAPIKey())
}
