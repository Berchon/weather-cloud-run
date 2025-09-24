package model

import (
	"fmt"
)

type CustomError struct {
	StatusCode int    `json:"status_code"`
	Err        string `json:"message"`
}

func NewCustomError(statusCode int, message string) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Err:        message,
	}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Erro %d: %v", e.StatusCode, e.Err)
}
