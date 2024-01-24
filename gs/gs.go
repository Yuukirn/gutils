package gs

import (
	"golang.org/x/exp/constraints"
)

func Filter[T any](s []T, f func(t T) bool) []T {
	var res []T
	for _, ss := range s {
		if f(ss) {
			res = append(res, ss)
		}
	}
	return res
}

func Map[T, U any](s []T, f func(t T) U) []U {
	var res = make([]U, 0, len(s))
	for i := range s {
		res = append(res, f(s[i]))
	}
	return res
}

func Contains[T comparable](s []T, t T) bool {
	for i := range s {
		if s[i] == t {
			return true
		}
	}
	return false
}

func Max[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	var m = s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > m {
			m = s[i]
		}
	}
	return m
}

func Min[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		var zero T
		return zero
	}
	var m = s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < m {
			m = s[i]
		}
	}
	return m
}
