package hevy

import (
	"encoding/json"
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

// GetRoutineFolderByID gets a routine folder by its ID.
func (c *Client) GetRoutineFolderByID(id int) (*RoutineFolder, error) {
	endpoint := fmt.Sprintf("%s/routine_folders/%d", c.baseURL, id)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("api-key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Endpoint:   endpoint,
			Body:       resp.Status,
		}
	}

	var folder RoutineFolder
	if err := json.NewDecoder(resp.Body).Decode(&folder); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &folder, nil
}

// GetRoutines gets a paginated list of routines.
func (c *Client) GetRoutines(page, pageSize int) (*RoutineResponse, error) {
	endpoint := fmt.Sprintf("%s/routines?page=%d&pageSize=%d", c.baseURL, page, pageSize)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("api-key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Endpoint:   endpoint,
			Body:       resp.Status,
		}
	}

	var response RoutineResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
