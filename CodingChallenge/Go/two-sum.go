package main

import (
	"fmt"
	"os"
	"reflect"
)

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

func main() {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"example pair in middle", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"pair at the end", []int{3, 2, 4}, 6, []int{1, 2}},
		{"duplicate values summing to target", []int{3, 3}, 6, []int{0, 1}},
		{"negative numbers", []int{-3, 4, 3, 90}, 0, []int{0, 2}},
		{"no solution exists", []int{1, 2, 3}, 100, []int{}},
		{"empty input", []int{}, 0, []int{}},
		{"single element cannot pair with itself", []int{5}, 10, []int{}},
		{"zeroes summing to target zero", []int{0, 4, 0}, 0, []int{0, 2}},
	}

	failed := 0
	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if reflect.DeepEqual(got, tt.want) {
			fmt.Printf("PASS: %s\n", tt.name)
		} else {
			failed++
			fmt.Printf("FAIL: %s -> twoSum(%v, %d) = %v, want %v\n", tt.name, tt.nums, tt.target, got, tt.want)
		}
	}

	if failed > 0 {
		fmt.Printf("\n%d/%d tests failed\n", failed, len(tests))
		os.Exit(1)
	}
	fmt.Printf("\nall %d tests passed\n", len(tests))
}
