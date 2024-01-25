package gutils

import (
	"errors"
	"reflect"
	"testing"
)

func TestAndR(t *testing.T) {
	type args[T any, U any] struct {
		r   Result[T]
		res Result[U]
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want Result[U]
	}
	tests := []testCase[int, string]{
		{
			name: "AndRTest1",
			args: args[int, string]{Ok(3), Ok("hello")},
			want: Ok("hello"),
		},
		{
			name: "AndRTest2",
			args: args[int, string]{Err[int](errors.New("error")), Ok("hello")},
			want: Err[string](errors.New("error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AndR(tt.args.r, tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AndR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAndThenR(t *testing.T) {
	type args[T any, U any] struct {
		r Result[T]
		f func(t T) Result[U]
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want Result[U]
	}
	tests := []testCase[int, string]{
		{
			name: "AndThenRTest1",
			args: args[int, string]{Ok(3), func(t int) Result[string] { return Ok("hello") }},
			want: Ok("hello"),
		},
		{
			name: "AndThenRTest2",
			args: args[int, string]{Err[int](errors.New("error")), func(t int) Result[string] { return Ok("hello") }},
			want: Err[string](errors.New("error")),
		},
		{
			name: "AndThenRTest3",
			args: args[int, string]{Ok(3), func(t int) Result[string] { return Err[string](errors.New("error")) }},
			want: Err[string](errors.New("error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AndThenR(tt.args.r, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AndThenR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErr(t *testing.T) {
	type args struct {
		err error
	}
	type testCase[T any] struct {
		name  string
		args  args
		wantR Result[T]
	}
	tests := []testCase[int]{
		{
			name:  "ErrTest1",
			args:  args{errors.New("error")},
			wantR: Err[int](errors.New("error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := Err[int](tt.args.err); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Err() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestMapOrR(t *testing.T) {
	type args[T any, U any] struct {
		r  Result[T]
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
			name: "MapOrRTest1",
			args: args[int, string]{Ok(3), "hello", func(t int) string { return "world" }},
			want: "world",
		},
		{
			name: "MapOrRTest2",
			args: args[int, string]{Err[int](errors.New("error")), "hello", func(t int) string { return "world" }},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapOrR(tt.args.r, tt.args.dv, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapOrR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapOrElseR(t *testing.T) {
	type args[T any, U any] struct {
		r Result[T]
		d func(err error) U
		f func(t T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want U
	}
	tests := []testCase[int, string]{
		{
			name: "MapOrElseRTest1",
			args: args[int, string]{Ok(3), func(err error) string { return "hello" }, func(t int) string { return "world" }},
			want: "world",
		},
		{
			name: "MapOrElseRTest2",
			args: args[int, string]{Err[int](errors.New("error")), func(err error) string { return "hello" }, func(t int) string { return "world" }},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapOrElseR(tt.args.r, tt.args.d, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapOrElseR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapR(t *testing.T) {
	type args[T any, U any] struct {
		r Result[T]
		f func(t T) U
	}
	type testCase[T any, U any] struct {
		name string
		args args[T, U]
		want Result[U]
	}
	tests := []testCase[int, string]{
		{
			name: "MapRTest1",
			args: args[int, string]{Ok(3), func(t int) string { return "hello" }},
			want: Ok("hello"),
		},
		{
			name: "MapRTest2",
			args: args[int, string]{Err[int](errors.New("error")), func(t int) string { return "hello" }},
			want: Err[string](errors.New("error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapR(tt.args.r, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOk(t *testing.T) {
	type args[T any] struct {
		t T
	}
	type testCase[T any] struct {
		name  string
		args  args[T]
		wantR Result[T]
	}
	tests := []testCase[int]{
		{
			name:  "OkTest1",
			args:  args[int]{3},
			wantR: Ok(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := Ok(tt.args.t); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Ok() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestResult_And(t *testing.T) {
	type args[T any] struct {
		res Result[T]
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args[T]
		want Result[T]
	}
	tests := []testCase[int]{
		{
			name: "AndTest1",
			r:    Ok(3),
			args: args[int]{Ok(6)},
			want: Ok(6),
		},
		{
			name: "AndTest2",
			r:    Err[int](errors.New("error")),
			args: args[int]{Err[int](errors.New("new error"))},
			want: Err[int](errors.New("error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.And(tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_AndThen(t *testing.T) {
	type args[T any] struct {
		f func(t T) Result[T]
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args[T]
		want Result[T]
	}
	tests := []testCase[int]{
		{
			name: "AndThenTest1",
			r:    Ok(3),
			args: args[int]{func(t int) Result[int] { return Ok(6) }},
			want: Ok(6),
		},
		{
			name: "AndThenTest2",
			r:    Err[int](errors.New("error")),
			args: args[int]{func(t int) Result[int] { return Ok(6) }},
			want: Err[int](errors.New("error")),
		},
		{
			name: "AndThenTest3",
			r:    Ok(3),
			args: args[int]{func(t int) Result[int] { return Err[int](errors.New("new error")) }},
			want: Err[int](errors.New("new error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.AndThen(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AndThen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_Expect(t *testing.T) {
	type args struct {
		msg string
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args
		want T
	}
	tests := []testCase[int]{
		{
			name: "ExpectTest1",
			r:    Ok(3),
			args: args{"error"},
			want: 3,
		},
		//{
		//	name: "ExpectTest2",
		//	r:    Err[int](errors.New("error")),
		//	args: args{"error"},
		//	want: 0,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Expect(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_ExpectErr(t *testing.T) {
	type args struct {
		msg string
	}
	type testCase[T any] struct {
		name    string
		r       Result[T]
		args    args
		wantErr bool
	}
	tests := []testCase[int]{
		//{
		//	name:    "ExpectErrTest1",
		//	r:       Ok(3),
		//	args:    args{"error"},
		//	wantErr: false,
		//},
		{
			name:    "ExpectErrTest2",
			r:       Err[int](errors.New("error")),
			args:    args{"error"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.ExpectErr(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("ExpectErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResult_IsErr(t *testing.T) {
	type testCase[T any] struct {
		name string
		r    Result[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "IsErrTest1",
			r:    Ok(3),
			want: false,
		},
		{
			name: "IsErrTest2",
			r:    Err[int](errors.New("error")),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.IsErr(); got != tt.want {
				t.Errorf("IsErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_IsErrAnd(t *testing.T) {
	type args struct {
		f func(err error) bool
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args
		want bool
	}
	tests := []testCase[int]{
		{
			name: "IsErrAndTest1",
			r:    Ok(3),
			args: args{func(err error) bool { return true }},
			want: false,
		},
		{
			name: "IsErrAndTest2",
			r:    Err[int](errors.New("error")),
			args: args{func(err error) bool { return true }},
			want: true,
		},
		{
			name: "IsErrAndTest3",
			r:    Err[int](errors.New("error")),
			args: args{func(err error) bool { return false }},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.IsErrAnd(tt.args.f); got != tt.want {
				t.Errorf("IsErrAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_IsOk(t *testing.T) {
	type testCase[T any] struct {
		name string
		r    Result[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "IsOkTest1",
			r:    Ok(3),
			want: true,
		},
		{
			name: "IsOkTest2",
			r:    Err[int](errors.New("error")),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.IsOk(); got != tt.want {
				t.Errorf("IsOk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_IsOkAnd(t *testing.T) {
	type args struct {
		f func(t int) bool
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args
		want bool
	}
	tests := []testCase[int]{
		{
			name: "IsOkAndTest1",
			r:    Ok(3),
			args: args{func(t int) bool { return true }},
			want: true,
		},
		{
			name: "IsOkAndTest2",
			r:    Err[int](errors.New("error")),
			args: args{func(t int) bool { return true }},
			want: false,
		},
		{
			name: "IsOkAndTest3",
			r:    Ok(3),
			args: args{func(t int) bool { return false }},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.IsOkAnd(tt.args.f); got != tt.want {
				t.Errorf("IsOkAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_Map(t *testing.T) {
	type args[T any] struct {
		f func(t T) T
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args[T]
		want Result[T]
	}
	tests := []testCase[int]{
		{
			name: "MapTest1",
			r:    Ok(3),
			args: args[int]{func(t int) int { return 6 }},
			want: Ok(6),
		},
		{
			name: "MapTest2",
			r:    Err[int](errors.New("error")),
			args: args[int]{func(t int) int { return 6 }},
			want: Err[int](errors.New("error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Map(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_MapErr(t *testing.T) {
	type args struct {
		f func(err error) error
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args
		want Result[T]
	}
	tests := []testCase[int]{
		{
			name: "MapErrTest1",
			r:    Ok(3),
			args: args{func(err error) error { return errors.New("new error") }},
			want: Ok(3),
		},
		{
			name: "MapErrTest2",
			r:    Err[int](errors.New("error")),
			args: args{func(err error) error { return errors.New("new error") }},
			want: Err[int](errors.New("new error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.MapErr(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_MapOr(t *testing.T) {
	type args[T any] struct {
		dv T
		f  func(t T) T
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "MapOrTest1",
			r:    Ok(3),
			args: args[int]{6, func(t int) int { return 9 }},
			want: 9,
		},
		{
			name: "MapOrTest2",
			r:    Err[int](errors.New("error")),
			args: args[int]{6, func(t int) int { return 9 }},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.MapOr(tt.args.dv, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_MapOrElse(t *testing.T) {
	type args[T any] struct {
		d func(err error) T
		f func(t T) T
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "MapOrElseTest1",
			r:    Ok(3),
			args: args[int]{func(err error) int { return 6 }, func(t int) int { return 9 }},
			want: 9,
		},
		{
			name: "MapOrElseTest2",
			r:    Err[int](errors.New("error")),
			args: args[int]{func(err error) int { return 6 }, func(t int) int { return 9 }},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.MapOrElse(tt.args.d, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapOrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_Or(t *testing.T) {
	type args[T any] struct {
		res Result[T]
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args[T]
		want Result[T]
	}
	tests := []testCase[int]{
		{
			name: "OrTest1",
			r:    Ok(3),
			args: args[int]{Ok(6)},
			want: Ok(3),
		},
		{
			name: "OrTest2",
			r:    Err[int](errors.New("error")),
			args: args[int]{Ok(6)},
			want: Ok(6),
		},
		{
			name: "OrTest3",
			r:    Err[int](errors.New("error")),
			args: args[int]{Err[int](errors.New("new error"))},
			want: Err[int](errors.New("new error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Or(tt.args.res); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_Ok(t *testing.T) {
	type testCase[T any] struct {
		name string
		r    Result[T]
		want Option[T]
	}
	tests := []testCase[int]{
		{
			name: "OkTest1",
			r:    Ok(3),
			want: Some(3),
		},
		{
			name: "OkTest2",
			r:    Err[int](errors.New("error")),
			want: None[int](),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Ok(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ok() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_OrElse(t *testing.T) {
	type args[T any] struct {
		f func(err error) Result[T]
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args[T]
		want Result[T]
	}
	tests := []testCase[int]{
		{
			name: "OrElseTest1",
			r:    Ok(3),
			args: args[int]{func(err error) Result[int] { return Ok(6) }},
			want: Ok(3),
		},
		{
			name: "OrElseTest2",
			r:    Err[int](errors.New("error")),
			args: args[int]{func(err error) Result[int] { return Ok(6) }},
			want: Ok(6),
		},
		{
			name: "OrElseTest3",
			r:    Err[int](errors.New("error")),
			args: args[int]{func(err error) Result[int] { return Err[int](errors.New("new error")) }},
			want: Err[int](errors.New("new error")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.OrElse(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_Unwrap(t *testing.T) {
	type testCase[T any] struct {
		name string
		r    Result[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "UnwrapTest1",
			r:    Ok(3),
			want: 3,
		},
		//{
		//	name: "UnwrapTest2",
		//	r:    Err[int](errors.New("error")),
		//	want: 0,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.Unwrap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_UnwrapOr(t *testing.T) {
	type args[T any] struct {
		d T
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "UnwrapOrTest1",
			r:    Ok(3),
			args: args[int]{6},
			want: 3,
		},
		{
			name: "UnwrapOrTest2",
			r:    Err[int](errors.New("error")),
			args: args[int]{6},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.UnwrapOr(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_UnwrapOrDefault(t *testing.T) {
	type testCase[T any] struct {
		name string
		r    Result[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "UnwrapOrDefaultTest1",
			r:    Ok(3),
			want: 3,
		},
		{
			name: "UnwrapOrDefaultTest2",
			r:    Err[int](errors.New("error")),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.UnwrapOrDefault(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResult_UnwrapOrElse(t *testing.T) {
	type args[T any] struct {
		f func(err error) T
	}
	type testCase[T any] struct {
		name string
		r    Result[T]
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "UnwrapOrElseTest1",
			r:    Ok(3),
			args: args[int]{func(err error) int { return 6 }},
			want: 3,
		},
		{
			name: "UnwrapOrElseTest2",
			r:    Err[int](errors.New("error")),
			args: args[int]{func(err error) int { return 6 }},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.UnwrapOrElse(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapOrElse() = %v, want %v", got, tt.want)
			}
		})
	}
}
