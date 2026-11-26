package hevy

import (
	"time"
)

// RoutinePagination represents the pagination information for routines.
type RoutinePagination struct {
	Page      int `json:"page"`
	PageCount int `json:"page_count"`
}

// GetPage returns the current page number.
func (r RoutinePagination) GetPage() int {
	return r.Page
}

// GetPageCount returns the total number of pages.
func (r RoutinePagination) GetPageCount() int {
	return r.PageCount
}

// RoutineResponse represents a response containing routines.
type RoutineResponse struct {
	RoutinePagination
	Routines []Routine `json:"routines"`
}

// GetPage returns the current page number.
func (r RoutineResponse) GetPage() int {
	return r.Page
}

// GetPageCount returns the total number of pages.
func (r RoutineResponse) GetPageCount() int {
	return r.PageCount
}

// RoutineFolderPagination represents the pagination information for routine folders.
type RoutineFolderPagination struct {
	Page      int `json:"page"`
	PageCount int `json:"page_count"`
}

// GetPage returns the current page number.
func (r RoutineFolderPagination) GetPage() int {
	return r.Page
}

// GetPageCount returns the total number of pages.
func (r RoutineFolderPagination) GetPageCount() int {
	return r.PageCount
}

// RoutineFolderResponse represents a response containing routine folders.
type RoutineFolderResponse struct {
	RoutineFolderPagination
	RoutineFolders []RoutineFolder `json:"routine_folders"`
}

// GetPage returns the current page number.
func (r RoutineFolderResponse) GetPage() int {
	return r.Page
}

// GetPageCount returns the total number of pages.
func (r RoutineFolderResponse) GetPageCount() int {
	return r.PageCount
}

// WorkoutPagination represents the pagination information for workouts.
type WorkoutPagination struct {
	Page      int `json:"page"`
	PageCount int `json:"page_count"`
}

// GetPage returns the current page number.
func (w WorkoutPagination) GetPage() int {
	return w.Page
}

// GetPageCount returns the total number of pages.
func (w WorkoutPagination) GetPageCount() int {
	return w.PageCount
}

// WorkoutResponse represents a response containing workouts.
type WorkoutResponse struct {
	WorkoutPagination
	Workouts []Workout `json:"workouts"`
}

// GetPage returns the current page number.
func (w WorkoutResponse) GetPage() int {
	return w.Page
}

// GetPageCount returns the total number of pages.
func (w WorkoutResponse) GetPageCount() int {
	return w.PageCount
}

// Routine represents a routine in Hevy.
type Routine struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	FolderID  *int       `json:"folder_id"` // nullable
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
	Exercises []Exercise `json:"exercises"`
}

// RoutineFolder represents a routine folder in Hevy.
type RoutineFolder struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Workout represents a workout in Hevy.
type Workout struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	RoutineID   *string    `json:"routine_id"`  // nullable
	Description *string    `json:"description"` // nullable
	StartTime   time.Time  `json:"start_time"`
	EndTime     time.Time  `json:"end_time"`
	UpdatedAt   time.Time  `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
	Media       []any      `json:"media"` // Not Used
	Exercises   []Exercise `json:"exercises"`
}

// Exercise represents an exercise in Hevy.
type Exercise struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Set  []Set  `json:"sets"`
}

// Set represents a set in Hevy.
type Set struct {
	ID        string   `json:"id"`
	Weight    *float64 `json:"weight"`   // nullable
	Reps      *int     `json:"reps"`     // nullable
	Duration  *float64 `json:"duration"` // nullable
	Completed bool     `json:"completed"`
}
