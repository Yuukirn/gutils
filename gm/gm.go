package gm

import "gutils"

func Keys[K comparable, V any](m map[K]V) []K {
	var keys = make([]K, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	var values = make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func Filter[K comparable, V any](m map[K]V, f func(k K) bool) map[K]V {
	var res = make(map[K]V)
	for k, v := range m {
		if f(k) {
			res[k] = v
		}
	}
	return res
}

func ContainsKey[K comparable, V any](m map[K]V, k K) bool {
	_, exist := m[k]
	return exist
}

func Get[K comparable, V any](m map[K]V, k K) gutils.Option[V] {
	if m == nil || len(m) == 0 {
		return gutils.None[V]()
	}

	v, exist := m[k]
	if !exist {
		return gutils.None[V]()
	}
	return gutils.Some(v)
}

func GetOr[K comparable, V any](m map[K]V, k K, dv V) V {
	return Get(m, k).UnwrapOr(dv)
}

func GetOrInsert[K comparable, V any](m map[K]V, k K, dv V) V {
	o := Get(m, k)
	if o.IsNone() {
		m[k] = dv
		return dv
	}
	return o.Some()
}

func GetOrDefault[K comparable, V any](m map[K]V, k K) V {
	return Get(m, k).UnwrapOrDefault()
}

func Map[K1, K2 comparable, V1, V2 any](m map[K1]V1, f func(K1, V1) (K2, V2)) map[K2]V2 {
	res := make(map[K2]V2, len(m))
	for k1, v1 := range m {
		k2, v2 := f(k1, v1)
		res[k2] = v2
	}
	return res
}

func Merge[K comparable, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}
