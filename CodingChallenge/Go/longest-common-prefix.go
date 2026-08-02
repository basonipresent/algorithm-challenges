package main

/**
 * Longest Common Prefix
 * Input: strs = ["flower","flow","flight"]
 * Output: "fl"
 * Explanation: The longest common prefix is "fl".
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, s := range strs {
		minLen := min(len(prefix), len(s))
		i := 0
		for i < minLen && prefix[i] == s[i] {
			i++
		}
		prefix = prefix[:i]
	}
	return prefix
}
