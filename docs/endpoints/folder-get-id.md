# Hevy Folder Endpoint

GET v1/routine_folders/{folderId} - Get a routine folder by its id

## Parameters

* **workoutId:** (path) - The id of the workout
* **Auth:** Header `api-key: $HEVY_API_KEY`
* **Fixture:** `testdata/folder.json`

## Sample Request

```sh
curl -X 'GET' -H "api-key: $HEVY_API_KEY" "https://api.hevyapp.com/v1/routine_folders/3075159"
```

## Response Fields

* See `docs/models/folder.md` for `folder` fields
