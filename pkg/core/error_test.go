package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApiClientError(t *testing.T) {
	// Arrange
	expected := &ApiClientError{
		Reason: "some_reason",
		StatusCode: 1,
		Message: "some_message",
		Response: &Response{},
	}

	// Act
	actual := NewApiClientError("some_reason", 1, "some_message", &Response{})

	// Assert
	assert.Equal(t, expected, actual)
}
