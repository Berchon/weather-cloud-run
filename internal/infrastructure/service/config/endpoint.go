package config

import "fmt"

type Endpoint struct {
	url string
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (endpoint *Endpoint) GetUrl() string {
	return endpoint.url
}

func (endpoint *Endpoint) SetUrl(path string, paramKeys ...interface{}) {
	endpoint.url = buildURI(path, paramKeys...)
}

func buildURI(path string, paramKeys ...interface{}) string {
	return fmt.Sprintf(path, paramKeys...)
}
