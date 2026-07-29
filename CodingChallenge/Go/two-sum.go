package main

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
