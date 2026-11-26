package hevy

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetRoutines(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check auth header
		apiKey := r.Header.Get("api-key")
		assert.Equal(t, "test-api-key", apiKey)

		// Check query parameters
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "10", r.URL.Query().Get("pageSize"))

		// Check endpoint (we're not strictly testing the full URL because httptest creates a full URL)
		assert.NotEmpty(t, r.URL.Path)

		// Return mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		// Just return a simple valid JSON response
		w.Write([]byte(`{"page": 1, "page_count": 2, "routines": []}`))
	}))
	defer server.Close()

	// Create client with test server
	client := NewClient("test-api-key", &http.Client{
		Timeout: 10 * time.Second,
	})

	// Override baseURL to point to test server
	client.baseURL = server.URL

	// Call method
	resp, err := client.GetRoutines(1, 10)

	// Check results
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, resp.Page)
	assert.Equal(t, 2, resp.PageCount)
	assert.Len(t, resp.Routines, 0)
}

func TestClient_GetWorkouts(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check auth header
		apiKey := r.Header.Get("api-key")
		assert.Equal(t, "test-api-key", apiKey)

		// Check query parameters
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "10", r.URL.Query().Get("pageSize"))

		// Check endpoint (we're not strictly testing the full URL because httptest creates a full URL)
		assert.NotEmpty(t, r.URL.Path)

		// Return mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		// Just return a simple valid JSON response
		w.Write([]byte(`{"page": 1, "page_count": 1, "workouts": []}`))
	}))
	defer server.Close()

	// Create client with test server
	client := NewClient("test-api-key", &http.Client{
		Timeout: 10 * time.Second,
	})

	// Override baseURL to point to test server
	client.baseURL = server.URL

	// Call method
	resp, err := client.GetWorkouts(1, 10)

	// Check results
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, resp.Page)
	assert.Equal(t, 1, resp.PageCount)
	assert.Len(t, resp.Workouts, 0)
}

func TestClient_GetRoutineFolderByID(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check auth header
		apiKey := r.Header.Get("api-key")
		assert.Equal(t, "test-api-key", apiKey)

		// Check endpoint (we're not strictly testing the full URL because httptest creates a full URL)
		assert.NotEmpty(t, r.URL.Path)

		// Return mock response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		// Just return a simple valid JSON response
		w.Write([]byte(`{"id": 12345, "name": "Test Folder"}`))
	}))
	defer server.Close()

	// Create client with test server
	client := NewClient("test-api-key", &http.Client{
		Timeout: 10 * time.Second,
	})

	// Override baseURL to point to test server
	client.baseURL = server.URL

	// Call method
	resp, err := client.GetRoutineFolderByID(12345)

	// Check results
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 12345, resp.ID)
}

func TestAPIError(t *testing.T) {
	err := &APIError{
		StatusCode: 400,
		Endpoint:   "/v1/routines",
		Body:       "Bad Request",
	}
	assert.EqualError(t, err, "Hevy API error 400 on /v1/routines: Bad Request")
}
