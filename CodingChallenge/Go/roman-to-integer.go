package main

/**
* Input: s = "MCMXCIV"
* Output: 1994
* Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
 */
type romanNumeral struct {
	symbol string
	value  int
}

var romanNumeralMap = map[string]int{
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}

func romanToInt(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		value := romanNumeralMap[string(s[i])]
		if i+1 < len(s) && value < romanNumeralMap[string(s[i+1])] {
			result -= value
		} else {
			result += value
		}
	}
	return result
}
