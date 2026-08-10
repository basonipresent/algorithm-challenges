package main

/*
 * Word Ladder
 * Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
 * Output: 5
 * Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog", return its length 5.
 */
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	if !wordSet[endWord] {
		return 0
	}

	beginSet := make(map[string]bool)
	beginSet[beginWord] = true
	endSet := make(map[string]bool)
	endSet[endWord] = true
	visited := make(map[string]bool)
	steps := 1

	for len(beginSet) > 0 && len(endSet) > 0 {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}

		nextWordSet := make(map[string]bool)
		for word := range beginSet {
			for i := 0; i < len(word); i++ {
				for ch := 'a'; ch <= 'z'; ch++ {
					if ch == rune(word[i]) {
						continue
					}
					newWord := word[:i] + string(ch) + word[i+1:]
					if endSet[newWord] {
						return steps + 1
					}
					if wordSet[newWord] && !visited[newWord] {
						visited[newWord] = true
						nextWordSet[newWord] = true
					}
				}
			}
		}
		beginSet = nextWordSet
		steps++
	}
	return 0
}
