# hevyai

Hevy CLI built with local LLM

## Dev Setup

```sh
# Start inference serving in separate terminal
unsloth run --model unsloth/Qwen3-Coder-30B-A3B-Instruct-GGUF:Q4_K_M \
    -c 49152 \
    -fa 1 \
    --parallel 1 \
    --load-mode none \
    --cache-type-k q4_0 \
    --cache-type-v q4_0 \
    --temp 0.7 \
    --top-p 0.8 \
    --top-k 20 \
    --repeat-penalty 1.05 \
    --reasoning off \
    --no-mmproj

# Start coding agent
unsloth start pi
```

Use **Unsloth Desktop App** to chat out plans and use web search, use the pi coding agent for dev tasks in VSCode.

## Prompt Examples

* Implement Step 2 of `./plans/workouts-cmd.md` and run `go test ./...`
* Implement `ListWorkouts` in `internal/hevy` and add a table-driven test using `httptest` and `testdata/workouts.json`. Inspect only the struct fields needed for workouts pagination and set/exercise summaries.
* Read `./docs/routine.md` and `testdata/routines.json` before implementing routine fetching.