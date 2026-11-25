# Hevy Routine Endpoint

GET v1/workouts/{workoutId} - Get a workout by its id

## Parameters

* **workoutId:** (path) - The id of the workout
* **Auth:** Header `api-key: $HEVY_API_KEY`
* **Fixture:** `testdata/workout.json`

## Sample Request

```sh
curl -X 'GET' -H "api-key: $HEVY_API_KEY" "https://api.hevyapp.com/v1/workouts/077c165b-09a3-4585-9fd2-04dbe776e3cd"
```

## Key Response Fields

* `workouts` (array)
  * `id` (string)
  * `title` (string)
  * `routine_id` (string|null)
  * `description` (string|null)
  * `start_time` (ISO 8601 timestamp)
  * `end_time` (ISO 8601 timestamp)
  * `updated_at` (ISO 8601 timestamp)
  * `created_at` (ISO 8601 timestamp)
  * 'media' (array) - Not Used
  * `exercises` (array)
    * See `docs/models/exercise.md` for `exercise` and `set` fields
