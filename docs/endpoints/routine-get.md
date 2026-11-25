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
    * See `docs/models/exercise.md` for `exercise` and `set` fields
