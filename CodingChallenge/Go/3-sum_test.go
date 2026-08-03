package main

import (
	"fmt"
	"sort"
	"testing"
)

// tripletKeys normalizes a 3Sum result into a sorted list of comma-joined
// strings so tests can compare results regardless of triplet/result ordering.
func tripletKeys(triplets [][]int) []string {
	keys := make([]string, len(triplets))
	for i, tr := range triplets {
		keys[i] = fmt.Sprintf("%d,%d,%d", tr[0], tr[1], tr[2])
	}
	sort.Strings(keys)
	return keys
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "example from doc comment",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "no triplet sums to zero",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "empty input",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "extra duplicate zeros collapse to one triplet",
			nums: []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "duplicate values must not produce duplicate triplets",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
		{
			name: "small mixed set",
			nums: []int{1, -1, -1, 0},
			want: [][]int{{-1, 0, 1}},
		},
		{
			name: "multiple distinct triplets",
			nums: []int{3, 0, -2, -1, 1, 2},
			want: [][]int{{-2, -1, 3}, {-2, 0, 2}, {-1, 0, 1}},
		},
		{
			name: "fewer than three elements",
			nums: []int{1, 2},
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			gotKeys := tripletKeys(got)
			wantKeys := tripletKeys(tt.want)
			if len(gotKeys) != len(wantKeys) {
				t.Fatalf("threeSum(%v) = %v, want %v", tt.nums, got, tt.want)
			}
			for i := range gotKeys {
				if gotKeys[i] != wantKeys[i] {
					t.Errorf("threeSum(%v) = %v, want %v", tt.nums, got, tt.want)
					break
				}
			}
		})
	}
}
