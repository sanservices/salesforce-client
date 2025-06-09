package salesforce

import (
	"net/http"
	"time"
)

type Config struct {
	Url    string `yaml:"Url"`
	ApiKey string `yaml:"ApiKey"`
}
type Client struct {
	Config     *Config
	HTTPClient *http.Client
}

// New returns service instance
func New(config *Config) *Client {
	return &Client{
		HTTPClient: createHTTPClient(),
		Config:     config,
	}
}

// Function to create Http Client
func createHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        300,
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     100,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}
