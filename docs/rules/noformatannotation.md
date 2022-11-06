# noformatannotation

## Why

The `noformatannotation` rule is to help codebases avoid useless `fmt.Sprintf` call. Gomega's assertions will pass all optional annotations through `fmt.Sprintf` for
the user, so no reason to call `fmt.Sprintf` for annotations.

## Fixable

The `noformatannotation` rule can automatically remove useless calls to `fmt.Sprintf`.

## Invalid

The following examples are all cases flagged by the `noformatannotation` rule.

```go
	Expect("hello").To(Equal("hello"), fmt.Sprintf("greeting for %s", user.Name))
	Ω("hello").To(Equal("hello"), fmt.Sprintf("greeting for %s", user.Name))
	Eventually("hello").Should(Equal("hello"), fmt.Sprintf("greeting for %s", user.Name))
	Consistently("hello").Should(Equal("hello"), fmt.Sprintf("greeting for %s", user.Name))
```

## Valid

The following examples are considered valid by the `requireannotation` rule.

```go
	Expect("hello").To(Equal("hello"), "greeting for %s", user.Name)
	Ω("hello").To(Equal("hello"), "greeting for %s", user.Name)
	Eventually("hello").Should(Equal("hello"), "greeting for %s", user.Name)
	Consistently("hello").Should(Equal("hello"), "greeting for %s", user.Name)
```

