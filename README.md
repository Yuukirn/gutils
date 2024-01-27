# gutils
Go generic utilities for personal use.

Subpackages:
- gm - Generic operations for maps
- gs - Generic operations for slices
- stream - Stream Processing for map and slice

This package also supports Option and Result types, which are inspired by Rust, 
and provides the same API as Rust.

However, Go itself does not support generic methods, so the API is not as elegant as Rust.
For example, the definition of `Map` method of `Option` type is as follows:

```go
func (o Option[T]) Map(f func(t T) T) Option[T]
```
It can't transform the type of `Option` itself, so it can only return `Option[T]` instead of `Option[U]`.

So if you want to transform the type of `Option`, you need to use MapO function:
```go
func MapO[T, U any](o Option[T], f func(t T) U) Option[U]
```

Besides, this package has some useful utilities, such as `Ref` to get the pointer of a rvalue, 
and `Deref` to get the value of a pointer safely.