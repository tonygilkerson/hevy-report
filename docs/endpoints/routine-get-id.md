# Hevy Routine Endpoint

GET v1/routines/{routineId} - Get a routine by its id

## Parameters

* **routineId:** (path) - The id of the routine
* **Auth:** Header `api-key: $HEVY_API_KEY`
* **Fixture:** `testdata/routine.json`

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
    * See `docs/models/exercise.md` for `exercise` and `set` fields
