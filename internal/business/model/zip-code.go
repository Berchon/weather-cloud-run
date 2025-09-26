package model

import (
	"net/http"
	"regexp"
)

type ZipCode string

func BuildZipCode(stringZipCode string) (*ZipCode, *CustomError) {
	zipCode := ZipCode(stringZipCode)
	if zipCode.IsValidZipCode() {
		return &zipCode, nil
	}

	return nil, NewCustomError(http.StatusUnprocessableEntity, "invalid zipcode")
}

func (zipCode *ZipCode) IsValidZipCode() bool {
	regex := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return regex.MatchString(zipCode.ToString())
}

func (zipCode *ZipCode) ToString() string {
	return string(*zipCode)
}
