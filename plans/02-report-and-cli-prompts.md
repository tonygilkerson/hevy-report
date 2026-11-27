# Prompt

Give Pi each prompt one at a time. Be sure enter the `/new` command before each prompt to start a fresh session.

## Prompt 1

Read the following:

* plans/02-report-and-cli.md
* docs/cli-design.md
* internal/hevy/models.go
* internal/hevy/routines.go
* internal/hevy/workouts.go

Implement only Step 1 in `internal/report/`. Stop once `go build ./internal/report` succeeds.

Prompt 1 take 2:

Read the following:

- plans/02-report-and-cli.md
- docs/cli-design.md
- internal/hevy/models.go
- internal/hevy/routines.go
- internal/hevy/workouts.go

Implement only Step 1 in `internal/report/`. Import `github.com/tonygilkerson/hevy-report/internal/hevy` and use `hevy.Routine`, `hevy.RoutineFolder`, and `hevy.Workout`—do not redefine those models in `package report`. Stop once `go build ./internal/report` succeeds.