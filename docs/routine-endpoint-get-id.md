# Hevy Routine Endpoint

GET v1/routines/{routineId} - Get a routine by its id

## Parameters

* **routineId:** (path) - The id of the routine

## Sample Request

```sh
curl -X 'GET' -H "api-key: $HEVY_API_KEY" "https://api.hevyapp.com/v1/routines/6a89e153-d5f2-4015-91c1-d741dc8902e8"
```

## Key Response Fields

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
