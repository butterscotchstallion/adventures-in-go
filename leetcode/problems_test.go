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

	isPalindrome1 := IsPalindrome(1)
	if !isPalindrome1 {
		t.Errorf("isPalindrome1 should be true")
	}

	isPalindrome2 := IsPalindrome(13)
	if isPalindrome2 {
		t.Errorf("isPalindrome2 should be false")
	}

	isPalindrome3 := IsPalindrome(9999999)
	if !isPalindrome3 {
		t.Errorf("isPalindrome3 should be false")
	}

	isPalindrome4 := IsPalindrome(-121)
	if isPalindrome4 {
		t.Errorf("isPalindrome4 should be false")
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
