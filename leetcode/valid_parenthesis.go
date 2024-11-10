package leetcode

/**
 * An input string is valid if:
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * Every close bracket has a corresponding open bracket of the same type.
 *
 */
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, r := range s {
		if _, ok := pairs[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
