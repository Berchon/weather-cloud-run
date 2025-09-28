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

type viaCepService struct {
	endpoint *config.Endpoint
	client   config.HTTPDoer
}

func NewViaCepService(client config.HTTPDoer) gateway.ViaCepService {
	if client == nil {
		client = config.NewHTTPClient(3 * time.Second)
	}
	return &viaCepService{
		endpoint: config.NewEndpoint(),
		client:   client,
	}
}

func (s *viaCepService) GetAddressByZipCode(ctx context.Context, zipCode model.ZipCode) (*string, *model.CustomError) {
	path := fmt.Sprintf(configs.GetViaCepPath(), zipCode)
	ep := s.endpoint.
		SetBaseURL(configs.GetViaCepBaseUrl()).
		SetPath(path)

	url, err := ep.Build()
	if err != nil {
		fmt.Println("Error building endpoint:", err)
		return nil, model.NewCustomError(http.StatusInternalServerError,
			fmt.Sprintf("Error building endpoint: %v", err))
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error creating request: %v", err))
	}

	response, err := s.client.Do(req)
	if err != nil {
		return nil, model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error sending request: %v", err))
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusBadRequest {
		return nil, model.NewCustomError(http.StatusUnprocessableEntity, "invalid zipcode")
	}

	if response.StatusCode != http.StatusOK {
		return nil, model.NewCustomError(http.StatusInternalServerError,
			fmt.Sprintf("error reading response. Status code not OK: Status code returned %v", response.StatusCode))
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error reading response: %v", err))
	}

	var viaCepDto dto.ViaCepDto
	if err := json.Unmarshal(body, &viaCepDto); err != nil {
		return nil, model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error unmarshalling response: %v", err))
	}

	if viaCepDto.Error == "true" {
		return nil, model.NewCustomError(http.StatusNotFound, "can not find zipcode")
	}

	if viaCepDto.City == "" {
		return nil, model.NewCustomError(http.StatusInternalServerError, "city is empty in response")
	}

	return &viaCepDto.City, nil
}
