# Milestone 2: Report Formatting & CLI Wiring

* [ ] **Step 1 - Report options & JSON output (`internal/report/`):** Define an `Options` struct (`JSON`, `Units`, `Detail`) and formatter entry points writing to an `io.Writer` for routines and workouts, implementing indented `--json` output.

* [ ] **Step 2 - Text report formatting routines (`internal/report/`):** Implement concrete text formatting functions or methods for routines using shared exercise/set rules, `null`-safe optional fields.

* [ ] **Step 3 - Text report formatting workouts (`internal/report/`):** Implement concrete text formatting functions or methods (not just an interface) for workouts using shared exercise/set rules, `null`-safe optional fields, and direct workout `title` headers without extra ID lookups.

* [ ] **Step 4 - Report unit tests (`internal/report/`):** Add table-driven tests covering `--json`, unit conversions, detail levels, nullable set fields, and workout title rendering.

* [ ] **Step 5 - CLI flags & command wiring (`cmd/hevy-report/`):** Implement CLI flag/command parsing per `docs/cli-design.md`, wire `internal/hevy` and `internal/report`, and map errors to clean messages and exit codes.

* [ ] **Step 6 - CLI tests & verification (`cmd/hevy-report/`):** Add CLI flag/command tests against a test server and run `go test ./...` and `go vet ./...`.