package stream

import (
	"gutils"
	"gutils/gs"
)

type SliceStream[T any] struct {
	slice []T
}

func NewSliceStream[T any](s []T) *SliceStream[T] {
	var res = make([]T, len(s))
	copy(res, s)
	return &SliceStream[T]{res}
}

func (ss *SliceStream[T]) Filter(f func(T) bool) *SliceStream[T] {
	ss.slice = gs.Filter(ss.slice, f)
	return ss
}

func (ss *SliceStream[T]) Map(f func(T) T) *SliceStream[T] {
	ss.slice = gs.Map(ss.slice, f)
	return ss
}

func (ss *SliceStream[T]) Reverse() *SliceStream[T] {
	ss.slice = gs.Reverse(ss.slice)
	return ss
}

func (ss *SliceStream[T]) Append(s []T) *SliceStream[T] {
	ss.slice = append(ss.slice, s...)
	return ss
}

func (ss *SliceStream[T]) Prepend(s []T) *SliceStream[T] {
	ss.slice = append(s, ss.slice...)
	return ss
}

func (ss *SliceStream[T]) First() gutils.Option[T] {
	return ss.Get(0)
}

func (ss *SliceStream[T]) Last() gutils.Option[T] {
	return ss.Get(len(ss.slice) - 1)
}

func (ss *SliceStream[T]) Len() int {
	return len(ss.slice)
}

func (ss *SliceStream[T]) IsEmpty() bool {
	return len(ss.slice) == 0
}

func (ss *SliceStream[T]) Get(i int) gutils.Option[T] {
	if i < 0 || i >= len(ss.slice) {
		return gutils.None[T]()
	}
	return gutils.Some(ss.slice[i])
}

func (ss *SliceStream[T]) ToSlice() []T {
	return ss.slice
}

func (ss *SliceStream[T]) Fold(f func(T, T) T) T {
	return gs.Fold(ss.slice, f)
}

func (ss *SliceStream[T]) FoldWith(dv T, f func(T, T) T) T {
	return gs.FoldWith(ss.slice, dv, f)
}
