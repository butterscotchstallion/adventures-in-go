package leetcode

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	type testCase struct {
		name string
		args args
		want bool
	}
	var testCases = make([]testCase, 3)
	testCases = append(testCases, testCase{name: "case one", args: args{s: "()"}, want: true})

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
