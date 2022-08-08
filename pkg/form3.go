package form3

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/danimagb/api-client/pkg/accounts"
	"github.com/danimagb/api-client/pkg/core"
)

const (
	defaultBaseURL   = "https://api.form3.tech/"
)

type Form3 struct{
	httpClient *http.Client
	baseUrl	url.URL
	userAgent string
	timeout int
	Accounts *accounts.AccountsClient
}

type ClientOption func (*Form3) error


func NewClient(options ... ClientOption) (*Form3, error){
	httpClient := &http.Client{}

	baseUrl,_ := url.Parse(defaultBaseURL)

	form3 := &Form3{
		httpClient: httpClient,
		baseUrl: *baseUrl,
	}

	for _, option := range options{
		err := option(form3)

		if(err != nil){
			return nil, fmt.Errorf("error when creating Form3 Client %w", err)
		}
	}

	baseClient := &core.BaseClient{
		BaseUrl: form3.baseUrl,
		UserAgent: form3.userAgent,
		HttpClient: form3.httpClient,
		Timeout: form3.timeout,
	}


	form3.Accounts = accounts.New(baseClient)

	return form3, nil
}

func WithHttpClient(c *http.Client) ClientOption{
	return func(form3 *Form3) error {
		if c != nil{
			form3.httpClient = c
		}
		return nil
	}
}

func WithBaseUrl(url url.URL) ClientOption{
	return func(form3 *Form3) error {
		form3.baseUrl = url
		return nil
	}
}

func WithUserAgent(userAgent string) ClientOption{
	return func(form3 *Form3) error {
		form3.userAgent = userAgent
		return nil
	}
}

func WithTimeoutInMilliseconds(timeout int) ClientOption{
	return func(form3 *Form3) error {
		if timeout > 0{
			form3.timeout = timeout
			return nil
		}
		return fmt.Errorf("timeout must be greater than zero (actual timeout: %d)", timeout)
	}
}