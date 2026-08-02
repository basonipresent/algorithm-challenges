package main

/**
 * Input: num = 3749
 * Output: "MMMDCCXLIX"
 * Explanation:
 * 3000 = MMM as 1000 (M) + 1000 (M) + 1000 (M)
 * 700 = DCC as 500 (D) + 100 (C) + 100 (C)
 * 40 = XL as 10 (X) less of 50 (L)
 * 9 = IX as 1 (I) less of 10 (X)
 * Note: 49 is not 1 (I) less of 50 (L) because the conversion is based on decimal places
 **/

type romanEntry struct {
	value  int
	symbol string
}

var romanTable = []romanEntry{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func intToRoman(num int) string {
	result := ""
	for _, entry := range romanTable {
		if num == 0 {
			break
		}
		for num >= entry.value {
			result += entry.symbol
			num -= entry.value
		}
	}
	return result
}
