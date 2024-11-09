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
	openParenCount := 0
	closeParenCount := 0
	openBracketCount := 0
	closeBracketCount := 0
	openBraceCount := 0
	closeBraceCount := 0

	if len(s) == 1 {
		return false
	}

	for pos, char := range s {
		if char == '(' {
			openParenCount++
		}
		if char == ')' {
			if openParenCount == closeParenCount {
				return false
			}
			closeParenCount++
		}
		if char == '[' {
			openBracketCount++
			if pos < len(s)-1 && s[pos+1] != ']' {
				return false
			}
		}
		if char == ']' {
			if openBracketCount == closeBracketCount {
				return false
			}
			closeBracketCount++
		}
		if char == '{' {
			openBraceCount++
			if pos < len(s)-1 && s[pos+1] != '}' {
				return false
			}
		}
		if char == '}' {
			if openBraceCount == closeBraceCount {
				return false
			}
			closeBraceCount++
		}
	}
	return openParenCount == closeParenCount && openBraceCount == closeBraceCount
}
