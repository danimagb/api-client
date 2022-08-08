package core

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody(t *testing.T) {
	t.Run("Given RawResponse is nill should return default byte array", func(t *testing.T) {
		// Arrange
		response := &Response{}
		expected := []byte{}

		// Act
		actual := response.Body()

		// Assert
		assert.Equal(t, expected, actual)
	})

	t.Run("Given RawResponse is not nil should return RawResponse body", func(t *testing.T) {
		// Arrange
		body := "sample body"
		expected, err := json.Marshal(body)
		if err != nil{
			t.Errorf("Error marshaling body: %s", body)
		}

		httpResponse := &http.Response{}

		response := &Response{body: expected, RawResponse: httpResponse}

		// Act
		actual := response.Body()

		// Assert
		assert.Equal(t, expected, actual)
	})
}

func TestStatus(t *testing.T) {
	t.Run("Given RawResponse is nill should return default status", func(t *testing.T) {
		// Arrange
		response := &Response{}
		expected := ""

		// Act
		actual := response.Status()

		// Assert
		assert.Equal(t, expected, actual)
	})

	t.Run("Given RawResponse is not nil should return RawResponse status", func(t *testing.T) {
		// Arrange
		expected := "200 OK"

		httpResponse := &http.Response{Status: expected}

		response := &Response{RawResponse: httpResponse}

		// Act
		actual := response.Status()

		// Assert
		assert.Equal(t, expected, actual)
	})
}

func TestStatuCode(t *testing.T) {
	t.Run("Given RawResponse is nill should return default status code", func(t *testing.T) {
		// Arrange
		response := &Response{}
		expected := 0

		// Act
		actual := response.StatusCode()

		// Assert
		assert.Equal(t, expected, actual)
	})

	t.Run("Given RawResponse is not nil should return RawResponse status code", func(t *testing.T) {
		// Arrange
		expected := 200

		httpResponse := &http.Response{StatusCode: expected}

		response := &Response{RawResponse: httpResponse}

		// Act
		actual := response.StatusCode()

		// Assert
		assert.Equal(t, expected, actual)
	})
}

func TestIsSuccess(t *testing.T) {
	t.Run("Given StatusCode is >199 & <300 should return true", func(t *testing.T) {
		// Arrange
		httpResponse := &http.Response{StatusCode: 200}

		response := &Response{RawResponse: httpResponse}

		// Act
		actual := response.IsSuccess()

		// Assert
		assert.True(t, actual)
	})

	t.Run("Given StatusCode is not >199 || <300 should return false", func(t *testing.T) {
		// Arrange
		httpResponse := &http.Response{StatusCode: 400}

		response := &Response{RawResponse: httpResponse}

		// Act
		actual := response.IsSuccess()

		// Assert
		assert.False(t, actual)
	})
}

func TestIsError(t *testing.T) {
	t.Run("Given StatusCode is >399 should return true", func(t *testing.T) {
		// Arrange
		httpResponse := &http.Response{StatusCode: 400}

		response := &Response{RawResponse: httpResponse}

		// Act
		actual := response.IsError()

		// Assert
		assert.True(t, actual)
	})

	t.Run("Given StatusCode is not >399 should return false", func(t *testing.T) {
		// Arrange
		httpResponse := &http.Response{StatusCode: 400}

		response := &Response{RawResponse: httpResponse}

		// Act
		actual := response.IsSuccess()

		// Assert
		assert.False(t, actual)
	})
}

func TestUnmarshalJson(t *testing.T) {
	t.Run("Given body is empty should not throw an error", func(t *testing.T) {
		// Arrange
		type SampleType struct{
			EX 	string
		}
		sampleType := &SampleType{}

		sut := &Response{ body: []byte{}}

		// Act
		err := sut.UnmarshalJson(sampleType)

		// Assert
		assert.Nil(t, err)
	})

	t.Run("Given body is not empty should unmarshal", func(t *testing.T) {
		// Arrange
		type SampleType struct{
			EX 	string `json:"ex,omitempty"`
		}
		expected := &SampleType{
			EX: "test",
		}

		body, _ := json.Marshal(expected)

		sut := &Response{
			body: body,
		}

		// Act
		actual := &SampleType{}
		err := sut.UnmarshalJson(actual)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}