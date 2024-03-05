package client

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/danimagb/api-client/pkg/accounts"
	"github.com/danimagb/api-client/pkg/core"
)

const (
	defaultBaseURL   = "https://api.dummy.tech/"
)

type Client struct{
	httpClient *http.Client
	baseUrl	url.URL
	userAgent string
	timeout int
	Accounts *accounts.AccountsClient
}

type ClientOption func (*Client) error


func NewClient(options ... ClientOption) (*Client, error){
	httpClient := &http.Client{}

	baseUrl,_ := url.Parse(defaultBaseURL)

	client := &Client{
		httpClient: httpClient,
		baseUrl: *baseUrl,
	}

	for _, option := range options{
		err := option(client)

		if(err != nil){
			return nil, fmt.Errorf("error when creating Client %w", err)
		}
	}

	baseClient := &core.BaseClient{
		BaseUrl: client.baseUrl,
		UserAgent: client.userAgent,
		HttpClient: client.httpClient,
		Timeout: client.timeout,
	}


	client.Accounts = accounts.New(baseClient)

	return client, nil
}

func WithHttpClient(c *http.Client) ClientOption{
	return func(client *Client) error {
		if c != nil{
			client.httpClient = c
		}
		return nil
	}
}

func WithBaseUrl(url url.URL) ClientOption{
	return func(client *Client) error {
		client.baseUrl = url
		return nil
	}
}

func WithUserAgent(userAgent string) ClientOption{
	return func(client *Client) error {
		client.userAgent = userAgent
		return nil
	}
}

func WithTimeoutInMilliseconds(timeout int) ClientOption{
	return func(client *Client) error {
		if timeout > 0{
			client.timeout = timeout
			return nil
		}
		return fmt.Errorf("timeout must be greater than zero (actual timeout: %d)", timeout)
	}
}