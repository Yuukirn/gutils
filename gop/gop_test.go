package gop

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
			args: args[int, string]{o: Some(2), optb: Nil[string]()},
			want: Nil[string](),
		},
		{
			name: "AndTest2",
			args: args[int, string]{Nil[int](), Some[string]("foo")},
			want: Nil[string](),
		},
		{
			name: "AndTest3",
			args: args[int, string]{Some(2), Some[string]("foo")},
			want: Some("foo"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := And(tt.args.o, tt.args.optb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("And() = %v, want %v", got, tt.want)
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
			args: args[int, string]{o: Nil[int](), f: func(t int) string {
				return "foo"
			}},
			want: Nil[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AndThen(tt.args.o, tt.args.f); !reflect.DeepEqual(got, tt.want) {
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
			args: args[int, string]{o: Nil[int](), f: func(t int) string {
				return "foo"
			}},
			want: Nil[string](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.o, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
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
				Nil[int](), "ok", func(t int) string {
					return "foo"
				},
			},
			want: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapOr(tt.args.o, tt.args.dv, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapOr() = %v, want %v", got, tt.want)
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
			args: args[int, string]{Nil[int](), df, f},
			want: "ok",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapOrElse(tt.args.o, tt.args.df, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapOrElse() = %v, want %v", got, tt.want)
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
			if gotO := Nil[int](); !reflect.DeepEqual(gotO, tt.wantO) {
				t.Errorf("Nil() = %v, want %v", gotO, tt.wantO)
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
			o:    Nil[int](),
			args: args[int]{
				Some(3),
			},
			want: Nil[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.And(tt.args.optb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("And() = %v, want %v", got, tt.want)
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
			o:    Nil[int](),
			args: args[int]{f},
			want: Nil[int](),
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
			o:     Nil[int](),
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
//			if got := tt.o.IsNil(); got != tt.want {
//				t.Errorf("IsNil() = %v, want %v", got, tt.want)
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
//			if got := tt.o.Map(tt.args.f); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Map() = %v, want %v", got, tt.want)
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
//			if got := tt.o.MapOr(tt.args.dv, tt.args.f); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("MapOr() = %v, want %v", got, tt.want)
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
//			if got := tt.o.MapOrElse(tt.args.df, tt.args.f); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("MapOrElse() = %v, want %v", got, tt.want)
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
