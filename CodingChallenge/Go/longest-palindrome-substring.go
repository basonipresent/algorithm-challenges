package main

/**
 * Input: s = "babad"
 * Expected: "bab"
 */
func longestPalindrome(s string) string {
	runes := []rune(s)
	if len(runes) == 0 {
		return ""
	}

	start, maxLen := 0, 1

	expand := func(left, right int) {
		for left >= 0 && right < len(runes) && runes[left] == runes[right] {
			left--
			right++
		}
		length := right - left - 1
		if length > maxLen {
			maxLen = length
			start = left + 1
		}
	}

	for i := range runes {
		expand(i, i)   // odd-length palindromes centered on i
		expand(i, i+1) // even-length palindromes centered between i and i+1
	}

	return string(runes[start : start+maxLen])
}
