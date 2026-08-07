package main

/**
 * Two Sum II - Input Array Is Sorted
 * Input: numbers = [2,7,11,15], target = 9
 * Output: [1,2]
 */
func twoSumIISorted(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
