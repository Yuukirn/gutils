package gutils

func Zero[T any]() (t T) {
	return
}

func IsZero[T comparable](t T) bool {
	return t == Zero[T]()
}

func Ref[T any](t T) *T {
	return &t
}

func Deref[T any](t *T) T {
	if t == nil {
		return Zero[T]()
	}
	return *t
}
