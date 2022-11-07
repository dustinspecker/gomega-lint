# gomega-lint

> opinionated linter for [gomega](https://onsi.github.io/gomega/)

## Install

1. `go install github.com/dustinspecker/gomega-lint/cmd/gomega-lint`

## Usage

### Check for violations

```shell
gomega-lint ./...
```

### Fix

Some rules support automatically fixing violations.

```shell
gomega-lint -fix ./...
```

### Run individual rules

By default, `gomega-lint` runs all rules. To specify which rules to run, use the following:

```shell
gomega-lint -noformatannotation -requireannotation ./...
```

## Rules

- [noformatannotation](docs/rules/noformatannotation.md) (fixable)
- [requireannotation](docs/rules/requireannotation.md)

## License
MIT
