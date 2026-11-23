# Hevy Routines

Hevy API host used in all API endpoints: `https://api.hevyapp.com`

## Routine Endpoints

### GET /v1/routines

Get a list of paginated list of routines

Parameters

* **Endpoint:** `GET https://api.hevyapp.com/v1/routines`
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

* Top-level
  * `page` (int) - Current page number
  * `page_count` (int) - Total number of pages
  * `routines` (array)
        * `id` (string)
        * `title` (string)
        * `folder_id` (int|null)
        * `notes` (string)
        * `updated_at` (ISO 8601 timestamp)
        * `created_at` (ISO 8601 timestamp)
        * `exercises` (array)
          * `index` (int)
          * `title` (string)
          * `exercise_template_id` (string)
          * `notes` (string)
          * `rest_seconds` (int|null)
          * `sets` (array)
            * `index` (int)
            * `type` (string)
            * `weight_kg` (float|null)
            * `reps` (number|null)
