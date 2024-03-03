package stream

import (
	"github.com/Yuukirn/gutils"
	"github.com/Yuukirn/gutils/gm"
	"maps"
)

type MapStream[K comparable, V any] struct {
	m map[K]V
}

func NewMapStream[K comparable, V any](m map[K]V) *MapStream[K, V] {
	var res = make(map[K]V, len(m))
	maps.Copy(res, m)
	return &MapStream[K, V]{res}
}

func (ms *MapStream[K, V]) Map(f func(K, V) (K, V)) *MapStream[K, V] {
	ms.m = gm.Map(ms.m, f)
	return ms
}

func (ms *MapStream[K, V]) Filter(f func(K) bool) *MapStream[K, V] {
	ms.m = gm.Filter(ms.m, f)
	return ms
}

func (ms *MapStream[K, V]) Merge(m map[K]V) *MapStream[K, V] {
	ms.m = gm.Merge(ms.m, m)
	return ms
}

func (ms *MapStream[K, V]) Keys() []K {
	return gm.Keys(ms.m)
}

func (ms *MapStream[K, V]) Values() []V {
	return gm.Values(ms.m)
}

func (ms *MapStream[K, V]) Get(k K) gutils.Option[V] {
	return gm.Get(ms.m, k)
}

func (ms *MapStream[K, V]) GetOr(k K, dv V) V {
	return gm.GetOr(ms.m, k, dv)
}

func (ms *MapStream[K, V]) GetOrInsert(k K, dv V) V {
	return gm.GetOrInsert(ms.m, k, dv)
}

func (ms *MapStream[K, V]) GetOrDefault(k K) V {
	return gm.GetOrDefault(ms.m, k)
}

func (ms *MapStream[K, V]) Len() int {
	return len(ms.m)
}

func (ms *MapStream[K, V]) IsEmpty() bool {
	return len(ms.m) == 0
}

func (ms *MapStream[K, V]) ContainsKey(k K) bool {
	return gm.ContainsKey(ms.m, k)
}

func (ms *MapStream[K, V]) ToMap() map[K]V {
	return ms.m
}
