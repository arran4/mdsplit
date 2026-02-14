# Upgrade Report: go-subcommand v0.0.17

## Summary
The upgrade to `github.com/arran4/go-subcommand` v0.0.17 introduces several changes and a critical bug that prevents the generated code from compiling without manual intervention.

## Issues Found

### 1. Compilation Error: Missing Package Qualification
The generated `cmd/mdsplit/root.go` file contains a call to `Run(...)` instead of `mdsplit.Run(...)`. The `mdsplit` package is imported correctly as `github.com/arran4/mdsplit`, but the function call lacks the package prefix.

**Error:**
```
cmd/mdsplit/root.go:115:10: undefined: Run
```

**Workaround:**
Manually edit `cmd/mdsplit/root.go` to change `Run(...)` to `mdsplit.Run(...)`. However, this file is overwritten by `go generate`.

### 2. `go generate` Execution Context
Running `go generate ./...` from the root directory fails because the generated command in `cmd/mdsplit/main.go` attempts to run `gosubc generate` (or equivalent) which seems to expect `go.mod` in the current working directory (`cmd/mdsplit`), but `go.mod` is in the root.

**Error:**
```
generate failed: go.mod not found in the root of the repository: open go.mod: no such file or directory
```

**Workaround:**
Run `go run github.com/arran4/go-subcommand/cmd/gosubc generate ./cmd/mdsplit` from the root directory.

### 3. New Files Created
- `cmd/errors.go`: Created in `cmd` package. This introduces a new package `cmd` in the repository.
- `cmd/mdsplit/root_test.go`: A new test file for the root command.

## Verification
After applying the manual fix for the compilation error, the following verification steps passed:
- `go test ./...` passed.
- CLI build (`go build ./cmd/mdsplit`) succeeded.
- CLI execution (`./mdsplit --help` and usage on a test file) worked as expected.

## Recommendation
The upgrade cannot be completed successfully without fixing the code generation bug in `gosubc` or applying a patch to the generated code post-generation.
