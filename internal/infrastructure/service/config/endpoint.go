package config

import (
	"errors"
	"fmt"
	"net/url"
)

// Endpoint é um builder seguro para criar URLs com base, path e query params
type Endpoint struct {
	baseURL string
	path    string
	query   url.Values
	err     error
}

// NewEndpoint cria um endpoint vazio
func NewEndpoint() *Endpoint {
	return &Endpoint{query: url.Values{}}
}

// SetBaseURL define a base URL já formatada (ex: "http://host")
func (e *Endpoint) SetBaseURL(baseURL string) *Endpoint {
	if e.err != nil {
		return e
	}

	parsed, err := url.Parse(baseURL)
	if err != nil {
		e.err = fmt.Errorf("invalid baseURL: %w", err)
		return e
	}

	e.baseURL = parsed.Scheme + "://" + parsed.Host
	e.path = parsed.Path
	e.query = parsed.Query()

	return e
}

// SetPath define o path já formatado (ex: "/api/123/abc")
func (e *Endpoint) SetPath(path string) *Endpoint {
	if e.err != nil {
		return e
	}

	if e.baseURL == "" {
		e.err = errors.New("baseURL must be set before path")
		return e
	}

	e.path = path
	return e
}

// AddQueryParam adiciona query parameter, sobrescrevendo se já existir
func (e *Endpoint) AddQueryParam(key, value string) *Endpoint {
	if e.err != nil {
		return e
	}

	if e.baseURL == "" {
		e.err = errors.New("baseURL must be set before adding query parameters")
		return e
	}

	e.query.Set(key, value)
	return e
}

// GetUrl retorna a URL atual e erro acumulado
func (e *Endpoint) GetUrl() (string, error) {
	if e.err != nil {
		return "", e.err
	}
	if e.baseURL == "" {
		return "", errors.New("URL not set")
	}

	u := &url.URL{
		Scheme:   "http",
		Host:     "",
		Path:     e.path,
		RawQuery: e.query.Encode(),
	}

	parsed, err := url.Parse(e.baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid baseURL: %w", err)
	}
	u.Scheme = parsed.Scheme
	u.Host = parsed.Host

	return u.String(), nil
}

// Build retorna a URL final e erro acumulado (idêntico ao GetUrl, serve como finalizador)
func (e *Endpoint) Build() (string, error) {
	return e.GetUrl()
}
