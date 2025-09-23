package main

import "github.com/Berchon/weather-cloud-run/internal/infra/webapp"

func main() {
	app := webapp.New()
	app.Start()
}
