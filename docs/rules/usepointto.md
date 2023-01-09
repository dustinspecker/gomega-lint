# usepointto

## Why

The `usepointo` rule is to help codebases avoid dereferencing and instead rely on `gstruct.PointTo`. This will help prevent accidental panics if the pointer be dereferenced is actually `nil`.

## Invalid

The following examples are all cases flagged by the `usepointto` rule.

```go
	Expect(*valuePtr).To(Equal("hello"))
	gomega.Expect(*valuePtr).To(gomega.Equal("hello"))
```

## Valid

The following examples are considered valid by the `usepointto` rule.

```go
	Expect(valuePtr).To(PointTo(Equal("hello")))
	gomega.Expect(valuePtr).To(PointTo(gomega.Equal("hello")))
```
