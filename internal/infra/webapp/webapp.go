package webapp

// import (
// 	"fmt"

// 	"github.com/Berchon/weather-cloud-run/internal/infra/configs"
// )

type WebApp interface {
	Start()
}

type webApp struct{}

func New() WebApp {
	return &webApp{}
}

func (webApp *webApp) Start() {

}
