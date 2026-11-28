# Prompt

Give Pi each prompt one at a time. Be sure enter the `/new` command before each prompt to start a fresh session.

## Prompt 1

Read the following:

* plans/02-report-and-cli.md
* docs/cli-design.md
* internal/hevy/models.go
* internal/hevy/routines.go
* internal/hevy/workouts.go

Then perform the following tasks:

* Implement only Step 1 in `internal/report/`.
* Import `github.com/tonygilkerson/hevy-report/internal/hevy` and use `hevy.*` types directly do not redefine those models or create extra JSON wrapper structs.
* Return an error if `opts.Writer` is nil.
* Stop once `go build ./internal/report` succeeds.

## Prompt 2

Read the following:

* plans/02-report-and-cli.md
* docs/cli-design.md
* internal/hevy/models.go
* internal/hevy/routines.go

Then perform the following tasks:

* Implement only Step 2 in `internal/report/`.
* Stop once `go build ./internal/report` succeeds.
