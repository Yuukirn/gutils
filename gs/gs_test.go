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

func TestReverse(t *testing.T) {
	type args[T any] struct {
		s []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "ReverseTest1",
			args: args[int]{[]int{1, 2, 3}},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args[T constraints.Integer | constraints.Float] struct {
		s []T
	}
	type testCase[T constraints.Integer | constraints.Float] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "SumTest1",
			args: args[int]{[]int{1, 2, 3}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumWith(t *testing.T) {
	type args[T constraints.Integer | constraints.Float] struct {
		s  []T
		dv T
	}
	type testCase[T constraints.Integer | constraints.Float] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "SumWithTest1",
			args: args[int]{[]int{1, 2, 3}, 10},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumWith(tt.args.s, tt.args.dv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFold(t *testing.T) {
	type args[T any] struct {
		s []T
		f func(T, T) T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "FoldTest1",
			args: args[int]{[]int{1, 2, 3}, func(a, b int) int { return a + b }},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fold(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fold() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestFoldWith(t *testing.T) {
	type args[T any] struct {
		s  []T
		dv T
		f  func(T, T) T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "FoldWithTest1",
			args: args[int]{[]int{1, 2, 3}, 10, func(a, b int) int { return a + b }},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FoldWith(tt.args.s, tt.args.dv, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoldWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
