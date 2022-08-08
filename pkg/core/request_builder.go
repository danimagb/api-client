package core

import (
	"context"
	"net/url"
	"strings"
)

// RequestBuilder is the main interface with some methods to build and handle a request.
type RequestBuilder interface {
	WithPath(value string ) RequestBuilder
	WithBody(body interface{}) RequestBuilder
	WithQueryParam(key, value string) RequestBuilder
	WithResultWriteTo(value interface{}) RequestBuilder
	WithErrorWriteTo(value interface{}) RequestBuilder
	WithContext(context context.Context) RequestBuilder
	Build() (*Request)
}

type requestBuilderImpl struct {
	httpMethod      	string
	path 				[]string
	queryParam 			url.Values
	body            	interface{}
	resultWriter     	interface{}
	errorWriter      	interface{}
	context 			context.Context
}

func NewRequestBuilder(method string) *requestBuilderImpl {
	return &requestBuilderImpl{
		httpMethod: method,
		queryParam: url.Values{},
	}
}

func (r requestBuilderImpl) WithPath(value string) RequestBuilder{
	r.path = append(r.path, value)
	return &r
}

func (r requestBuilderImpl) WithBody(body interface{}) RequestBuilder{
	r.body = body
	return &r
}

func (r requestBuilderImpl) WithQueryParam(param, value string) RequestBuilder{
	r.queryParam.Set(param, value)
	return &r
}

func (r requestBuilderImpl) WithContext(value context.Context) RequestBuilder{
	r.context = value
	return &r
}

func (r requestBuilderImpl) WithResultWriteTo(value interface {}) RequestBuilder{
	r.resultWriter = value
	return &r
}

func (r requestBuilderImpl) WithErrorWriteTo(value interface {}) RequestBuilder{
	r.errorWriter = value
	return &r
}

func (r requestBuilderImpl) Build() (*Request){
	finalPath := strings.Join(r.path, "/")

	return &Request{
		Method: r.httpMethod,
		Path: finalPath,
		QueryParam: r.queryParam,
		Body: r.body,
		Result: r.resultWriter,
		Error: r.errorWriter,
		Context: r.context,
	}
}