package main

/*
 * Input: digits = "23"
 * Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
 */
func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	backtrack(0, digits, "", &result)

	return result
}

func backtrack(index int, digits string, current string, result *[]string) {
	if index == len(digits) {
		*result = append(*result, current)
		return
	}
	c := string(digits[index])
	for _, letter := range digitToLettersMap[c] {
		backtrack(index+1, digits, current+letter, result)
	}
}

var digitToLettersMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}
