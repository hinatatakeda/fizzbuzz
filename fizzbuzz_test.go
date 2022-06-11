package example

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {

	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "1〜20の値でFizzBuzzテスト",
			args: args{
				numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			},
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz"},
		}}
	///{
	//	name: "0以下の数字は変換されずそのまま返ってくること",
	//	args: args{
	//		numbers: []int{0, -1, -2, -3, -4, -5},
	//	},
	//	want: []string{"0", "-1", "-2", "-3", "-4", "-5"},
	///},

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendNum(t *testing.T) {
	type args struct {
		s []string
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendNum(tt.args.s, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBuzz(tt.args.n); got != tt.want {
				t.Errorf("isBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFizz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFizz(tt.args.n); got != tt.want {
				t.Errorf("isFizz() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFizzBuzz(tt.args.n); got != tt.want {
				t.Errorf("isFizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
