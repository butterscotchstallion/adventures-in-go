package leetcode

import (
	"strconv"
	"unicode/utf8"
)

func SplitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n /= 10
	}
	result := []int{}
	for i := range slc {
		result = append(result, slc[len(slc)-1-i])
	}
	return result
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func IsPalindrome(x int) bool {
	strInput := strconv.Itoa(x)
	reversed := Reverse(strInput)
	return reversed == strconv.Itoa(x)
}

func IsPalindrome2(x int) bool {
	digits := SplitInt(x)
	if x < 0 {
		return false
	}
	if len(digits) == 1 {
		return true
	}

	reversed := []int{}

	for i := len(digits) - 1; i >= 0; i-- {
		reversed = append(reversed, digits[i])
	}
	for j := len(reversed) - 1; j >= 0; j-- {
		if reversed[j] != digits[j] {
			return false
		}
	}
	return true
}
