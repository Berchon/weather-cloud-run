package model_test

import (
	"testing"

	"github.com/Berchon/weather-cloud-run/internal/business/model"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomError(t *testing.T) {
	t.Run("Should create a CustomError with given status code and message", func(t *testing.T) {
		want := &model.CustomError{
			StatusCode: 400,
			Err:        "Bad Request",
		}
		got := model.NewCustomError(400, "Bad Request")

		assert.NotNil(t, got)
		assert.Equal(t, want.StatusCode, got.StatusCode)
		assert.Equal(t, want.Err, got.Err)
	})
}

func TestError(t *testing.T) {
	t.Run("Should return a string to a CustomError", func(t *testing.T) {
		want := "Bad Request"
		got := model.NewCustomError(400, "Bad Request").Error()

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}
