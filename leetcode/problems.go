package leetcode

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

func IsPalindrome(x int) bool {
	digits := SplitInt(x)
	reversed := []int{}

	if x < 0 {
		return false
	}

	if len(digits) == 1 {
		return true
	}

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
