package main

/**
 * Koko Eating Bananas
 * Input: []int{3, 6, 7, 11}, 8
 * Output: 4
 * Explanation: Koko can eat 4 bananas per hour, and finish all the piles in 8 hours.
 */
func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 0
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}

	for left <= right {
		mid := left + (right-left)/2
		hoursNeeded := 0
		for _, pile := range piles {
			hoursNeeded += (pile + mid - 1) / mid
		}
		if hoursNeeded <= h {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
