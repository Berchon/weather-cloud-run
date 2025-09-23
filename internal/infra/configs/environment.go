package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type cfg struct {
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	ViaCepAPIUrl  string `mapstructure:"VIACEP_API_URL"`
	WeatherAPIUrl string `mapstructure:"WEATHER_API_URL"`
	WeatherAPIKey string `mapstructure:"WEATHER_API_KEY"`
}

var (
	config   *cfg
	lastPath string
)

func LoadConfig(path string) error {
	lastPath = path

	viper.SetDefault("WEB_SERVER_PORT", "8080")
	viper.SetDefault("VIACEP_API_URL", "https://viacep.com.br/ws/")
	viper.SetDefault("WEATHER_API_URL", "https://api.weatherapi.com/v1/current.json")
	viper.SetDefault("WEATHER_API_KEY", "")

	viper.SetConfigFile(fmt.Sprintf("%s/.env", path))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Warning: is not possible to read .env file, using defaults values from env:", err)
	}

	var c cfg
	if err := viper.Unmarshal(&c); err != nil {
		return fmt.Errorf("Error to unmarshal config: %w", err)
	}

	config = &c
	return nil
}

func RefreshConfig() error {
	viper.Reset()
	return LoadConfig(lastPath)
}

func GetWebServerPort() string {
	return config.WebServerPort
}

func SetWebServerPort(port string) {
	config.WebServerPort = port
}

func GetViaCepAPIUrl() string {
	return config.ViaCepAPIUrl
}

func SetViaCepAPIUrl(url string) {
	config.ViaCepAPIUrl = url
}

func GetWeatherAPIUrl() string {
	return config.WeatherAPIUrl
}

func SetWeatherAPIUrl(url string) {
	config.WeatherAPIUrl = url
}

func GetWeatherAPIKey() string {
	return config.WeatherAPIKey
}

func SetWeatherAPIKey(key string) {
	config.WeatherAPIKey = key
}
