package core

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildHttpRequest(t *testing.T) {
	url := url.URL{
		Scheme: "http",
		Host:   "example.com",
	}

	t.Run("Given an error parsing the url should return an error", func(t *testing.T) {
		// Arrange
		request := NewRequestBuilder(http.MethodGet).
			WithPath("://"). // Forcing an invalid path
			Build()

		// Act
		actual, err := request.buildHttpRequest(url)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, actual)
	})

	t.Run("Given request with body should return http request with body", func(t *testing.T) {
		// Arrange
		body := "sample body"
		buf, err := json.Marshal(body)
		if err != nil{
			t.Errorf("Error marshaling body: %s", body)
		}

		expected := ioutil.NopCloser(bytes.NewReader(buf))

		request := NewRequestBuilder(http.MethodPost).
			WithBody(body).
			Build()

		// Act
		actual, err := request.buildHttpRequest(url)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected, actual.Body)
	})

	t.Run("Given request without body should return http request without body", func(t *testing.T) {
		// Arrange
		expected := http.NoBody

		request := NewRequestBuilder(http.MethodPost).Build()

		// Act
		actual, err := request.buildHttpRequest(url)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected, actual.Body)
	})

	t.Run("Given an error when marshaling body should return error", func(t *testing.T) {
		// Arrange
		// this represents something that is impossible to parse to json
		body := map[string]interface{}{
			"foo": make(chan int),
		}

		request := NewRequestBuilder(http.MethodPost).
			WithBody(body).
			Build()

		// Act
		actual, err := request.buildHttpRequest(url)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, actual)
	})

	t.Run("Given an error when creating the httpRequest should return an error", func(t *testing.T) {
		// Arrange
		request := NewRequestBuilder("INVALID METHOD").Build()

		// Act
		actual, err := request.buildHttpRequest(url)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, actual)
	})
}

func TestGetContext(t *testing.T) {
	t.Run("Given request with context should return that same context", func(t *testing.T) {
		// Arrange
		expected := context.Background()
		expected = context.WithValue(expected, "myKey", "myValue")

		request := NewRequestBuilder(http.MethodPost).
			WithContext(expected).
			Build()

		// Act
		actual := request.getContext()

		// Assert
		assert.Equal(t, expected, actual)
	})

	t.Run("Given request without context should return default context", func(t *testing.T) {
		// Arrange
		expected := context.Background()

		request := NewRequestBuilder(http.MethodPost).
			Build()

		// Act
		actual := request.getContext()

		// Assert
		assert.Equal(t, expected, actual)
	})
}
