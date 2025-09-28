package service

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"

// 	"github.com/Berchon/weather-cloud-run/internal/business/gateway"
// 	"github.com/Berchon/weather-cloud-run/internal/business/model"
// 	"github.com/Berchon/weather-cloud-run/internal/infrastructure/configs"
// 	"github.com/Berchon/weather-cloud-run/internal/infrastructure/service/config"
// 	"github.com/Berchon/weather-cloud-run/internal/infrastructure/service/dto"
// )

// type viaCepService2 struct {
// 	endpoint *config.Endpoint
// }

// func NewViaCepService2() gateway.ViaCepService {
// 	endpoint := config.NewEndpoint()
// 	return &viaCepService{
// 		endpoint: endpoint,
// 	}
// }

// func (s *viaCepService2) GetAddressByZipCode(ctx context.Context, zipCode model.ZipCode) (*string, *model.CustomError) {
// 	url := fmt.Sprintf("%s%s", configs.GetViaCepBaseUrl(), configs.GetViaCepPath())
// 	s.endpoint.SetUrl(url, zipCode)
// 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.endpoint.GetUrl(), nil)
// 	if err != nil {
// 		return nil, model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error creating request: %v", err))
// 	}

// 	response, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return nil, model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error sending request: %v", err))
// 	}
// 	defer response.Body.Close()

// 	if response.StatusCode == http.StatusBadRequest {
// 		return nil, model.NewCustomError(http.StatusUnprocessableEntity, "invalid zipcode")
// 	}

// 	if response.StatusCode != http.StatusOK {
// 		return nil, model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error reading response. Status code not OK: Status code returned %v", response.StatusCode))
// 	}

// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error reading response: %v", err))
// 	}

// 	var viaCepDto dto.ViaCepDto
// 	if err := json.Unmarshal(body, &viaCepDto); err != nil {
// 		return nil, model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error unmarshalling response: %v", err))
// 	}

// 	if viaCepDto.Error == "true" {
// 		return nil, model.NewCustomError(http.StatusNotFound, "can not find zipcode")
// 	}

// 	return &viaCepDto.City, nil
// }
