# Hevy CLI Design (`hevy-report`)

CLI specification for querying Hevy routines and workouts and rendering human-readable or JSON reports.

## Subcommands

- `hevy-report routines` — List saved routines and their exercises/sets (`GET /v1/routines`).
- `hevy-report workouts` — List logged workouts (`GET /v1/workouts`), or fetch a single workout by ID (`hevy-report workouts <id>`).

## Flags

### Shared Flags

- `--json` (`bool`, default `false`): Output structured JSON instead of human-readable text.
- `--units` (`string`, default `lbs`): Display weight units (`lbs` or `kg`); convert from `weight_kg`.
- `--detail` (`string`, default `detailed`): Output verbosity level: `list`, `compact`, `detailed`, or `expanded`.

### Pagination & Filtering

- `--page` (`int`, default `1`): Starting page number (must be `>= 1`).
- `--page-size` (`int`, default `10`): Page size sent to Hevy API (respect API max).
- `--all` (`bool`, default `false`): Fetch pages until `page >= page_count` (with a safety cap).
- `--since` (`string`, optional): Filter workouts on or after a date (`YYYY-MM-DD`) or RFC3339 timestamp.

## Report Layouts (`routines`)

### Rendering Rules

- **Folder name:** Resolve `folder_id` via `GET /v1/routine_folders` (omit `(folder: ...)` if `folder_id` is `null`).
- **Supersets:** Exercises with a non-null `superset_id` render with `- | <exercise name>`.
- **Set labels:** `warmup` -> `W`, `dropset` -> `D`, `failure` -> `F`, `normal` -> 1-based working set counter (`1`, `2`, `3`).
- **Null handling:** If `weight_kg` is `null`, omit weight; if `reps` is `null`, render `...`.
- **Notes (`expanded`):** Prefix each non-empty line of routine or exercise `notes` with `( `.

### Level Templates

- `--detail list`: `<routine name> (folder: <folder name>)`
- `--detail compact`: routine line + `- <exercise name>` (or `- | <exercise name>`)
- `--detail detailed` (default): `compact` + indented set rows (`<set label>  <weight> <units> X <reps> reps`)
- `--detail expanded`: `detailed` + routine notes under routine line and exercise notes under exercise line

### Concrete Example (`--detail expanded`)

```text
Upper A (folder: YMCA)
( focus on upper body
( focus on slow full range during exercises
- Lat Pulldown (Cable)
    ( Go slow on this one, push out chest
    W   60 lbs X 25 reps
    W   60 lbs X 20 reps
    1   75 lbs X 12 reps
    2   90 lbs X 10 reps
    3  105 lbs X  8 reps
- | Crunch
    1  X 25 reps
- | Weighted Snow Angels
    ...
```

## Report Layouts (`workouts`)

TBD

## Output Behavior

- Human-readable mode uses `text/tabwriter` and writes to a provided `io.Writer` (defaulting to `os.Stdout` in `cmd/hevy-report`).
- Guard against `nil` / null fields (`weight_kg`, `reps`, `duration_seconds`, `rest_seconds`) without panicking or fabricating values.
- Write human/JSON reports to stdout; write errors and diagnostics to stderr with non-zero exit codes on failure.

