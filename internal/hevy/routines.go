package hevy

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
