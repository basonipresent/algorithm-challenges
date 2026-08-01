package main

/**
 * Length of Longest Substring
 * Input: s = "abcabcbb"
 * Output: 3
 * Explanation: The answer is "abc", with the length of 3.
 */
func lengthOfLongestSubstring(s string) int {
	runes := []rune(s)
	left := 0
	maxLen := 0
	charMap := make(map[rune]int)

	for right := 0; right < len(runes); right++ {
		currentChar := runes[right]
		if duplicateIndex, ok := charMap[currentChar]; ok && duplicateIndex >= left {
			left = duplicateIndex + 1
		}
		charMap[currentChar] = right
		currentWindowSize := right - left + 1
		if currentWindowSize > maxLen {
			maxLen = currentWindowSize
		}
	}

	return maxLen
}
