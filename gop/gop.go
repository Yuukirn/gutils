package gop

import "errors"

// TODO: add ok_or...

type Option[T any] struct {
	some T
	none bool
}

func (o *Option[T]) And(optb Option[T]) Option[T] {
	if o.IsNil() {
		return Nil[T]()
	}
	return optb
}

func And[T, U any](o Option[T], optb Option[U]) Option[U] {
	if o.IsNil() {
		return Nil[U]()
	}
	return optb
}

func (o *Option[T]) AndThen(f func(t T) T) Option[T] {
	if o.IsNil() {
		return Nil[T]()
	}
	return Some(f(o.Some()))
}

func AndThen[T, U any](o Option[T], f func(t T) U) Option[U] {
	if o.IsNil() {
		return Nil[U]()
	}
	return Some(f(o.Some()))
}

func (o *Option[T]) Or(optb Option[T]) Option[T] {
	if o.IsNil() {
		return optb
	}
	return Some(o.Some())
}

func (o *Option[T]) OrElse(f func() T) Option[T] {
	if o.IsNil() {
		return Some(f())
	}
	return Some(o.Some())
}

func (o *Option[T]) Xor(optb Option[T]) Option[T] {
	if o.IsNil() && optb.IsSome() {
		return Some(optb.Some())
	}
	if o.IsSome() && optb.IsNil() {
		return Some(o.Some())
	}
	return Nil[T]()
}

func (o *Option[T]) Map(f func(t T) T) Option[T] {
	if o.IsNil() {
		return Nil[T]()
	}
	return Some(f(o.Some()))
}

func Map[T, U any](o Option[T], f func(t T) U) Option[U] {
	if o.IsNil() {
		return Nil[U]()
	}
	return Some(f(o.Some()))
}

func (o *Option[T]) MapOr(dv T, f func(t T) T) T {
	if o.IsNil() {
		return dv
	}
	return f(o.Some())
}

func MapOr[T, U any](o Option[T], dv U, f func(t T) U) U {
	if o.IsNil() {
		return dv
	}
	return f(o.Some())
}

func (o *Option[T]) MapOrElse(df func() T, f func(t T) T) T {
	if o.IsNil() {
		return df()
	}
	return f(o.Some())
}

func MapOrElse[T, U any](o Option[T], df func() U, f func(t T) U) U {
	if o.IsNil() {
		return df()
	}
	return f(o.Some())
}

func (o *Option[T]) IsSome() bool {
	return !o.IsNil()
}

func (o *Option[T]) IsSomeAnd(f func(t T) bool) bool {
	if o.IsSome() {
		return f(o.Some())
	}
	return false
}

func (o *Option[T]) IsNil() bool {
	return o.none == true
}

func (o *Option[T]) Some() T {
	return o.some
}

func (o *Option[T]) Get() (T, bool) {
	var zero T
	if o.IsNil() {
		return zero, false
	}
	return o.Some(), true
}

func (o *Option[T]) GetOrElse(dv T) T {
	if o.IsNil() {
		return dv
	}
	return o.Some()
}

func (o *Option[T]) GetOrInsert(dv T) T {
	if o.IsNil() {
		o.some = dv
		o.none = false
		return dv
	}
	return o.Some()
}

func (o *Option[T]) GetOrInsertWith(f func() T) T {
	if o.IsNil() {
		o.some = f()
		o.none = false
		return o.Some()
	}
	return o.Some()
}

func (o *Option[T]) Unwrap() T {
	if o.IsNil() {
		panic(errors.New("none option"))

	}
	return o.Some()
}

func (o *Option[T]) UnwrapOr(dv T) T {
	if o.IsNil() {
		return dv
	}
	return o.Some()
}

func (o *Option[T]) UnwrapOrElse(f func() T) T {
	if o.IsNil() {
		return f()
	}
	return o.Some()
}

func (o *Option[T]) UnwrapOrDefault() T {
	if o.IsNil() {
		var zero T
		return zero
	}
	return o.Some()
}

func Some[T any](t T) (o Option[T]) {
	o.some = t
	return
}

func Nil[T any]() (o Option[T]) {
	o.none = true
	return
}
