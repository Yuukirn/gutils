package gm

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
		f func(k K) bool
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want map[K]V
	}
	tests := []testCase[string, int]{
		{
			name: "FilterTest1",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}, f: func(k string) bool { return k == "a" }},
			want: map[string]int{"a": 1},
		},
		{
			name: "FilterTest2",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}, f: func(k string) bool { return k == "d" }},
			want: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOr(t *testing.T) {
	type args[K comparable, V any] struct {
		m  map[K]V
		k  K
		dv V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want V
	}
	tests := []testCase[string, int]{
		{
			name: "GetOrTest1",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}, k: "a", dv: 0},
			want: 1,
		},
		{
			name: "GetOrTest2",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}, k: "d", dv: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOr(tt.args.m, tt.args.k, tt.args.dv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOrDefault(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
		k K
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want V
	}
	tests := []testCase[string, int]{
		{
			name: "GetOrDefaultTest1",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}, k: "a"},
			want: 1,
		},
		{
			name: "GetOrDefaultTest2",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}, k: "d"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOrDefault(tt.args.m, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOrInsert(t *testing.T) {
	type args[K comparable, V any] struct {
		m  map[K]V
		k  K
		dv V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want V
	}
	tests := []testCase[string, int]{
		{
			name: "GetOrInsertTest1",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}, k: "a", dv: 0},
			want: 1,
		},
		{
			name: "GetOrInsertTest2",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}, k: "d", dv: 0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOrInsert(tt.args.m, tt.args.k, tt.args.dv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeys(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want []K
	}
	tests := []testCase[string, int]{
		{
			name: "KeysTest1",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}},
			want: []string{"a", "b", "c"},
		},
		{
			name: "KeysTest2",
			args: args[string, int]{m: map[string]int{}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Keys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[K1 comparable, V1 any, K2 comparable, V2 any] struct {
		m map[K1]V1
		f func(K1, V1) (K2, V2)
	}
	type testCase[K1 comparable, V1 any, K2 comparable, V2 any] struct {
		name string
		args args[K1, V1, K2, V2]
		want map[K2]V2
	}
	tests := []testCase[string, int, int, string]{
		{
			name: "MapTest1",
			args: args[string, int, int, string]{m: map[string]int{"a": 1, "b": 2, "c": 3}, f: func(k string, v int) (int, string) { return v, k }},
			want: map[int]string{1: "a", 2: "b", 3: "c"},
		},
		{
			name: "MapTest2",
			args: args[string, int, int, string]{m: map[string]int{}, f: func(k string, v int) (int, string) { return v, k }},
			want: map[int]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args[K comparable, V any] struct {
		m1 map[K]V
		m2 map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want map[K]V
	}
	tests := []testCase[string, int]{
		{
			name: "MergeTest1",
			args: args[string, int]{m1: map[string]int{"a": 1, "b": 2, "c": 3}, m2: map[string]int{"a": 4, "d": 5}},
			want: map[string]int{"a": 4, "b": 2, "c": 3, "d": 5},
		},
		{
			name: "MergeTest2",
			args: args[string, int]{m1: map[string]int{}, m2: map[string]int{"a": 4, "d": 5}},
			want: map[string]int{"a": 4, "d": 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.m1, tt.args.m2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want []V
	}
	tests := []testCase[string, int]{
		{
			name: "ValuesTest1",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "ValuesTest2",
			args: args[string, int]{m: map[string]int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Values(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsKey(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
		k K
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want bool
	}
	tests := []testCase[string, int]{
		{
			name: "ContainsKeyTest1",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}, k: "a"},
			want: true,
		},
		{
			name: "ContainsKeyTest2",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}, k: "d"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsKey(tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("ContainsKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
