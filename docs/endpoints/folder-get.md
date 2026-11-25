# Hevy Folder Endpoint

GET v1/routine_folders - Get a paginated list of routine folders

## Parameters

* **Auth:** Header `api-key: $HEVY_API_KEY`
* **Query params:**
  * `page` (int, default 1) - Page number (Must be 1 or greater)
  * `pageSize` (int, default 10) - Number of items on the requested page (Max 10)
* **Fixture:** `testdata/folders.json`

## Sample Request

```sh
curl -X 'GET' -H "api-key: $HEVY_API_KEY" "https://api.hevyapp.com/v1/routine_folders?page=1&pageSize=10"
```

## Response Fields

* `page` (int) - Current page number
* `page_count` (int) - Total number of pages
* `routine_folders` (array)
  * See `docs/models/folder.md` for `folder` fields
