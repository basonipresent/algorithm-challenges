package main

/**
 * Valid Parentheses
 * Input: s = "()[]{}"
 * Output: true
 */
func isValid(s string) bool {
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]rune, 0, len(s))
	for _, c := range s {
		switch c {
		case '(', '[', '{': // opening parentheses
			stack = append(stack, c)
		case ')', ']', '}': // closing parentheses
			if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
