package core

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {

	t.Run("Given no extra build values should return a request with default values", func(t *testing.T) {
		// Arrange
		expected := &Request{
			Method: http.MethodGet,
			Path: "",
			QueryParam: url.Values{},
			Body: nil,
			Result: nil,
			Error: nil,
			Context: nil,
		}

		// Act
		actual := NewRequestBuilder(http.MethodGet).Build()

		// Assert
		assert.Equal(t, expected, actual)
	})

	t.Run("Given multiple paths should return a request with unified path", func(t *testing.T) {
		// Arrange
		expected := "some_path_1/some_path_2"

		// Act
		actual := NewRequestBuilder(http.MethodGet).
			WithPath("some_path_1").
			WithPath("some_path_2").
			Build()

		// Assert
		assert.Equal(t, expected, actual.Path)
	})

	t.Run("Given query parameters should return a request with respective url values", func(t *testing.T) {
		// Arrange
		expected := make(url.Values)
		expected.Add("some_parameter","some_value")

		// Act
		actual := NewRequestBuilder(http.MethodGet).
			WithQueryParam("some_parameter", "some_value").
			Build()

		// Assert
		assert.Equal(t, expected, actual.QueryParam)
	})

	t.Run("Given a body should return a request with respective body", func(t *testing.T) {
		// Arrange
		type SampleBody struct{}
		expected := new(SampleBody)

		// Act
		actual := NewRequestBuilder(http.MethodGet).
			WithBody(expected).
			Build()

		// Assert
		assert.Equal(t, expected, actual.Body)
	})

	t.Run("Given a context should return a request with respective context", func(t *testing.T) {
		// Arrange
		expected := context.Background()

		// Act
		actual := NewRequestBuilder(http.MethodGet).
			WithContext(expected).
			Build()

		// Assert
		assert.Equal(t, expected, actual.Context)
	})

	t.Run("Given a result writer should return a request with respective result writer", func(t *testing.T) {
		// Arrange
		type SampleWriteResult struct{}
		expected := new(SampleWriteResult)

		// Act
		actual := NewRequestBuilder(http.MethodGet).
			WithResultWriteTo(expected).
			Build()

		// Assert
		assert.Equal(t, expected, actual.Result)
	})

	t.Run("Given a error writer should return a request with respective error writer", func(t *testing.T) {
		// Arrange
		type SampleErrorWriter struct{}
		expected := new(SampleErrorWriter)

		// Act
		actual := NewRequestBuilder(http.MethodGet).
			WithErrorWriteTo(expected).
			Build()

		// Assert
		assert.Equal(t, expected, actual.Error)
	})

}