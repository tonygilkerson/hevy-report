package hevy

import (
	"fmt"
	"net/http"
	"time"
)

// Client is a Hevy API client.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewClientWithHTTPClient creates a new Hevy API client with a custom HTTP client.
func NewClient(
	// apiKey Used in auth header to access Hevy API
	apiKey string,

	// httpClient An http.Client used to access Hevy API
	// If nil a default client is created
	httpClient *http.Client,
) *Client {

	// Create a default if none was provided
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 10 * time.Second}
	}

	return &Client{
		baseURL:    "https://api.hevyapp.com/v1",
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

// APIError represents an error returned by the Hevy API.
type APIError struct {
	StatusCode int
	Endpoint   string
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Hevy API error %d on %s: %s", e.StatusCode, e.Endpoint, e.Body)
}

func (e *APIError) Unwrap() error {
	return fmt.Errorf("HTTP status %d", e.StatusCode)
}

