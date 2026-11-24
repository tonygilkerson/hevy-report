# Hevy Routine Endpoint

GET v1/workouts - Get a paginated list of workouts

## Parameters

* **Auth:** Header `api-key: $HEVY_API_KEY`
* **Query params:**
  * `page` (int, default 1) - Page number (Must be 1 or greater)
  * `pageSize` (int, default 10) - Number of items on the requested page (Max 10)
* **Fixture:** `testdata/workouts.json`

## Sample Request

```sh
curl -X 'GET' -H "api-key: $HEVY_API_KEY" "https://api.hevyapp.com/v1/workouts?page=1&pageSize=10"
```

## Key Response Fields

* `page` (int) - Current page number
* `page_count` (int) - Total number of pages
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
