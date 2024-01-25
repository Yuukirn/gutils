package gutils

import "gutils/common"

type Result[T any] struct {
	ok  T
	err error
}

func (r *Result[T]) Ok() Option[T] {
	if r.IsErr() {
		return None[T]()
	}
	return Some(r.ok)
}

func (r *Result[T]) IsOk() bool {
	return !r.IsErr()
}

func (r *Result[T]) IsOkAnd(f func(t T) bool) bool {
	if r.IsErr() {
		return false
	}
	return f(r.ok)
}

func (r *Result[T]) IsErr() bool {
	return r.err != nil
}

func (r *Result[T]) IsErrAnd(f func(err error) bool) bool {
	if r.IsErr() {
		return f(r.err)
	}
	return false
}

func (r *Result[T]) Map(f func(t T) T) Result[T] {
	if r.IsErr() {
		return Err[T](r.err)
	}
	return Ok(f(r.ok))
}

func MapR[T, U any](r Result[T], f func(t T) U) Result[U] {
	if r.IsErr() {
		return Err[U](r.err)
	}
	return Ok(f(r.ok))
}

func (r *Result[T]) MapOr(dv T, f func(t T) T) T {
	if r.IsErr() {
		return dv
	}
	return f(r.ok)
}

func MapOrR[T, U any](r Result[T], dv U, f func(t T) U) U {
	if r.IsErr() {
		return dv
	}
	return f(r.ok)
}

func (r *Result[T]) MapOrElse(d func(err error) T, f func(t T) T) T {
	if r.IsErr() {
		return d(r.err)
	}
	return f(r.ok)
}

func MapOrElseR[T, U any](r Result[T], d func(err error) U, f func(t T) U) U {
	if r.IsErr() {
		return d(r.err)
	}
	return f(r.ok)
}

func (r *Result[T]) MapErr(f func(err error) error) Result[T] {
	if r.IsErr() {
		return Err[T](f(r.err))
	}
	return *r
}

func (r *Result[T]) And(res Result[T]) Result[T] {
	if r.IsErr() {
		return *r
	}
	return res
}

func AndR[T, U any](r Result[T], res Result[U]) Result[U] {
	if r.IsErr() {
		return Err[U](r.err)
	}
	return res
}

func (r *Result[T]) AndThen(f func(t T) Result[T]) Result[T] {
	if r.IsErr() {
		return *r
	}
	return f(r.ok)
}

func AndThenR[T, U any](r Result[T], f func(t T) Result[U]) Result[U] {
	if r.IsErr() {
		return Err[U](r.err)
	}
	return f(r.ok)
}

func (r *Result[T]) Or(res Result[T]) Result[T] {
	if r.IsErr() {
		return res
	}
	return *r
}

func (r *Result[T]) OrElse(f func(err error) Result[T]) Result[T] {
	if r.IsErr() {
		return f(r.err)
	}
	return *r
}

func (r *Result[T]) Unwrap() T {
	if r.IsErr() {
		panic("called `Result::Unwrap()` on an `Err` value")
	}
	return r.ok
}

func (r *Result[T]) UnwrapOr(d T) T {
	if r.IsErr() {
		return d
	}
	return r.ok
}

func (r *Result[T]) UnwrapOrElse(f func(err error) T) T {
	if r.IsErr() {
		return f(r.err)
	}
	return r.ok
}

func (r *Result[T]) UnwrapOrDefault() T {
	if r.IsErr() {
		return common.Zero[T]()
	}
	return r.ok
}

func (r *Result[T]) Expect(msg string) T {
	if r.IsErr() {
		panic(msg)
	}
	return r.ok
}

func (r *Result[T]) ExpectErr(msg string) error {
	if r.IsErr() {
		return r.err
	}
	panic(msg)
}

func Ok[T any](t T) (r Result[T]) {
	r.ok = t
	return
}

func Err[T any](err error) (r Result[T]) {
	r.err = err
	return
}
