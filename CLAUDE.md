# Hevy CLI Report Tool

Go CLI that queries the Hevy REST API and prints human-readable reports for workouts, routines, and exercise history.

## Commands

- Build: `go build ./cmd/hevy-report`
- Test: `go test ./...`
- Check: `go vet ./... && gofmt -l .`

## Project Layout

- `docs/`: Concise per-endpoint API notes and curl examples
- `testdata/`: Small redacted JSON fixtures for unit tests
- `cmd/hevy-report/`: CLI entry point, subcommands, and flags
- `internal/hevy/`: Hevy API client, typed models, pagination
- `internal/report/`: Human-readable and JSON report formatters

## Go & API Conventions

- Prefer standard library (`net/http`, `encoding/json`, `text/tabwriter`).
- Read API key from `HEVY_API_KEY`; never hardcode or log secrets.
- Always pass `context.Context`, set an HTTP client timeout, and use typed structs with `json` tags (no `map[string]any`).
- Wrap errors with endpoint/HTTP status context (`fmt.Errorf("...: %w", err)`).
- Handle pagination and rate limits; default CLI output to human-readable tables, with a `--json` flag.

## Workflow & Testing

- Keep edits scoped to the active task; write a short checklist in `./plans/<task>.md` before multi-step work and check items off as tests pass.
- Do not paste large OpenAPI specs or full fixtures into context; inspect specific fields on demand.
- Write table-driven tests using `httptest` and `testdata/` fixtures (no live API calls). Run `gofmt`, `go vet`, and `go test ./...` before finishing.

## Chat Style

- Keep chat replies to 1–2 sentences (or just "Done.") with no code walkthroughs, step narration, or summaries unless asked.