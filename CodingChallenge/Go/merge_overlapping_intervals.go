package main

import (
	"fmt"
	"sort"
)

/**
 * Merge Overlapping Intervals
 * Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
 * Output: [[1,6],[8,10],[15,18]]
 * Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
 */
func merge(intervals [][]int) [][]int {
	merged := make([][]int, 0, len(intervals))
	if len(intervals) == 0 {
		return merged
	}

	sorted := append([][]int(nil), intervals...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i][0] < sorted[j][0]
	})
	fmt.Printf("sorted: %v \n", sorted)

	merged = append(merged, sorted[0])
	for _, interval := range sorted[1:] {
		last := merged[len(merged)-1]
		fmt.Printf("last: %v, interval: %v \n", last, interval)
		if interval[0] <= last[1] {
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
			continue
		}
		merged = append(merged, interval)
	}
	return merged
}
