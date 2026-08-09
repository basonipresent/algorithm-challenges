package main

/**
 * Minimum Window Substring
 * Input: s = "ADOBECODEBANC", t = "ABC"
 * Output: "BANC"
 * Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
 */
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	mapT := make([]int, 128)
	count := len(t)
	start, end := 0, 0
	minLen, startIndex := int(^uint(0)>>1), 0

	for _, char := range t {
		mapT[char]++
	}

	for end < len(s) {
		if mapT[s[end]] > 0 {
			count--
		}
		mapT[s[end]]--
		end++

		for count == 0 {
			if end-start < minLen {
				startIndex = start
				minLen = end - start
			}

			if mapT[s[start]] == 0 {
				count++
			}
			mapT[s[start]]++
			start++
		}
	}

	if minLen == int(^uint(0)>>1) {
		return ""
	}

	return s[startIndex : startIndex+minLen]
}
