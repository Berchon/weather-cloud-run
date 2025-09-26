//go:generate mockery --dir=. --output=../../infrastructure/service/mock --name=ViaCepService --structname=MockViaCepService --outpkg=mock --filename=via_cep_service.go --disable-version-string
package gateway

import (
	"context"

	"github.com/Berchon/weather-cloud-run/internal/business/model"
)

type ViaCepService interface {
	GetAddressByZipCode(ctx context.Context, zipCode model.ZipCode) (*string, *model.CustomError)
}
