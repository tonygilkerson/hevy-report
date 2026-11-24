# Hevy Routine Endpoint

GET v1/workouts/{workoutId} - Get a workout by its id

## Parameters

* **workoutId:** (path) - The id of the workout

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
    * `index` (int)
    * `title` (string)
    * `notes` (string)
    * `exercise_template_id` (string) - example: 05293BCA The id of the exercise template. This can be used to fetch the exercise template.
    * `superset_id` (number|null)
    * `sets` (array)
      * `index` (int)
      * `type` (string) - The type of set. This can be one of 'normal', 'warmup', 'dropset', 'failure'
      * `weight_kg` (number|null)
      * `reps` (number|null)
      * `rpe` (number|null)
      * `distance_meters` (number|null)
      * `duration_seconds` (number|null)
      * `custom_metric` (number|null) - Currently only used to log floors or steps for stair machine exercises
