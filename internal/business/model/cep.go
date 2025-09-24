package model

import (
	"fmt"
	"net/http"
	"regexp"
)

type CEP string

func BuildCEP(stringCep string) (*CEP, *CustomError) {
	cep := CEP(stringCep)
	if cep.IsValidCEP() {
		return &cep, nil
	}

	return nil, NewCustomError(http.StatusBadRequest, fmt.Sprintf("CEP [%s] is not valid", stringCep))
}

func (cep *CEP) IsValidCEP() bool {
	regex := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return regex.MatchString(cep.ToString())
}

func (cep *CEP) ToString() string {
	return string(*cep)
}
