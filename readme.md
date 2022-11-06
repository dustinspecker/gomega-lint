# gomega-lint

> opinionated linter for [gomega](https://onsi.github.io/gomega/)

## Install

1. `go install github.com/dustinspecker/gomega-lint@latest`

## Usage

### Check for violations

1. `gomega-lint ./...`

### Fix

Some rules support automatically fixing violations.

1. `gomega-lint -fix ./...`

## Rules

- [noformatannotation](docs/rules/noformatannotation.md) (fixable)
- [requireannotation](docs/rules/requireannotation.md)

## License
MIT
