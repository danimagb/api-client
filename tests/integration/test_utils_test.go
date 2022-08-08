package tests

import (
	"net/url"
	"os"
	"testing"

	form3 "github.com/danimagb/api-client/pkg"
)

const(
	defaultHost string = "http://localhost:8080/"
)


func SetupNewClient(t *testing.T) *form3.Form3{
	host := os.Getenv("API_URL")
	if len(host) == 0{
		host = defaultHost
	}

	u, err := url.Parse(host)
	if err != nil {
		t.Errorf("Error parsing url to create form3 client: %v", err)
	}

	client, err := form3.NewClient(
		form3.WithBaseUrl(*u),
		form3.WithTimeoutInMilliseconds(30),
	)

	if err != nil {
		t.Errorf("Error creating form3 client: %v", err)
	}

	return client
}