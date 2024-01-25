package gutils

import (
	"errors"
	"reflect"
	"testing"
)

func TestAnd(t *testing.T) {
	type args[T any, U any] struct {
		o    Option[T]
		optb Option[U]
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want Option[U]
	}
	tests := []testCase[int, string]{
		{
			name: "AndTest1",
			args: args[int, string]{o: Some(2), optb: None[string]()},
			want: None[string](),
		},
		{
			name: "AndTest2",
			args: args[int, string]{None[int](), Some[string]("foo")},
			want: None[string](),
		},
		{
			name: "AndTest3",
			args: args[int, string]{Some(2), Some[string]("foo")},
			want: Some("foo"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AndO(tt.args.o, tt.args.optb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AndO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAndThenO(t *testing.T) {
	type args[T any, U any] struct {
		o    Option[T]
		f    func(t T) Option[U]
		want Option[U]
	}
	tests := []struct {
		name string
		args args[int, string]
	}{
		{
			name: "AndThenTest1",
			args: args[int, string]{o: Some(2), f: func(t int) Option[string] {
				return Some("foo")
			}, want: Some("foo")},
		},
		{
			name: "AndThenTest2",
			args: args[int, string]{o: None[int](), f: func(t int) Option[string] {
				return Some("foo")
			}, want: None[string]()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AndThenO(tt.args.o, tt.args.f); !reflect.DeepEqual(got, tt.args.want) {
				t.Errorf("AndThenO() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[T any, U any] struct {
		o Option[T]
		f func(t T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want Option[U]
	}
	tests := []testCase[int, string]{
		{
			name: "MapTest1",
			args: args[int, string]{o: Some(2), f: func(t int) string {
				return "foo"
			}},
			want: Some("foo"),
		},
		{
			name: "MapTest2",
			args: args[int, string]{o: None[int](), f: func(t int) string {
				return "foo"
			}},
			want: None[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapO(tt.args.o, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapOr(t *testing.T) {
	type args[T any, U any] struct {
		o  Option[T]
		dv U
		f  func(t T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want U
	}
	tests := []testCase[int, string]{
		{
			name: "MapOrTest1",
			args: args[int, string]{
				Some(2), "ok", func(t int) string {
					return "foo"
				},
			},
			want: "foo",
		},
		{
			name: "MapOrTest2",
			args: args[int, string]{
				None[int](), "ok", func(t int) string {
					return "foo"
				},
			},
			want: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapOrO(tt.args.o, tt.args.dv, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapOrO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapOrElse(t *testing.T) {
	df := func() string {
		return "ok"
	}
	f := func(t int) string {
		return "foo"
	}
	type args[T any, U any] struct {
		o  Option[T]
		df func() U
		f  func(t T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want U
	}
	tests := []testCase[int, string]{
		{
			name: "MapOrElseTest1",
			args: args[int, string]{
				Some(2), df, f,
			},
			want: "foo",
		},
		{
			name: "MapOrElseTest2",
			args: args[int, string]{None[int](), df, f},
			want: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapOrElseO(tt.args.o, tt.args.df, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapOrElseO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNil(t *testing.T) {
	type testCase[T any] struct {
		name  string
		wantO Option[T]
	}
	tests := []testCase[int]{
		{
			name:  "NilTest1",
			wantO: Option[int]{none: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotO := None[int](); !reflect.DeepEqual(gotO, tt.wantO) {
				t.Errorf("None() = %v, want %v", gotO, tt.wantO)
			}
		})
	}
}

func TestOption_And(t *testing.T) {
	type args[T any] struct {
		optb Option[T]
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want Option[T]
	}
	tests := []testCase[int]{
		{
			name: "Option_AndTest1",
			o:    Some(2),
			args: args[int]{
				Some(3),
			},
			want: Some(3),
		},
		{
			name: "Option_AndTest2",
			o:    None[int](),
			args: args[int]{
				Some(3),
			},
			want: None[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.And(tt.args.optb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AndO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_AndThen(t *testing.T) {
	type args[T any] struct {
		f func(t T) Option[T]
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want Option[T]
	}
	tests := []testCase[int]{
		{
			name: "Option_AndThenTest1",
			o:    Some(2),
			args: args[int]{func(t int) Option[int] {
				return Some(3)
			}},
			want: Some(3),
		},
		{
			name: "Option_AndThenTest2",
			o:    None[int](),
			args: args[int]{func(t int) Option[int] {
				return Some(3)
			}},
			want: None[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.AndThen(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AndThen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_Get(t *testing.T) {
	type testCase[T any] struct {
		name  string
		o     Option[T]
		want  T
		want1 bool
	}
	tests := []testCase[int]{
		{
			name:  "Option_GetTest1",
			o:     Some(2),
			want:  2,
			want1: true,
		},
		{
			name:  "Option_GetTest2",
			o:     None[int](),
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.o.Get()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOption_GetOrElse(t *testing.T) {
	type args[T any] struct {
		dv T
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Option_GetOrElseTest1",
			o:    Some(2),
			args: args[int]{3},
			want: 2,
		},
		{
			name: "Option_GetOrElseTest2",
			o:    None[int](),
			args: args[int]{3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.GetOrElse(tt.args.dv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_GetOrInsert(t *testing.T) {
	type args[T any] struct {
		dv T
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Option_GetOrInsertTest1",
			o:    Some(2),
			args: args[int]{3},
			want: 2,
		},
		{
			name: "Option_GetOrInsertTest2",
			o:    None[int](),
			args: args[int]{3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.GetOrInsert(tt.args.dv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_GetOrInsertWith(t *testing.T) {
	type args[T any] struct {
		f func() T
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Option_GetOrInsertWithTest1",
			o:    Some(2),
			args: args[int]{func() int {
				return 3
			}},
			want: 2,
		},
		{
			name: "Option_GetOrInsertWithTest2",
			o:    None[int](),
			args: args[int]{func() int {
				return 3
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.GetOrInsertWith(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrInsertWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_IsNone(t *testing.T) {
	type testCase[T any] struct {
		name string
		o    Option[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Option_IsNilTest1",
			o:    Some(2),
			want: false,
		},
		{
			name: "Option_IsNilTest2",
			o:    None[int](),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.IsNone(); got != tt.want {
				t.Errorf("IsNone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_IsSome(t *testing.T) {
	type testCase[T any] struct {
		name string
		o    Option[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Option_IsSomeTest1",
			o:    Some(2),
			want: true,
		},
		{
			name: "Option_IsSomeTest2",
			o:    None[int](),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.IsSome(); got != tt.want {
				t.Errorf("IsSome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_IsSomeAnd(t *testing.T) {
	type args[T any] struct {
		f func(t T) bool
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Option_IsSomeAndTest1",
			o:    Some(2),
			args: args[int]{func(t int) bool {
				return t == 2
			}},
			want: true,
		},
		{
			name: "Option_IsSomeAndTest2",
			o:    None[int](),
			args: args[int]{func(t int) bool {
				return t == 2
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.IsSomeAnd(tt.args.f); got != tt.want {
				t.Errorf("IsSomeAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_Map(t *testing.T) {
	type args[T any] struct {
		f func(t T) T
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want Option[T]
	}
	tests := []testCase[int]{
		{
			name: "Option_MapTest1",
			o:    Some(2),
			args: args[int]{func(t int) int {
				return 2 * t
			}},
			want: Some(4),
		},
		{
			name: "Option_MapTest2",
			o:    None[int](),
			args: args[int]{func(t int) int {
				return 2 * t
			}},
			want: None[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Map(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_MapOr(t *testing.T) {
	type args[T any] struct {
		dv T
		f  func(t T) T
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Option_MapOrTest1",
			o:    Some(2),
			args: args[int]{
				dv: 3,
				f: func(t int) int {
					return 2 * t
				},
			},
			want: 4,
		},
		{
			name: "Option_MapOrTest2",
			o:    None[int](),
			args: args[int]{
				dv: 3,
				f: func(t int) int {
					return 2 * t
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.MapOr(tt.args.dv, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_MapOrElse(t *testing.T) {
	type args[T any] struct {
		df func() T
		f  func(t T) T
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Option_MapOrElseTest1",
			o:    Some(2),
			args: args[int]{
				df: func() int {
					return 3
				},
				f: func(t int) int {
					return 2 * t
				},
			},
			want: 4,
		},
		{
			name: "Option_MapOrElseTest2",
			o:    None[int](),
			args: args[int]{
				df: func() int {
					return 3
				},
				f: func(t int) int {
					return 2 * t
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.MapOrElse(tt.args.df, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapOrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_Or(t *testing.T) {
	type args[T any] struct {
		optb Option[T]
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want Option[T]
	}
	tests := []testCase[int]{
		{
			name: "Option_OrTest1",
			o:    Some(2),
			args: args[int]{
				optb: Some(3),
			},
			want: Some(2),
		},
		{
			name: "Option_OrTest2",
			o:    None[int](),
			args: args[int]{
				optb: Some(3),
			},
			want: Some(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Or(tt.args.optb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_OrElse(t *testing.T) {
	type args[T any] struct {
		f func() Option[T]
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want Option[T]
	}
	tests := []testCase[int]{
		{
			name: "Option_OrElseTest1",
			o:    Some(2),
			args: args[int]{
				f: func() Option[int] {
					return Some(3)
				},
			},
			want: Some(2),
		},
		{
			name: "Option_OrElseTest2",
			o:    None[int](),
			args: args[int]{
				f: func() Option[int] {
					return Some(3)
				},
			},
			want: Some(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.OrElse(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_Some(t *testing.T) {
	type testCase[T any] struct {
		name string
		o    Option[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Option_SomeTest1",
			o:    Some(2),
			want: 2,
		},
		{
			name: "Option_SomeTest2",
			o:    None[int](),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Some(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Some() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_Unwrap(t *testing.T) {
	type testCase[T any] struct {
		name string
		o    Option[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Option_UnwrapTest1",
			o:    Some(2),
			want: 2,
		},
		// {
		// 	name: "Option_UnwrapTest2",
		// 	o:    None[int](),
		// 	want: 0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Unwrap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_UnwrapOr(t *testing.T) {
	type args[T any] struct {
		dv T
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Option_UnwrapOrTest1",
			o:    Some(2),
			args: args[int]{3},
			want: 2,
		},
		{
			name: "Option_UnwrapOrTest2",
			o:    None[int](),
			args: args[int]{3},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.UnwrapOr(tt.args.dv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_UnwrapOrElse(t *testing.T) {
	type args[T any] struct {
		f func() T
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Option_UnwrapOrElseTest1",
			o:    Some(2),
			args: args[int]{func() int {
				return 3
			}},
			want: 2,
		},
		{
			name: "Option_UnwrapOrElseTest2",
			o:    None[int](),
			args: args[int]{func() int {
				return 3
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.UnwrapOrElse(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapOrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_UnwrapOrDefault(t *testing.T) {
	type testCase[T any] struct {
		name string
		o    Option[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Option_UnwrapOrDefaultTest1",
			o:    Some(2),
			want: 2,
		},
		{
			name: "Option_UnwrapOrDefaultTest2",
			o:    None[int](),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.UnwrapOrDefault(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_Xor(t *testing.T) {
	type args[T any] struct {
		optb Option[T]
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want Option[T]
	}
	tests := []testCase[int]{
		{
			name: "Option_XorTest1",
			o:    Some(2),
			args: args[int]{Some(3)},
			want: None[int](),
		},
		{
			name: "Option_XorTest2",
			o:    None[int](),
			args: args[int]{Some(3)},
			want: Some(3),
		},
		{
			name: "Option_XorTest3",
			o:    Some(2),
			args: args[int]{None[int]()},
			want: Some(2),
		},
		{
			name: "Option_XorTest4",
			o:    None[int](),
			args: args[int]{None[int]()},
			want: None[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Xor(tt.args.optb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Xor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSome(t *testing.T) {
	type args[T any] struct {
		t T
	}
	type testCase[T any] struct {
		name  string
		args  args[T]
		wantO Option[T]
	}
	tests := []testCase[int]{
		{
			name:  "SomeTest1",
			args:  args[int]{t: 2},
			wantO: Option[int]{some: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotO := Some(tt.args.t); !reflect.DeepEqual(gotO, tt.wantO) {
				t.Errorf("Some() = %v, want %v", gotO, tt.wantO)
			}
		})
	}
}

func TestOption_OkOr(t *testing.T) {
	type args[T any] struct {
		err error
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want Result[T]
	}
	tests := []testCase[int]{
		{
			name: "Option_OkOrTest1",
			o:    Some(2),
			args: args[int]{errors.New("foo")},
			want: Ok[int](2),
		},
		{
			name: "Option_OkOrTest2",
			o:    None[int](),
			args: args[int]{errors.New("foo")},
			want: Err[int](errors.New("foo")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.OkOr(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OkOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOption_OkOrElse(t *testing.T) {
	type args[T any] struct {
		err func() error
	}
	type testCase[T any] struct {
		name string
		o    Option[T]
		args args[T]
		want Result[T]
	}
	tests := []testCase[int]{
		{
			name: "Option_OkOrElseTest1",
			o:    Some(2),
			args: args[int]{func() error {
				return errors.New("foo")
			}},
			want: Ok[int](2),
		},
		{
			name: "Option_OkOrElseTest2",
			o:    None[int](),
			args: args[int]{func() error {
				return errors.New("foo")
			}},
			want: Err[int](errors.New("foo")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.OkOrElse(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OkOrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}
