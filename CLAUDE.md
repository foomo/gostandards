# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go library providing typed constants for common standards (ISO-4217 currency codes, ISO-3166 country codes, HTTP headers/encodings/auth prefixes). Import path: `github.com/foomo/gostandards`.

## Commands

```bash
make generate       # Run go generate (regenerate code from CSV sources)
make test           # Run tests
make test.race      # Run tests with -race flag
make test.update    # Run tests with -update flag
make lint           # Run golangci-lint
make lint.fix       # Auto-fix lint violations
make tidy           # Run go mod tidy
make outdated       # Show outdated direct dependencies
make docs           # Start VitePress docs dev server
make godocs         # Open go docs
```

Run a single test:
```bash
go test -tags=safe -run TestName ./iso4217/...
```

Tests use build tag `safe` — always include `-tags=safe` when running tests directly.

## Prerequisites

[mise](https://mise.jdx.dev) must be installed. It manages tool versions for `lefthook` and `golangci-lint` (see `.mise.toml`).

## Architecture

Each standard lives in its own package exporting a string-based type with constants:

- **`iso4217/`** — `Code` type (string) with currency code constants (e.g., `CodeUSD`). Code is **generated** — do not edit `code_gen.go` directly.
- **`iso3166/`** — `Alpha2` and `Alpha3` types (strings) with country code constants (e.g., `Alpha2US`, `Alpha3USA`). Code is **generated** — do not edit `code_gen.go` directly.
- **`http/`** — `Header`, `Encoding`, and `AuthPrefix` types (strings) with standard HTTP constants
- **`cmd/generate/`** — Shared code generator that produces `code_gen.go` for both `iso3166` and `iso4217` from remote CSV sources using Go templates. Run `make generate` or `go generate ./...` to regenerate.

## Documentation

Uses [VitePress](https://vitepress.dev) for project documentation, deployed as a GitHub Pages site. Source files live in `docs/`. Run `make docs` to start the local dev server (`bun run dev`).

## Linting

Uses golangci-lint v2 with `default: all` (all linters enabled) minus a specific exclusion list. Build tag `safe` is required. The linter also runs `gofmt` and `goimports` as formatters.
