package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setEnvMock() {
	os.Setenv("WEB_SERVER_PORT", "1234")
	os.Setenv("VIACEP_BASE_URL", "http://mockviacep.com")
	os.Setenv("VIACEP_PATH", "/mock/ws/%s/json")
	os.Setenv("WEATHER_BASE_URL", "http://mockweather.com")
	os.Setenv("WEATHER_PATH", "/mock/v1/current.json")
	os.Setenv("WEATHER_API_KEY", "mock-key")
}

func unsetEnvMock() {
	os.Unsetenv("WEB_SERVER_PORT")
	os.Unsetenv("VIACEP_BASE_URL")
	os.Unsetenv("VIACEP_PATH")
	os.Unsetenv("WEATHER_BASE_URL")
	os.Unsetenv("WEATHER_PATH")
	os.Unsetenv("WEATHER_API_KEY")
}

func Test_LoadConfig(t *testing.T) {
	t.Run("When .env not found, should return default values", func(t *testing.T) {
		unsetEnvMock()
		err := LoadConfig(".")
		assert.NoError(t, err)

		assert.Equal(t, "8080", GetWebServerPort())
		assert.Equal(t, "https://viacep.com.br", GetViaCepBaseUrl())
		assert.Equal(t, "/ws/%s/json", GetViaCepPath())
		assert.Equal(t, "https://api.weatherapi.com", GetWeatherBaseUrl())
		assert.Equal(t, "/v1/current.json", GetWeatherPath())
		assert.Equal(t, "", GetWeatherAPIKey())
	})

	t.Run("When load .env successfully, should return file values", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()

		err := LoadConfig(".")
		assert.NoError(t, err)

		assert.Equal(t, "1234", GetWebServerPort())
		assert.Equal(t, "http://mockviacep.com", GetViaCepBaseUrl())
		assert.Equal(t, "/mock/ws/%s/json", GetViaCepPath())
		assert.Equal(t, "http://mockweather.com", GetWeatherBaseUrl())
		assert.Equal(t, "/mock/v1/current.json", GetWeatherPath())
		assert.Equal(t, "mock-key", GetWeatherAPIKey())
	})
}

func Test_RefreshConfig(t *testing.T) {
	t.Run("When reload configs, should return new values", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()

		_ = LoadConfig(".")

		os.Setenv("WEB_SERVER_PORT", "7777")
		os.Setenv("VIACEP_BASE_URL", "http://updated-viacep/")

		err := RefreshConfig()
		assert.NoError(t, err)

		assert.Equal(t, "7777", GetWebServerPort())
		assert.Equal(t, "http://updated-viacep/", GetViaCepBaseUrl())
	})

	t.Run("When reload configs and .env not found, should return default values", func(t *testing.T) {
		setEnvMock()
		_ = LoadConfig(".")
		unsetEnvMock()

		err := RefreshConfig()
		assert.NoError(t, err)

		assert.Equal(t, "8080", GetWebServerPort()) // default value
		assert.Equal(t, "https://viacep.com.br", GetViaCepBaseUrl())
		assert.Equal(t, "/ws/%s/json", GetViaCepPath())
		assert.Equal(t, "https://api.weatherapi.com", GetWeatherBaseUrl())
		assert.Equal(t, "/v1/current.json", GetWeatherPath())
		assert.Equal(t, "", GetWeatherAPIKey())
	})
}

func Test_SetAndGet(t *testing.T) {
	t.Run("WebServerPort", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()
		_ = LoadConfig(".")

		assert.Equal(t, "1234", GetWebServerPort())
		SetWebServerPort("9999")
		assert.Equal(t, "9999", GetWebServerPort())
	})

	t.Run("ViaCepBaseUrl", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()
		_ = LoadConfig(".")

		assert.Equal(t, "http://mockviacep.com", GetViaCepBaseUrl())
		SetViaCepBaseUrl("http://changed-viacep")
		assert.Equal(t, "http://changed-viacep", GetViaCepBaseUrl())
	})

	t.Run("ViaCepPath", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()
		_ = LoadConfig(".")

		assert.Equal(t, "/mock/ws/%s/json", GetViaCepPath())
		SetViaCepPath("/changed-mock/ws/%s/json")
		assert.Equal(t, "/changed-mock/ws/%s/json", GetViaCepPath())
	})

	t.Run("WeatherBaseUrl", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()
		_ = LoadConfig(".")

		assert.Equal(t, "http://mockweather.com", GetWeatherBaseUrl())
		SetWeatherBaseUrl("http://changed-weather")
		assert.Equal(t, "http://changed-weather", GetWeatherBaseUrl())
	})

	t.Run("WeatherPath", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()
		_ = LoadConfig(".")

		assert.Equal(t, "/mock/v1/current.json", GetWeatherPath())
		SetWeatherPath("/changed-mock/v1/current.json")
		assert.Equal(t, "/changed-mock/v1/current.json", GetWeatherPath())
	})

	t.Run("WeatherAPIKey", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()
		_ = LoadConfig(".")

		assert.Equal(t, "mock-key", GetWeatherAPIKey())
		SetWeatherAPIKey("NEW-KEY")
		assert.Equal(t, "NEW-KEY", GetWeatherAPIKey())
	})
}
