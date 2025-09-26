package model_test

import (
	"testing"

	"github.com/Berchon/weather-cloud-run/internal/business/model"
	"github.com/stretchr/testify/assert"
)

func TestBuildZipCode(t *testing.T) {
	t.Run("Should return a zip code When given a valid value", func(t *testing.T) {
		stringZipCode := "12345678"
		want := model.ZipCode(stringZipCode)
		got, err := model.BuildZipCode(stringZipCode)

		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Equal(t, want, *got)
	})

	t.Run("Should return a error When given a invalid value", func(t *testing.T) {
		stringZipCode := "abc"
		want := "invalid zipcode"
		got, err := model.BuildZipCode(stringZipCode)

		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.Equal(t, want, err.Error())
	})
}

func TestIsValidZipCode(t *testing.T) {
	t.Run("Should return true When zip code is valid", func(t *testing.T) {
		zipCode1 := model.ZipCode("12345678")
		zipCode2 := model.ZipCode("12345-678")

		assert.True(t, zipCode1.IsValidZipCode())
		assert.True(t, zipCode2.IsValidZipCode())
	})

	t.Run("Should return false When zip code is invalid with a character", func(t *testing.T) {
		zipCode1 := model.ZipCode("1234567a")
		zipCode2 := model.ZipCode("12345-67a")

		assert.False(t, zipCode1.IsValidZipCode())
		assert.False(t, zipCode2.IsValidZipCode())
	})

	t.Run("Should return false When zip code is invalid lenth", func(t *testing.T) {
		zipCode1 := model.ZipCode("12345-6789")
		zipCode2 := model.ZipCode("12345-6789")
		zipCode3 := model.ZipCode("12345-67")
		zipCode4 := model.ZipCode("12345-67")
		zipCode5 := model.ZipCode("12-345678")

		assert.False(t, zipCode1.IsValidZipCode())
		assert.False(t, zipCode2.IsValidZipCode())
		assert.False(t, zipCode3.IsValidZipCode())
		assert.False(t, zipCode4.IsValidZipCode())
		assert.False(t, zipCode5.IsValidZipCode())
	})
}
