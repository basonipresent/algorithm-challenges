package main

/**
 * Longest Repeating Character Replacement
 * Input: s = "AABABBA", k = 1
 * Output: 4
 * Explanation: Replace the one 'A' in the middle with 'B' and form the longest repeating character substring.
 */
func characterReplacement(s string, k int) int {
	mapS := make([]int, 128)
	start, end := 0, 0
	maxCount := 0
	maxLen := 0
	for end < len(s) {
		mapS[s[end]]++
		if mapS[s[end]] > maxCount {
			maxCount = mapS[s[end]]
		}
		end++

		if end-start-maxCount > k {
			mapS[s[start]]--
			start++
		}
		maxLen = max(maxLen, end-start)
	}
	return maxLen
}
