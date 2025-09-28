package config

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func NewTestResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

type ErrorReader struct{}

func (e *ErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("erro simulado de leitura")
}
func (e *ErrorReader) Close() error { return nil }
