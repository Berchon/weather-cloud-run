package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndpoint_Build(t *testing.T) {
	t.Run("should create a new endpoint when NewEndpoint is called", func(t *testing.T) {
		e := NewEndpoint()
		assert.NotNil(t, e)
		assert.Equal(t, 0, len(e.query))
		assert.Nil(t, e.err)
	})

	t.Run("should return valid URL when SetBaseURL is called with valid URL", func(t *testing.T) {
		e := NewEndpoint().SetBaseURL(fmt.Sprintf("http://example.com/%s", "api"))
		url, err := e.Build()
		assert.NoError(t, err)
		assert.Equal(t, "http://example.com/api", url)
	})

	t.Run("should return error when SetBaseURL is called with invalid URL", func(t *testing.T) {
		e := NewEndpoint().SetBaseURL("://invalid-url") // invalid scheme
		_, err := e.Build()
		assert.Error(t, err)
	})

	t.Run("should return error when SetPath is called before baseURL is set", func(t *testing.T) {
		e := NewEndpoint().SetPath(fmt.Sprintf("/api/%d", 1))
		_, err := e.Build()
		assert.Error(t, err)
		assert.Equal(t, "baseURL must be set before path", err.Error())
	})

	t.Run("should return valid URL when SetPath is called after baseURL is set", func(t *testing.T) {
		e := NewEndpoint().SetBaseURL("http://example.com").SetPath(fmt.Sprintf("/api/%d", 123))
		url, err := e.Build()
		assert.NoError(t, err)
		assert.Equal(t, "http://example.com/api/123", url)
	})

	t.Run("should return error when AddQueryParam is called before baseURL is set", func(t *testing.T) {
		e := NewEndpoint().AddQueryParam("key", "value")
		_, err := e.Build()
		assert.Error(t, err)
		assert.Equal(t, "baseURL must be set before adding query parameters", err.Error())
	})

	t.Run("should return URL with query params when AddQueryParam is called after baseURL is set", func(t *testing.T) {
		e := NewEndpoint().SetBaseURL("http://example.com").
			AddQueryParam("key", "value").
			AddQueryParam("key2", "value2")

		url, err := e.Build()
		assert.NoError(t, err)
		assert.Contains(t, url, "key=value")
		assert.Contains(t, url, "key2=value2")
	})

	t.Run("should overwrite existing query param when AddQueryParam is called with same key", func(t *testing.T) {
		e := NewEndpoint().SetBaseURL("http://example.com").
			AddQueryParam("key", "value").
			AddQueryParam("key", "newvalue")

		url, err := e.Build()
		assert.NoError(t, err)
		assert.Contains(t, url, "key=newvalue")
		assert.NotContains(t, url, "key=value")
	})

	t.Run("should return complex URL when baseURL, path and multiple query params are set", func(t *testing.T) {
		e := NewEndpoint().
			SetBaseURL(fmt.Sprintf("https://example.com/%s", "v1")).
			SetPath(fmt.Sprintf("/api/%d/%s", 10, "test")).
			AddQueryParam("a", "1").
			AddQueryParam("b", "2")

		url, err := e.Build()
		assert.NoError(t, err)
		assert.Equal(t, "https://example.com/api/10/test?a=1&b=2", url)
	})

	t.Run("should return error when GetUrl is called and baseURL is not set", func(t *testing.T) {
		e := NewEndpoint()
		_, err := e.GetUrl()
		assert.Error(t, err)
		assert.Equal(t, "URL not set", err.Error())
	})

	t.Run("should accumulate errors when multiple methods fail", func(t *testing.T) {
		e := NewEndpoint()
		e.SetPath("/api/1")
		e.AddQueryParam("key", "value")
		_, err := e.Build()
		assert.Error(t, err)
		// first error should be reported
		assert.Equal(t, "baseURL must be set before path", err.Error())
	})

	// Additional robust tests
	t.Run("should return URL with empty path when path is not set", func(t *testing.T) {
		e := NewEndpoint().SetBaseURL("http://example.com")
		url, err := e.Build()
		assert.NoError(t, err)
		assert.Equal(t, "http://example.com", url)
	})

	t.Run("should handle empty query parameters correctly", func(t *testing.T) {
		e := NewEndpoint().SetBaseURL("http://example.com").SetPath("/test")
		url, err := e.Build()
		assert.NoError(t, err)
		assert.Equal(t, "http://example.com/test", url)
	})

	t.Run("should handle multiple SetBaseURL calls gracefully", func(t *testing.T) {
		e := NewEndpoint().
			SetBaseURL("http://example.com").
			SetBaseURL("http://example.org")
		url, err := e.Build()
		assert.NoError(t, err)
		assert.Equal(t, "http://example.org", url)
	})

	t.Run("should handle multiple SetPath calls gracefully", func(t *testing.T) {
		e := NewEndpoint().
			SetBaseURL("http://example.com").
			SetPath("/path1").
			SetPath("/path2")
		url, err := e.Build()
		assert.NoError(t, err)
		assert.Equal(t, "http://example.com/path2", url)
	})
}
