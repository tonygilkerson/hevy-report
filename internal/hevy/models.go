package hevy

import (
	"time"
)

// RoutinePagination represents the pagination information for routines.
type RoutinePagination struct {
	Page      int `json:"page"`
	PageCount int `json:"page_count"`
}

// RoutineResponse represents a response containing routines.
type RoutineResponse struct {
	RoutinePagination
	Routines []Routine `json:"routines"`
}

// RoutineFolderPagination represents the pagination information for routine folders.
type RoutineFolderPagination struct {
	Page      int `json:"page"`
	PageCount int `json:"page_count"`
}

// RoutineFolderResponse represents a response containing routine folders.
type RoutineFolderResponse struct {
	RoutineFolderPagination
	RoutineFolders []RoutineFolder `json:"routine_folders"`
}

// WorkoutPagination represents the pagination information for workouts.
type WorkoutPagination struct {
	Page      int `json:"page"`
	PageCount int `json:"page_count"`
}

// WorkoutResponse represents a response containing workouts.
type WorkoutResponse struct {
	WorkoutPagination
	Workouts []Workout `json:"workouts"`
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
