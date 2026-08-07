package main

/**
 * Group Anagrams
 * Input: ["eat","tea","tan","ate","nat","bat"]
 * Output: [["ate","eat","tea"], ["nat","tan"], ["bat"]]
 */
func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]int][]string)
	for _, word := range strs {
		count := [26]int{}
		for _, char := range word {
			count[char-'a']++
		}
		anagrams[count] = append(anagrams[count], word)
	}

	result := make([][]string, 0, len(anagrams))
	for _, ans := range anagrams {
		result = append(result, ans)
	}
	return result
}
