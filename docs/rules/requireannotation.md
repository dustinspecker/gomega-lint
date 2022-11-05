# requireannotation

## Why

The `requireannotation` rule is to help codebases provide more context for assertions than just the failing case. `expected nil to not be nil` is not helpful, while `user should get an error when deleting another user's post: expected nil to not be nil` can be helpful to read in output..

## Invalid

The following examples are all cases flagged by the `requireannotation` rule.

```go
	Expect("hello").To(Equal("hello"))
	Ω("hello").To(Equal("hello"))
	Eventually("hello").Should(Equal("hello"))
	Consistently("hello").Should(Equal("hello"))
```

## Valid

The following examples are considered valid by the `requireannotation` rule.

```go
	Expect("hello").To(Equal("hello"), "helpful annotation")
	Ω("hello").To(Equal("hello"), "helpful annotation")
	Eventually("hello").Should(Equal("hello"), "helpful annotation")
	Consistently("hello").Should(Equal("hello"), "helpful annotation")
```
