package aocutils

import (
	"reflect"
	"testing"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test that it outputs 'Hello'",
			want: "Hello.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinMaxInt(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{
			name:  "Test simple array.",
			args:  args{[]int{0, 1, 2}},
			want:  0,
			want1: 2,
		},
		{
			name:  "Test two element array.",
			args:  args{[]int{0, 1}},
			want:  0,
			want1: 1,
		},
		{
			name:  "Test singel element array.",
			args:  args{[]int{0}},
			want:  0,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MinMaxInt(tt.args.array)
			if got != tt.want {
				t.Errorf("MinMaxInt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MinMaxInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReadfile(t *testing.T) {
	type args struct {
		filename string
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
			if got := Readfile(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Readfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
