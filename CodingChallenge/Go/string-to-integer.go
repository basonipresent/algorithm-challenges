package main

import "math"

/**
 * String to Integer (atoi)
 * Input: s = "42"
 * Output: 42
 * Explanation: The first non-whitespace character is '42', which is a numerical digit.
 *             Then we take as many characters as possible, which results in 42.
 */
func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	// skip leading spaces
	i, n := 0, len(s)
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n {
		return 0
	}

	// check sign
	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}

	// convert digits to integer
	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')
		result = result*10 + digit
		if sign*result > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign*result < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}
	return sign * result
}
