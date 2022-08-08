package core

import (
	"fmt"
)


type ApiClientError struct {
	Reason 		string
	StatusCode 	int
	Message    	string
	Response    *Response
}

func NewApiClientError(reason string, statusCode int, message string, response *Response) *ApiClientError {
	return &ApiClientError{
		Reason : reason,
		StatusCode: statusCode,
		Message: message,
		Response: response,
	}
}

func (clientError *ApiClientError) Error() string {
	return fmt.Sprintf("%s (Status Code: %d | Message: '%s')", clientError.Reason, clientError.StatusCode, clientError.Message)
}