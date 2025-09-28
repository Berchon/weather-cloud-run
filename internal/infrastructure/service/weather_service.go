package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Berchon/weather-cloud-run/internal/business/gateway"
	"github.com/Berchon/weather-cloud-run/internal/business/model"
	"github.com/Berchon/weather-cloud-run/internal/infrastructure/configs"
	"github.com/Berchon/weather-cloud-run/internal/infrastructure/service/config"
	"github.com/Berchon/weather-cloud-run/internal/infrastructure/service/dto"
)

type weatherService struct {
	endpoint *config.Endpoint
	client   config.HTTPDoer
}

func NewWeatherService(client config.HTTPDoer) gateway.WeatherService {
	if client == nil {
		client = config.NewHTTPClient(3 * time.Second)
	}

	return &weatherService{
		endpoint: config.NewEndpoint(),
		client:   client,
	}
}

func (s *weatherService) GetWeatherByCity(ctx context.Context, city string) (*map[string]float64, *model.CustomError) {
	ep := s.endpoint.
		SetBaseURL(configs.GetWeatherBaseUrl()).
		SetPath(configs.GetWeatherPath()).
		AddQueryParam("key", configs.GetWeatherAPIKey()).
		AddQueryParam("q", city).
		AddQueryParam("aqi", "no")

	url, err := ep.Build()
	if err != nil {
		fmt.Println("Error building endpoint:", err)
		return nil, model.NewCustomError(http.StatusInternalServerError,
			fmt.Sprintf("Error building endpoint: %v", err))
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, model.NewCustomError(http.StatusInternalServerError,
			fmt.Sprintf("error creating request: %v", err))
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, model.NewCustomError(http.StatusInternalServerError,
			fmt.Sprintf("error sending request: %v", err))
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, model.NewCustomError(http.StatusInternalServerError,
			fmt.Sprintf("error reading response: %v", err))
	}

	// tratamento de erro da WeatherAPI
	if resp.StatusCode != http.StatusOK {
		var apiErr dto.WeatherErrorDto
		if err := json.Unmarshal(body, &apiErr); err == nil && apiErr.Error.Message != "" {
			return nil, model.NewCustomError(resp.StatusCode,
				fmt.Sprintf("Error code: %d. Description: %s", apiErr.Error.Code, apiErr.Error.Message))
		}
		return nil, model.NewCustomError(resp.StatusCode,
			fmt.Sprintf("unexpected error from weather api: %s", string(body)))
	}

	// sucesso → parse da resposta
	var weatherDto dto.WeatherDto
	if err := json.Unmarshal(body, &weatherDto); err != nil {
		return nil, model.NewCustomError(http.StatusInternalServerError,
			fmt.Sprintf("error unmarshalling response: %v", err))
	}

	// conversões
	tempC := weatherDto.Current.TempC
	tempF := tempC*1.8 + 32
	tempK := tempC + 273

	result := map[string]float64{
		"temp_C": tempC,
		"temp_F": tempF,
		"temp_K": tempK,
	}

	return &result, nil
}
