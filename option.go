package gutils

import (
	"errors"
	"gutils/common"
)

type Option[T any] struct {
	some T
	none bool
}

func (o *Option[T]) And(optb Option[T]) Option[T] {
	if o.IsNone() {
		return None[T]()
	}
	return optb
}

func AndO[T, U any](o Option[T], optb Option[U]) Option[U] {
	if o.IsNone() {
		return None[U]()
	}
	return optb
}

func (o *Option[T]) AndThen(f func(t T) Option[T]) Option[T] {
	if o.IsNone() {
		return None[T]()
	}
	return f(o.Some())
}

func AndThenO[T, U any](o Option[T], f func(t T) Option[U]) Option[U] {
	if o.IsNone() {
		return None[U]()
	}
	return f(o.Some())
}

func (o *Option[T]) Or(optb Option[T]) Option[T] {
	if o.IsNone() {
		return optb
	}
	return Some(o.Some())
}

func (o *Option[T]) OrElse(f func() Option[T]) Option[T] {
	if o.IsNone() {
		return f()
	}
	return Some(o.Some())
}

func (o *Option[T]) Xor(optb Option[T]) Option[T] {
	if o.IsNone() && optb.IsSome() {
		return Some(optb.Some())
	}
	if o.IsSome() && optb.IsNone() {
		return Some(o.Some())
	}
	return None[T]()
}

func (o *Option[T]) Map(f func(t T) T) Option[T] {
	if o.IsNone() {
		return None[T]()
	}
	return Some(f(o.Some()))
}

func MapO[T, U any](o Option[T], f func(t T) U) Option[U] {
	if o.IsNone() {
		return None[U]()
	}
	return Some(f(o.Some()))
}

func (o *Option[T]) MapOr(dv T, f func(t T) T) T {
	if o.IsNone() {
		return dv
	}
	return f(o.Some())
}

func MapOrO[T, U any](o Option[T], dv U, f func(t T) U) U {
	if o.IsNone() {
		return dv
	}
	return f(o.Some())
}

func (o *Option[T]) MapOrElse(df func() T, f func(t T) T) T {
	if o.IsNone() {
		return df()
	}
	return f(o.Some())
}

func MapOrElseO[T, U any](o Option[T], df func() U, f func(t T) U) U {
	if o.IsNone() {
		return df()
	}
	return f(o.Some())
}

func (o *Option[T]) OkOr(err error) Result[T] {
	if o.IsNone() {
		return Err[T](err)
	}
	return Ok[T](o.Some())
}

func (o *Option[T]) OkOrElse(err func() error) Result[T] {
	if o.IsNone() {
		return Err[T](err())
	}
	return Ok[T](o.Some())
}

func (o *Option[T]) IsSome() bool {
	return !o.IsNone()
}

func (o *Option[T]) IsSomeAnd(f func(t T) bool) bool {
	if o.IsSome() {
		return f(o.Some())
	}
	return false
}

func (o *Option[T]) IsNone() bool {
	return o.none
}

func (o *Option[T]) Some() T {
	return o.some
}

func (o *Option[T]) Get() (T, bool) {
	if o.IsNone() {
		return common.Zero[T](), false
	}
	return o.Some(), true
}

func (o *Option[T]) GetOrElse(dv T) T {
	if o.IsNone() {
		return dv
	}
	return o.Some()
}

func (o *Option[T]) GetOrInsert(dv T) T {
	if o.IsNone() {
		o.some = dv
		o.none = false
		return dv
	}
	return o.Some()
}

func (o *Option[T]) GetOrInsertWith(f func() T) T {
	if o.IsNone() {
		o.some = f()
		o.none = false
		return o.Some()
	}
	return o.Some()
}

func (o *Option[T]) Unwrap() T {
	if o.IsNone() {
		panic(errors.New("called `Option::Unwrap()` on a `None` value"))
	}
	return o.Some()
}

func (o *Option[T]) UnwrapOr(dv T) T {
	if o.IsNone() {
		return dv
	}
	return o.Some()
}

func (o *Option[T]) UnwrapOrElse(f func() T) T {
	if o.IsNone() {
		return f()
	}
	return o.Some()
}

func (o *Option[T]) UnwrapOrDefault() T {
	if o.IsNone() {
		return common.Zero[T]()
	}
	return o.Some()
}

func Some[T any](t T) (o Option[T]) {
	o.some = t
	return
}

func None[T any]() (o Option[T]) {
	o.none = true
	return
}
