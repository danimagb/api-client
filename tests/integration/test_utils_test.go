package tests

import (
	"net/url"
	"os"
	"testing"

	client "github.com/danimagb/api-client/pkg"
)

const(
	defaultHost string = "http://localhost:8080/"
)


func SetupNewClient(t *testing.T) *client.Client{
	host := os.Getenv("API_URL")
	if len(host) == 0{
		host = defaultHost
	}

	u, err := url.Parse(host)
	if err != nil {
		t.Errorf("Error parsing url to create client: %v", err)
	}

	client, err := client.NewClient(
		client.WithBaseUrl(*u),
		client.WithTimeoutInMilliseconds(1000),
	)

	if err != nil {
		t.Errorf("Error creating client client: %v", err)
	}

	return client
}