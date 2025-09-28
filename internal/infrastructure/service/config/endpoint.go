package config

import (
	"errors"
	"fmt"
	"net/url"
)

type Endpoint struct {
	baseURL string
	path    string
	query   url.Values
	err     error
}

func NewEndpoint() *Endpoint {
	return &Endpoint{query: url.Values{}}
}

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

func (e *Endpoint) AddQueryParam(key, value string) *Endpoint {
	if e.err != nil {
		return e
	}

	if e.baseURL == "" {
		e.err = errors.New("baseURL must be set before adding query parameters")
		return e
	}

	if key == "" || value == "" {
		e.err = errors.New("key and value must be non-empty for query parameters")
		return e
	}

	e.query.Set(key, value)
	return e
}

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

func (e *Endpoint) Build() (string, error) {
	return e.GetUrl()
}
