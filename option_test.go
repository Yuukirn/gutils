package gutils

import (
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

func TestAndThen(t *testing.T) {
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
			name: "AndThenTest1",
			args: args[int, string]{o: Some(2), f: func(t int) string {
				return "foo"
			}},
			want: Some("foo"),
		},
		{
			name: "AndThenTest2",
			args: args[int, string]{o: None[int](), f: func(t int) string {
				return "foo"
			}},
			want: None[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AndThenO(tt.args.o, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AndThen() = %v, want %v", got, tt.want)
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
	f := func(t int) int {
		return 2 * t
	}
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
			name: "Option_AndThenTest1",
			o:    Some(2),
			args: args[int]{f},
			want: Some(4),
		},
		{
			name: "Option_AndThenTest2",
			o:    None[int](),
			args: args[int]{f},
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

//
//func TestOption_GetOrElse(t *testing.T) {
//	type args[T any] struct {
//		dv T
//	}
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		args args[T]
//		want T
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.GetOrElse(tt.args.dv); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetOrElse() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_GetOrInsert(t *testing.T) {
//	type args[T any] struct {
//		dv T
//	}
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		args args[T]
//		want T
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.GetOrInsert(tt.args.dv); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetOrInsert() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_GetOrInsertWith(t *testing.T) {
//	type args[T any] struct {
//		f func() T
//	}
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		args args[T]
//		want T
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.GetOrInsertWith(tt.args.f); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetOrInsertWith() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_IsNil(t *testing.T) {
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		want bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.IsNone(); got != tt.want {
//				t.Errorf("IsNone() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_IsSome(t *testing.T) {
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		want bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.IsSome(); got != tt.want {
//				t.Errorf("IsSome() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_IsSomeAnd(t *testing.T) {
//	type args[T any] struct {
//		f func(t T) bool
//	}
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		args args[T]
//		want bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.IsSomeAnd(tt.args.f); got != tt.want {
//				t.Errorf("IsSomeAnd() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_Map(t *testing.T) {
//	type args[T any] struct {
//		f func(t T) T
//	}
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		args args[T]
//		want Option
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.MapO(tt.args.f); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("MapO() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_MapOr(t *testing.T) {
//	type args[T any] struct {
//		dv T
//		f  func(t T) T
//	}
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		args args[T]
//		want T
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.MapOrO(tt.args.dv, tt.args.f); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("MapOrO() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_MapOrElse(t *testing.T) {
//	type args[T any] struct {
//		df func() T
//		f  func(t T) T
//	}
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		args args[T]
//		want T
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.MapOrElseO(tt.args.df, tt.args.f); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("MapOrElseO() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_Or(t *testing.T) {
//	type args[T any] struct {
//		optb Option
//	}
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		args args[T]
//		want Option
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.Or(tt.args.optb); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Or() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_OrElse(t *testing.T) {
//	type args[T any] struct {
//		f func() T
//	}
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		args args[T]
//		want Option
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.OrElse(tt.args.f); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("OrElse() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_Some(t *testing.T) {
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		want T
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.Some(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Some() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_Unwrap(t *testing.T) {
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		want T
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.Unwrap(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Unwrap() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_UnwrapOr(t *testing.T) {
//	type args[T any] struct {
//		dv T
//	}
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		args args[T]
//		want T
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.UnwrapOr(tt.args.dv); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("UnwrapOr() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_UnwrapOrDefault(t *testing.T) {
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		want T
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.UnwrapOrDefault(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("UnwrapOrDefault() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_UnwrapOrElse(t *testing.T) {
//	type args[T any] struct {
//		f func() T
//	}
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		args args[T]
//		want T
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.UnwrapOrElse(tt.args.f); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("UnwrapOrElse() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOption_Xor(t *testing.T) {
//	type args[T any] struct {
//		optb Option
//	}
//	type testCase[T any] struct {
//		name string
//		o    Option[T]
//		args args[T]
//		want Option
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.o.Xor(tt.args.optb); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Xor() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestSome(t *testing.T) {
//	type args[T any] struct {
//		t T
//	}
//	type testCase[T any] struct {
//		name  string
//		args  args[T]
//		wantO Option
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if gotO := Some(tt.args.t); !reflect.DeepEqual(gotO, tt.wantO) {
//				t.Errorf("Some() = %v, want %v", gotO, tt.wantO)
//			}
//		})
//	}
//}
