# Exercise & Set Fields

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
