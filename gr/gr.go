package gr

type Result[T any] struct {
	ok  T
	err error
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

func Ok[T any](t T) (r Result[T]) {
	r.ok = t
	return
}

func Err[T any](err error) (r Result[T]) {
	r.err = err
	return
}
