package leetcode

import (
	"reflect"
	"testing"
)

func Test_IsPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	var tests []struct {
		name string
		args args
		want bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitInt(t *testing.T) {
	type args struct {
		n int
	}
	var tests []struct {
		name string
		args args
		want []int
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitInt(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
