package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type cfg struct {
	WebServerPort  string `mapstructure:"WEB_SERVER_PORT"`
	ViaCepBaseUrl  string `mapstructure:"VIACEP_BASE_URL"`
	ViaCepPath     string `mapstructure:"VIACEP_PATH"`
	WeatherBaseUrl string `mapstructure:"WEATHER_BASE_URL"`
	WeatherPath    string `mapstructure:"WEATHER_PATH"`
	WeatherAPIKey  string `mapstructure:"WEATHER_API_KEY"`
}

var (
	config   *cfg
	lastPath string
)

func LoadConfig(path string) error {
	lastPath = path

	viper.SetDefault("WEB_SERVER_PORT", "8080")
	viper.SetDefault("VIACEP_BASE_URL", "https://viacep.com.br")
	viper.SetDefault("VIACEP_PATH", "/ws/%s/json")
	viper.SetDefault("WEATHER_BASE_URL", "https://api.weatherapi.com")
	viper.SetDefault("WEATHER_PATH", "/v1/current.json")
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

func GetViaCepBaseUrl() string {
	return config.ViaCepBaseUrl
}

func SetViaCepBaseUrl(url string) {
	config.ViaCepBaseUrl = url
}

func GetViaCepPath() string {
	return config.ViaCepPath
}

func SetViaCepPath(path string) {
	config.ViaCepPath = path
}

func GetWeatherBaseUrl() string {
	return config.WeatherBaseUrl
}

func SetWeatherBaseUrl(url string) {
	config.WeatherBaseUrl = url
}

func GetWeatherPath() string {
	return config.WeatherPath
}

func SetWeatherPath(path string) {
	config.WeatherPath = path
}

func GetWeatherAPIKey() string {
	return config.WeatherAPIKey
}

func SetWeatherAPIKey(key string) {
	config.WeatherAPIKey = key
}
