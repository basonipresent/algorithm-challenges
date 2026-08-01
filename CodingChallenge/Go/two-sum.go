package main

/**
 * Two Sum
 * Input: nums = [2,7,11,15], target = 9
 * Output: [0,1]
 * Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
 */
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, v := range nums {
		complement := target - v
		if j, ok := seen[complement]; ok && j != i {
			return []int{j, i}
		}
		seen[v] = i
	}
	return []int{}
}
