//go:generate mockery --dir=. --output=./mock --name=GetTemperatureByCepUsecase --structname=MockGetTemperatureByCepUsecase --outpkg=mock --filename=get_temperature_by_cep_usecase.go --disable-version-string
package usecase

import (
	"context"
	"fmt"

	"github.com/Berchon/weather-cloud-run/internal/business/model"
)

type GetTemperatureByCepUsecase interface {
	GetTemperatureByCep(ctx context.Context, cep model.CEP) (*string, *model.CustomError)
}

type getTemperatureByCepUsecase struct {
	// brasilApiService gateway.BrasilApiService
}

func NewGetTemperatureByCepUsecase() GetTemperatureByCepUsecase { //brasilApiService gateway.BrasilApiService) GetTemperatureByCepUsecase {
	return &getTemperatureByCepUsecase{
		// BrasilApiService: brasilApiService,
	}
}

func (uc *getTemperatureByCepUsecase) GetTemperatureByCep(ctx context.Context, cep model.CEP) (*string, *model.CustomError) {
	fmt.Println("usecase")
	return nil, nil
}
