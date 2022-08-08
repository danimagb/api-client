package core

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type Request struct {
	Method          string
	Path            string
	QueryParam 		url.Values
	Body            interface{}
	Result			interface{}
	Error  			interface{}
	Context         context.Context
}

func (r *Request) buildHttpRequest(baseUrl url.URL) (*http.Request, error){
	url, err := baseUrl.Parse(r.Path)

	if err != nil {
		return nil, err
	}

	url.RawQuery = r.QueryParam.Encode()

	body := []byte{}
	if r.Body != nil{
		body, err = json.Marshal(r.Body)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(r.Method, url.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	if request.Body != nil {
		request.Header.Set("Content-Type", "application/json")
    }
	request.Header.Set("Accept", "application/json")

	return request, nil
}

func (r *Request) getContext() context.Context {
	if r.Context == nil {
		return context.Background()
	}
	return r.Context
}
