package gs

import (
	"golang.org/x/exp/constraints"
	"reflect"
	"strconv"
	"testing"
)

func TestContains(t *testing.T) {
	type args[T comparable] struct {
		s []T
		t T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "ContainsTest1",
			args: args[int]{[]int{1, 2, 3}, 2},
			want: true,
		},
		{
			name: "ContainsTest2",
			args: args[int]{[]int{1, 2, 3}, 4},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args[T any] struct {
		s []T
		f func(t T) bool
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "FilterTest1",
			args: args[int]{[]int{1, 2, 3}, func(t int) bool { return t%2 == 0 }},
			want: []int{2},
		},
		{
			name: "FilterTest2",
			args: args[int]{[]int{1, 2, 3}, func(t int) bool { return t%2 == 1 }},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[T any, U any] struct {
		s []T
		f func(t T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want []U
	}
	tests := []testCase[int, string]{
		{
			name: "MapTest1",
			args: args[int, string]{[]int{1, 2, 3}, func(t int) string { return strconv.Itoa(t) }},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args[T constraints.Ordered] struct {
		s []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "MaxTest1",
			args: args[int]{[]int{1, 2, 3}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args[T constraints.Ordered] struct {
		s []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "MinTest1",
			args: args[int]{[]int{2, 1, 3}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}
