package hevy

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetWorkouts gets a paginated list of workouts.
func (c *Client) GetWorkouts(page, pageSize int) (*WorkoutResponse, error) {
	endpoint := fmt.Sprintf("%s/workouts?page=%d&pageSize=%d", c.baseURL, page, pageSize)

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

	var response WorkoutResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
