package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)


type MockedHttpClient struct {
	mock.Mock
}

func(m *MockedHttpClient) Do(req *http.Request) (*http.Response, error) {

	args := m.Called(req)

	if(args.Get(0) != nil){
		return args.Get(0).(*http.Response), args.Error(1)
	}

	return nil, args.Error(1)
}

func TestSend(t *testing.T) {
	url := url.URL{
		Scheme: "http",
		Host:   "example.com",
	}

	testCases := []struct{
		responseStatusCode int
		requestMethod string
	}{
		{100, "GET"},{200, "GET"},{300, "GET"},{400, "GET"},{500, "GET"},
		{100, "POST"},{200, "POST"},{300, "POST"},{400, "POST"},{500, "POST"},
		{100, "DELETE"},{200, "DELETE"},{300, "DELETE"},{400, "DELETE"},{500, "DELETE"},
		{100, "PATCH"},{200, "PATCH"},{300, "PATCH"},{400, "PATCH"},{500, "PATCH"},
	}

	for _, tc := range testCases{
		t.Run("Given httpClient returns HttpResponse regardless of status code should return Response", func(t *testing.T) {
			// Arrange
			apiReq := NewRequestBuilder(tc.requestMethod).
				Build()

			httpResponse := &http.Response{StatusCode: tc.responseStatusCode, Body: ioutil.NopCloser(bytes.NewBuffer(nil))}

			mockedHttpClient := new(MockedHttpClient)
			mockedHttpClient.On("Do", mock.Anything).Return(httpResponse, nil)

			expected := &Response{
				RawResponse: httpResponse,
				body: []byte{},
			}

			sut := &BaseClient{
				BaseUrl: url,
				HttpClient: mockedHttpClient,
				Timeout: 1,
			}

			// Act
			actual, err := sut.Send(apiReq)

			// Assert
			assert.Nil(t, err)
			assert.Equal(t, expected, actual)
		})
	}

	t.Run("Given httpClient returns an error should return an error", func(t *testing.T) {
		// Arrange
		apiReq := NewRequestBuilder(http.MethodGet).
			Build()

		httpError := fmt.Errorf("Some error occurred")

		mockedHttpClient := new(MockedHttpClient)
		mockedHttpClient.On("Do", mock.Anything).Return(nil, httpError)

		sut := &BaseClient{
			BaseUrl: url,
			HttpClient: mockedHttpClient,
			Timeout: 1,
		}

		// Act
		actual, err := sut.Send(apiReq)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, actual)
	})

	t.Run("Given an error building http request should return an error without calling http client Do()", func(t *testing.T) {
		// Arrange
		apiReq := NewRequestBuilder("INVALID METHOD").
			Build()

		httpResponse := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer(nil))}

		mockedHttpClient := new(MockedHttpClient)
		mockedHttpClient.On("Do", mock.Anything).Return(httpResponse, nil)

		sut := &BaseClient{
			BaseUrl: url,
			HttpClient: mockedHttpClient,
			Timeout: 1,
		}

		// Act
		actual, err := sut.Send(apiReq)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, actual)
		mockedHttpClient.AssertNotCalled(t,"Do", mock.Anything)
	})

	t.Run("Given an error parsing http response should return an error", func(t *testing.T) {
		// Arrange
		type SampleType struct{}

		apiReq := NewRequestBuilder("GET").
			WithResultWriteTo(new(SampleType)).
			Build()

		//Wrongly formatted json will cause an error parsing the body
		httpResponse := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString("{'},"))}

		mockedHttpClient := new(MockedHttpClient)
		mockedHttpClient.On("Do", mock.Anything).Return(httpResponse, nil)

		sut := &BaseClient{
			BaseUrl: url,
			HttpClient: mockedHttpClient,
			Timeout: 1,
		}

		// Act
		actual, err := sut.Send(apiReq)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, actual)
		mockedHttpClient.AssertCalled(t,"Do", mock.Anything)
	})
}


func TestHandleHttpResponse(t *testing.T) {
	t.Run("Given a valid Http Response should return Response", func(t *testing.T) {
		// Arrange
		apiReq := NewRequestBuilder("GET").
			Build()

		httpResponse := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBuffer(nil))}

		expected := &Response{
			RawResponse: httpResponse,
			body: []byte{},
		}

		sut := &BaseClient{}

		// Act
		actual, err := sut.handleHttpResponse(apiReq, httpResponse)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Given an invalid body in Http Response should return an error", func(t *testing.T) {
		// Arrange
		type SampleType struct{
		}

		apiReq := NewRequestBuilder("GET").
			WithResultWriteTo(new(SampleType)).
			Build()

		httpResponse := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewBufferString("{'},"))}


		sut := &BaseClient{}

		// Act
		actual, err := sut.handleHttpResponse(apiReq, httpResponse)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, actual)
	})
}


func TestParseResponseBody(t *testing.T) {
	t.Run("Given a success Response should parse response body into resultValue", func(t *testing.T) {
		// Arrange
		type SampleType struct{
			EX 	string `json:"ex,omitempty"`
		}
		expected := &SampleType{
			EX: "test",
		}

		body, _ := json.Marshal(expected)

		httpResponse := &http.Response{StatusCode: 200, Body: ioutil.NopCloser(bytes.NewReader(body))}

		defer httpResponse.Body.Close()

		responseBody, _ := ioutil.ReadAll(httpResponse.Body)

		apiResponse := &Response{
			RawResponse: httpResponse,
			body: responseBody,
		}

		sut := &BaseClient{}

		// Act
		actual := &SampleType{}
		err := sut.parseResponseBody(apiResponse, actual, nil)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})

	t.Run("Given an unsuccessful Response should parse response body into errorValue", func(t *testing.T) {
		// Arrange
		type SampleType struct{
			EX 	string `json:"ex,omitempty"`
		}
		expected := &SampleType{
			EX: "test",
		}

		body, _ := json.Marshal(expected)

		httpResponse := &http.Response{StatusCode: 500, Body: ioutil.NopCloser(bytes.NewReader(body))}

		defer httpResponse.Body.Close()

		responseBody, _ := ioutil.ReadAll(httpResponse.Body)

		apiResponse := &Response{
			RawResponse: httpResponse,
			body: responseBody,
		}

		sut := &BaseClient{}

		// Act
		actual := &SampleType{}
		err := sut.parseResponseBody(apiResponse, nil, actual)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	})
}