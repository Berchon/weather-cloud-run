package config

import (
	"testing"
	"time"
)

func TestNewHTTPClient(t *testing.T) {
	t.Run("fail - timeout menor ou igual a zero deve usar 3s", func(t *testing.T) {
		client := NewHTTPClient(0)
		if client.Timeout != 3*time.Second {
			t.Errorf("esperado timeout = 3s, obtido %v", client.Timeout)
		}
	})

	t.Run("success - timeout valido", func(t *testing.T) {
		client := NewHTTPClient(5 * time.Second)
		if client.Timeout != 5*time.Second {
			t.Errorf("esperado timeout = 5s, obtido %v", client.Timeout)
		}
	})
}
