package main

/**
 * Palindrome Number
 * Input: x = 121
 * Output: true
 * Explanation: 121 is a palindrome.
 */
func isPalindrome(x int) bool {
	// negative numbers are not palindromes
	if x < 0 {
		return false
	}
	// reverse x value
	original := x
	reversed := 0
	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return reversed == original
}
