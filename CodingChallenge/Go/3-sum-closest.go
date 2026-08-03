package main

import (
	"math"
	"sort"
)

/**
 * 3Sum Closest
 * Input: nums = [-1,2,1,-4], target = 1
 * Output: 2
 * Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 */
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]

			if math.Abs(float64(sum-target)) < math.Abs(float64(closest-target)) {
				closest = sum
			}

			if sum == target {
				return target
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return closest
}
