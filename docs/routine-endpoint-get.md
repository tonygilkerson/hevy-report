# Hevy Routine Endpoint

GET v1/routines - Get a paginated list of routines

## Parameters

* **Auth:** Header `api-key: $HEVY_API_KEY`
* **Query params:**
  * `page` (int, default 1) - Page number (Must be 1 or greater)
  * `pageSize` (int, default 10) - Number of items on the requested page (Max 10)
* **Fixture:** `testdata/routines.json`

## Sample Request

```sh
curl -X 'GET' -H "api-key: $HEVY_API_KEY" "https://api.hevyapp.com/v1/routines?page=1&pageSize=10"
```

## Key Response Fields

* `page` (int) - Current page number
* `page_count` (int) - Total number of pages
* `routines` (array)
  * `id` (string)
  * `title` (string)
  * `folder_id` (int|null)
  * `updated_at` (ISO 8601 timestamp)
  * `created_at` (ISO 8601 timestamp)
  * `exercises` (array)
    * `index` (int)
    * `title` (string)
    * `rest_seconds` (int|null)
    * `notes` (string)
    * `exercise_template_id` (string) - example: 05293BCA The id of the exercise template. This can be used to fetch the exercise template.
    * `superset_id` (number|null)
    * `sets` (array)
      * `index` (int)
      * `type` (string) - The type of set. This can be one of 'normal', 'warmup', 'dropset', 'failure'
      * `weight_kg` (number|null)
      * `reps` (number|null)
      * `rep_range` (dict|null) - Not Used.
      * `distance_meters` (number|null)
      * `duration_seconds` (number|null)
      * `custom_metric` (number|null) - Currently only used to log floors or steps for stair machine exercises
