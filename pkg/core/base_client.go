package core

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client interface{
	Send(apiReq *Request) (*Response, error)
}

const(
	defaultTimeoutInMilliseconds int = 500
)

type BaseClient struct {
	BaseUrl	url.URL
	UserAgent  string
	Timeout	int
	HttpClient HTTPClient
}

// Send makes the http request and returns a Response or error.
func (c *BaseClient) Send(apiReq *Request) (*Response, error) {

	httpReq, err := apiReq.buildHttpRequest(c.BaseUrl)

	if(err != nil){
		return nil, fmt.Errorf("Error while creating http request: %v", err)
	}

	httpReq.Header.Set("User-Agent", c.UserAgent)

	currentContext := apiReq.getContext()

	timeoutToApply := defaultTimeoutInMilliseconds
	if c.Timeout != 0{
		timeoutToApply = c.Timeout
	}

	ctx, cancel := context.WithTimeout(currentContext, time.Duration(timeoutToApply)*time.Millisecond)

	defer cancel()

	httpReq = httpReq.WithContext(ctx)

	resp, err := c.HttpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("Error while executing http request: %v", err)
	}


	apiResponse, err := c.handleHttpResponse(apiReq, resp)

	if err != nil{
		return nil, fmt.Errorf("Error while reading http response body: %v", err)
	}

	return apiResponse, nil
}

//Handles the http response by reading the body and returns a Response or error
func (c *BaseClient) handleHttpResponse(apiReq *Request, resp *http.Response) (*Response, error) {
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	apiResponse := &Response{
		RawResponse: resp,
		body: responseBody,
	}

	err = c.parseResponseBody(apiResponse, apiReq.Result, apiReq.Error)

	if(err != nil){
		return nil , err
	}

	return apiResponse, nil
}

// Parses the body of the http response
// If the response indicates success, parses the body into resultValue
// If the response does not indicate success, parses the body into errorValue
func (c *BaseClient) parseResponseBody(apiResponse *Response, resultValue interface{}, errorValue interface{}) error {
	if apiResponse.IsSuccess() && resultValue != nil{
		return apiResponse.UnmarshalJson(resultValue);

	} else if apiResponse.IsError() && errorValue != nil{
		return apiResponse.UnmarshalJson(errorValue)
	}

	return nil
}
