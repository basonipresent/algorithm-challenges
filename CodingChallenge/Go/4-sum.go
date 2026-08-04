package main

import "sort"

/**
 * Input: nums = [1,0,-1,0,-2,2], target = 0
 * Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 */
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	if len(nums) < 4 {
		return nil
	}
	result := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			result = append(result, findPairs(nums, target-nums[i]-nums[j], i, j)...)
		}
	}
	return result
}

// findPairs finds every pair within nums[j+1:] that, together with nums[i]
// and nums[j], sums to target using the two-pointer technique on the
// already-sorted slice.
func findPairs(nums []int, remaining, i, j int) [][]int {
	pairs := make([][]int, 0)
	left, right := j+1, len(nums)-1
	for left < right {
		total := nums[left] + nums[right]
		switch {
		case total == remaining:
			pairs = append(pairs, []int{nums[i], nums[j], nums[left], nums[right]})
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		case total > remaining:
			right--
		default:
			left++
		}
	}
	return pairs
}
