# Milestone 1: Hevy HTTP Client (`internal/hevy`)

Scope: Routines and workouts only (`GET /v1/routines`, `GET /v1/routine_folders`, `GET /v1/workouts`, `GET /v1/workouts/{id}`).

## Steps

* [X] **Step 1 - Client & error types (`client.go`):** Define `Client` with base URL (`https://api.hevyapp.com/v1`), `api-key` header auth, configurable `*http.Client` timeout, `context.Context` on all requests, and an `APIError` struct (`StatusCode`, `Endpoint`, `Body`) implementing `error` and `Unwrap`.
* [X] **Step 2 - Models (`models.go`):** Define typed structs for routine pagination responses, routines, routine folders, workouts, exercises, and sets using pointer/nullable types where JSON fields can be `null`.
* [X] **Step 3 - Pagination helper (`client.go` or `paginate.go`):** Implement a page-fetching helper supporting `page`, `pageSize`, and an `all` mode that loops until `page >= page_count` with a safety cap.
* [X] **Step 4 - Endpoint methods (`routines.go`):** Implement `GetRoutines`, `GetRoutineFolderByID`.
* [X] **Step 5 - Endpoint methods (`workouts.go`):** Implement `GetWorkouts`.
* [ ] **Step 6 - Unit tests (`client_test.go`):** Add table-driven tests using `httptest.NewServer` and `testdata/*.json` covering auth headers, query params, pagination, `null` fields, and HTTP non-2xx errors.
