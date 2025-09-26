package main

import "github.com/Berchon/weather-cloud-run/internal/infrastructure/webapp"

func main() {
	app := webapp.New()
	app.Start()
}
