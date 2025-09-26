//go:generate mockery --dir=. --output=./mock --name=GetTemperatureByZipCodeUsecase --structname=MockGetTemperatureByZipCodeUsecase --outpkg=mock --filename=get_temperature_by_zip_code_usecase.go --disable-version-string
package usecase

import (
	"context"
	"fmt"

	"github.com/Berchon/weather-cloud-run/internal/business/gateway"
	"github.com/Berchon/weather-cloud-run/internal/business/model"
)

type GetTemperatureByZipCodeUsecase interface {
	GetTemperatureByZipCode(ctx context.Context, zipCode model.ZipCode) (*string, *model.CustomError)
}

type getTemperatureByZipCodeUsecase struct {
	ViaCepService gateway.ViaCepService
}

func NewGetTemperatureByZipCodeUsecase(viaCepService gateway.ViaCepService) GetTemperatureByZipCodeUsecase {
	return &getTemperatureByZipCodeUsecase{
		ViaCepService: viaCepService,
	}
}

func (uc *getTemperatureByZipCodeUsecase) GetTemperatureByZipCode(ctx context.Context, zipCode model.ZipCode) (*string, *model.CustomError) {
	fmt.Println("usecase")
	city, err := uc.ViaCepService.GetAddressByZipCode(ctx, zipCode)
	if err != nil {
		return nil, err
	}

	fmt.Println("city", *city)
	return city, nil
}
