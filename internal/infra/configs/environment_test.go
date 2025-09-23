package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setEnvMock() {
	os.Setenv("WEB_SERVER_PORT", "1234")
	os.Setenv("VIACEP_API_URL", "http://mockviacep.com/ws/")
	os.Setenv("WEATHER_API_URL", "http://mockweather.com/api/")
	os.Setenv("WEATHER_API_KEY", "mockkey")
}

func unsetEnvMock() {
	os.Unsetenv("WEB_SERVER_PORT")
	os.Unsetenv("VIACEP_API_URL")
	os.Unsetenv("WEATHER_API_URL")
	os.Unsetenv("WEATHER_API_KEY")
}

func Test_LoadConfig(t *testing.T) {
	t.Run("When .env not found, should return default values", func(t *testing.T) {
		unsetEnvMock()
		err := LoadConfig(".")
		assert.NoError(t, err)

		assert.Equal(t, "8080", GetWebServerPort())
		assert.Equal(t, "https://viacep.com.br/ws/", GetViaCepAPIUrl())
		assert.Equal(t, "https://api.weatherapi.com/v1/current.json", GetWeatherAPIUrl())
		assert.Equal(t, "", GetWeatherAPIKey())
	})

	t.Run("When load .env successfully, should return file values", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()

		err := LoadConfig(".")
		assert.NoError(t, err)

		assert.Equal(t, "1234", GetWebServerPort())
		assert.Equal(t, "http://mockviacep.com/ws/", GetViaCepAPIUrl())
		assert.Equal(t, "http://mockweather.com/api/", GetWeatherAPIUrl())
		assert.Equal(t, "mockkey", GetWeatherAPIKey())
	})
}

func Test_RefreshConfig(t *testing.T) {
	t.Run("When reload configs, should return new values", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()

		_ = LoadConfig(".")

		os.Setenv("WEB_SERVER_PORT", "7777")
		os.Setenv("VIACEP_API_URL", "http://updated-viacep/")

		err := RefreshConfig()
		assert.NoError(t, err)

		assert.Equal(t, "7777", GetWebServerPort())
		assert.Equal(t, "http://updated-viacep/", GetViaCepAPIUrl())
	})

	t.Run("When reload configs and .env not found, should return default values", func(t *testing.T) {
		setEnvMock()
		_ = LoadConfig(".")
		unsetEnvMock()

		err := RefreshConfig()
		assert.NoError(t, err)

		assert.Equal(t, "8080", GetWebServerPort()) // default value
		assert.Equal(t, "https://viacep.com.br/ws/", GetViaCepAPIUrl())
		assert.Equal(t, "https://api.weatherapi.com/v1/current.json", GetWeatherAPIUrl())
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

	t.Run("ViaCepAPIUrl", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()
		_ = LoadConfig(".")

		assert.Equal(t, "http://mockviacep.com/ws/", GetViaCepAPIUrl())
		SetViaCepAPIUrl("http://changed-viacep/")
		assert.Equal(t, "http://changed-viacep/", GetViaCepAPIUrl())
	})

	t.Run("WeatherAPIUrl", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()
		_ = LoadConfig(".")

		assert.Equal(t, "http://mockweather.com/api/", GetWeatherAPIUrl())
		SetWeatherAPIUrl("http://changed-weather/")
		assert.Equal(t, "http://changed-weather/", GetWeatherAPIUrl())
	})

	t.Run("WeatherAPIKey", func(t *testing.T) {
		setEnvMock()
		defer unsetEnvMock()
		_ = LoadConfig(".")

		assert.Equal(t, "mockkey", GetWeatherAPIKey())
		SetWeatherAPIKey("NEWKEY")
		assert.Equal(t, "NEWKEY", GetWeatherAPIKey())
	})
}
