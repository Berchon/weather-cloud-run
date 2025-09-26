//go:generate mockery --dir=. --output=./mock --name=HTTPDoer --structname=MockHTTPDoer --outpkg=mock --filename=http_client.go --disable-version-string
package config

import (
	"net/http"
	"time"
)

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewHTTPClient(timeout time.Duration) *http.Client {
	if timeout <= 0 {
		timeout = 3 * time.Second
	}

	return &http.Client{
		Timeout: timeout,
	}
}
