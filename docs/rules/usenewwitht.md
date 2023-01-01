# usenewwitht

## Why

The `usenewwitht` rule is to help codebases avoid using the deprecated `NewGomegaWithT` and instead use `NewWithT`.

## Fixable

The `usenewwitht` rule can automatically replace usages of `NewGomegaWithT` with `NewWithT`.

## Invalid

The following examples are all cases flagged by the `usenewwitht` rule.

```go
	gomega.NewGomegaWithT(t)
	NewGomegaWithT(t)
```

## Valid

The following examples are considered valid by the `usenewwitht` rule.

```go
	gomega.NewWithT(t)
	NewWithT(t)
```

