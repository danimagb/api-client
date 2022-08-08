package accounts

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/danimagb/api-client/pkg/core"
	"github.com/danimagb/api-client/pkg/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedBaseClient struct {
	mock.Mock
}

func (m *MockedBaseClient) Send(req *core.Request) (*core.Response, error) {

	args := m.Called(req)

	if args.Get(0) != nil {
		return args.Get(0).(*core.Response), args.Error(1)
	}

	return nil, args.Error(1)
}

func TestFetch(t *testing.T) {
	t.Run("Given an error calling base client should return an error", func(t *testing.T) {
		// Arrange
		expectedError := fmt.Errorf("Some error occurred")

		mockedBaseClient := new(MockedBaseClient)
		mockedBaseClient.On("Send", mock.Anything).Return(nil, expectedError)

		sut := New(mockedBaseClient)
		// Act

		actual, err := sut.Fetch(context.Background(), uuid.New())

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, err, expectedError)
		assert.Nil(t, actual)
	})

	t.Run("Given a response with status code 200 should not return error", func(t *testing.T) {
		// Arrange
		statusCode := 200
		httpResponse := &http.Response{StatusCode: statusCode, Body: ioutil.NopCloser(bytes.NewBuffer(nil))}

		mockedResponse := &core.Response{
			RawResponse: httpResponse,
		}

		mockedBaseClient := new(MockedBaseClient)
		mockedBaseClient.On("Send", mock.Anything).Return(mockedResponse, nil)

		sut := New(mockedBaseClient)
		// Act

		actual, err := sut.Fetch(context.Background(), uuid.New())

		// Assert
		assert.Nil(t, err)
		assert.NotNil(t, actual)
	})

	t.Run("Given a response with status code other than 200 should return error", func(t *testing.T) {
		// Arrange
		statusCode := 500
		httpResponse := &http.Response{StatusCode: statusCode, Body: ioutil.NopCloser(bytes.NewBuffer(nil))}

		mockedResponse := &core.Response{
			RawResponse: httpResponse,
		}

		mockedBaseClient := new(MockedBaseClient)
		mockedBaseClient.On("Send", mock.Anything).Return(mockedResponse, nil)

		sut := New(mockedBaseClient)
		// Act

		actual, err := sut.Fetch(context.Background(), uuid.New())

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, actual)
		assert.IsType(t, err.(*core.ApiClientError), err)
		assert.IsType(t, http.StatusInternalServerError, err.(*core.ApiClientError).StatusCode)
	})

}


func TestCreate(t *testing.T) {
	t.Run("Given an error calling base client should return an error", func(t *testing.T) {
		// Arrange
		expectedError := fmt.Errorf("Some error occurred")

		mockedBaseClient := new(MockedBaseClient)
		mockedBaseClient.On("Send", mock.Anything).Return(nil, expectedError)

		sut := New(mockedBaseClient)
		// Act

		actual, err := sut.Create(context.Background(), &models.AccountRequest{})

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, err, expectedError)
		assert.Nil(t, actual)
	})

	t.Run("Given a response with status code 201 should not return error", func(t *testing.T) {
		// Arrange
		statusCode := 201
		httpResponse := &http.Response{StatusCode: statusCode, Body: ioutil.NopCloser(bytes.NewBuffer(nil))}

		mockedResponse := &core.Response{
			RawResponse: httpResponse,
		}

		mockedBaseClient := new(MockedBaseClient)
		mockedBaseClient.On("Send", mock.Anything).Return(mockedResponse, nil)

		sut := New(mockedBaseClient)

		// Act
		actual, err := sut.Create(context.Background(), &models.AccountRequest{})

		// Assert
		assert.Nil(t, err)
		assert.NotNil(t, actual)
	})

	t.Run("Given a response with status code other than 201 should return error", func(t *testing.T) {
		// Arrange
		statusCode := 500
		httpResponse := &http.Response{StatusCode: statusCode, Body: ioutil.NopCloser(bytes.NewBuffer(nil))}

		expected := &core.Response{
			RawResponse: httpResponse,
		}

		mockedBaseClient := new(MockedBaseClient)
		mockedBaseClient.On("Send", mock.Anything).Return(expected, nil)

		sut := New(mockedBaseClient)
		// Act

		actual, err := sut.Create(context.Background(), &models.AccountRequest{})

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, actual)
		assert.IsType(t, err.(*core.ApiClientError), err)
		assert.IsType(t, http.StatusInternalServerError, err.(*core.ApiClientError).StatusCode)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Given an error calling base client should return an error", func(t *testing.T) {
		// Arrange
		expectedError := fmt.Errorf("Some error occurred")

		mockedBaseClient := new(MockedBaseClient)
		mockedBaseClient.On("Send", mock.Anything).Return(nil, expectedError)

		sut := New(mockedBaseClient)
		// Act

		err := sut.Delete(context.Background(), uuid.New(), 0)

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, err, expectedError)
	})

	t.Run("Given a response with status code 204 should not return error", func(t *testing.T) {
		// Arrange
		statusCode := 204
		httpResponse := &http.Response{StatusCode: statusCode, Body: ioutil.NopCloser(bytes.NewBuffer(nil))}

		mockedResponse := &core.Response{
			RawResponse: httpResponse,
		}

		mockedBaseClient := new(MockedBaseClient)
		mockedBaseClient.On("Send", mock.Anything).Return(mockedResponse, nil)

		sut := New(mockedBaseClient)

		// Act
		err := sut.Delete(context.Background(), uuid.New(), 0)

		// Assert
		assert.Nil(t, err)
	})

	t.Run("Given a response with status code other than 204 should return error", func(t *testing.T) {
		// Arrange
		statusCode := 500
		httpResponse := &http.Response{StatusCode: statusCode, Body: ioutil.NopCloser(bytes.NewBuffer(nil))}

		expected := &core.Response{
			RawResponse: httpResponse,
		}

		mockedBaseClient := new(MockedBaseClient)
		mockedBaseClient.On("Send", mock.Anything).Return(expected, nil)

		sut := New(mockedBaseClient)
		// Act

		err := sut.Delete(context.Background(), uuid.New(), 0)

		// Assert
		assert.NotNil(t, err)
		assert.IsType(t, err.(*core.ApiClientError), err)
		assert.IsType(t, http.StatusInternalServerError, err.(*core.ApiClientError).StatusCode)
	})
}