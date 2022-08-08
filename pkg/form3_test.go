package form3

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {

	t.Run("Given no extra options should return a client with default values", func(t *testing.T) {
		// Act
		actual, err := NewClient()

		// Assert
		assert.Nil(t, err)
		assert.NotEmpty(t, actual)
		assert.Equal(t, defaultBaseURL, actual.baseUrl.String())
		assert.Empty(t, actual.userAgent)
		assert.Empty(t, actual.timeout)
		assert.NotEmpty(t, actual.Accounts)
	})

	t.Run("Given an option to set Http Client should return a client with that specific Http Client", func(t *testing.T) {
		// Arrange
		expected := &http.Client{Timeout: 1}

		// Act
		actual, err := NewClient(
			WithHttpClient(expected),
		)

		// Assert
		assert.Nil(t, err)
		assert.Same(t, expected, actual.httpClient)
	})

	t.Run("Given an option to set Timeout should return a client with that specific timeout", func(t *testing.T) {
		// Arrange
		expected := 10000

		// Act
		actual, err := NewClient(
			WithTimeoutInMilliseconds(expected),
		)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected, actual.timeout)
	})

	t.Run("Given an option to set a negative Timeout should return an error", func(t *testing.T) {
		// Arrange
		negativeTimeout := -10000

		// Act
		actual, err := NewClient(
			WithTimeoutInMilliseconds(negativeTimeout),
		)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, actual)
	})

	t.Run("Given an option to set a User Agent should return a client with that specific User Agent", func(t *testing.T) {
		// Arrange
		expected := "test/useragent"

		// Act
		actual, err := NewClient(
			WithUserAgent(expected),
		)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected , actual.userAgent)
	})

	t.Run("Given an option to set the Base Url should return a client with that specific Base Url", func(t *testing.T) {
		// Arrange
		expected := url.URL{
			Scheme: "http",
			Host:   "example.com",
		}

		// Act
		actual, err := NewClient(
			WithBaseUrl(expected),
		)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected.String() ,actual.baseUrl.String())
	})

}