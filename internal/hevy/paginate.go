package hevy

import (
	"context"
	"fmt"
)

// PaginationResponse defines the interface for pagination responses.
type PaginationResponse interface {
	GetPage() int
	GetPageCount() int
}

// FetchAllPages fetches all pages of a paginated response using the provided fetch function.
// It handles pagination by looping until page >= page_count with a safety cap.
func FetchAllPages(
	ctx context.Context, 
	fetch func(page, pageSize int) (PaginationResponse, error)
) ([]PaginationResponse, error) {
	var responses []PaginationResponse
	
	page := 1
	pageSize := 10 // Default page size
	safetyCap := 100 // Safety cap to prevent infinite loops
	
	for i := 0; i < safetyCap; i++ {
		resp, err := fetch(page, pageSize)
		if err != nil {
			return responses, fmt.Errorf("failed to fetch page %d: %w", page, err)
		}
		
		responses = append(responses, resp)
		
		// Check if we've reached the last page
		if page >= resp.GetPageCount() {
			break
		}
		
		page++
	}
	
	return responses, nil
}