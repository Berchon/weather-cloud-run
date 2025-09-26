package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpoint(t *testing.T) {
	t.Run("success - path sem params", func(t *testing.T) {
		ep := NewEndpoint()
		ep.SetUrl("http://host/test")
		assert.Equal(t, "http://host/test", ep.GetUrl())
	})

	t.Run("success - path com params", func(t *testing.T) {
		ep := NewEndpoint()
		ep.SetUrl("http://host/ws/%s/json", "12345-678")
		assert.Equal(t, "http://host/ws/12345-678/json", ep.GetUrl())
	})

	t.Run("success - buildURI direto", func(t *testing.T) {
		got := buildURI("http://host/%s", "abc")
		assert.Equal(t, "http://host/abc", got)
	})
}
