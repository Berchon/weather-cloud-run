package model_test

import (
	"testing"

	"github.com/Berchon/weather-cloud-run/internal/business/model"
	"github.com/stretchr/testify/assert"
)

func TestBuildCEP(t *testing.T) {
	t.Run("Should return a zip code When given a valid value", func(t *testing.T) {
		stringCep := "12345678"
		want := model.CEP(stringCep)
		got, err := model.BuildCEP(stringCep)

		assert.NotNil(t, got)
		assert.Nil(t, err)
		assert.Equal(t, want, *got)
	})

	t.Run("Should return a error When given a invalid value", func(t *testing.T) {
		stringCep := "abc"
		want := "Erro 400: CEP [abc] is not valid"
		got, err := model.BuildCEP(stringCep)

		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.Equal(t, want, err.Error())
	})
}

func TestIsValidCEP(t *testing.T) {
	t.Run("Should return true When zip code is valid", func(t *testing.T) {
		cep1 := model.CEP("12345678")
		cep2 := model.CEP("12345-678")

		assert.True(t, cep1.IsValidCEP())
		assert.True(t, cep2.IsValidCEP())
	})

	t.Run("Should return false When zip code is invalid with a character", func(t *testing.T) {
		cep1 := model.CEP("1234567a")
		cep2 := model.CEP("12345-67a")

		assert.False(t, cep1.IsValidCEP())
		assert.False(t, cep2.IsValidCEP())
	})

	t.Run("Should return false When zip code is invalid lenth", func(t *testing.T) {
		cep1 := model.CEP("12345-6789")
		cep2 := model.CEP("12345-6789")
		cep3 := model.CEP("12345-67")
		cep4 := model.CEP("12345-67")
		cep5 := model.CEP("12-345678")

		assert.False(t, cep1.IsValidCEP())
		assert.False(t, cep2.IsValidCEP())
		assert.False(t, cep3.IsValidCEP())
		assert.False(t, cep4.IsValidCEP())
		assert.False(t, cep5.IsValidCEP())
	})
}
