### Hevy CLI Report Tool (`hevy-report`) — Project Roadmap


#### Phase 3: Milestone 1 — Hevy HTTP Client (`./plans/01-hevy-client.md`)

* Build base HTTP client in `internal/hevy/`:
* Base URL `[https://api.hevyapp.com/v1](https://api.hevyapp.com/v1)`
* Auth header `api-key: $HEVY_API_KEY`
* `context.Context` support, client timeout, structured HTTP status error wrapping, and pagination helper


* Implement typed Go structs and `GetRoutines`, `GetWorkouts`, and `GetExerciseHistory`.
* Write table-driven unit tests using `net/http/httptest` and `testdata/*.json`.

#### Phase 4: Milestone 2 — Report Formatters (`./plans/02-reports.md`)

* Implement `text/tabwriter` human-readable formatters + `--json` output in `internal/report/`:
* Routines summary (title, exercises, target sets/reps/rest)
* Workouts summary (date, title, duration, exercises, top set / total volume)
* Exercise history progression (date, sets x reps @ weight, estimated 1RM or volume trend)


* Add unit tests verifying formatted text and JSON outputs against `testdata/`.

#### Phase 5: Milestone 3 — CLI Wiring & Polish (`./plans/03-cli.md`)

* Wire subcommands and flags in `cmd/hevy-report/` to `internal/hevy` and `internal/report`.
* Add exit-code handling, helpful `--help` usage text, and a brief `README.md` with build/run examples.

---

### Reminder for Each Coding Session on 24GB

* Run `/clear` between milestones.
* Point the model only to the **1 plan file**, **1 doc file**, and **1 fixture** needed for that step.