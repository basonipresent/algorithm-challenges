package main

import (
	"fmt"
	"sort"
	"testing"
)

// quadKeys normalizes a 4Sum result into a sorted list of comma-joined
// strings so tests can compare results regardless of quadruplet/result ordering.
func quadKeys(quads [][]int) []string {
	keys := make([]string, len(quads))
	for i, q := range quads {
		sorted := append([]int(nil), q...)
		sort.Ints(sorted)
		keys[i] = fmt.Sprintf("%d,%d,%d,%d", sorted[0], sorted[1], sorted[2], sorted[3])
	}
	sort.Strings(keys)
	return keys
}

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{
			name:   "example from doc comment",
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want:   [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name:   "all identical elements",
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			want:   [][]int{{2, 2, 2, 2}},
		},
		{
			name:   "fewer than four elements",
			nums:   []int{1, 2, 3},
			target: 6,
			want:   [][]int{},
		},
		{
			name:   "exactly four elements, matches",
			nums:   []int{1, 2, 3, 4},
			target: 10,
			want:   [][]int{{1, 2, 3, 4}},
		},
		{
			name:   "exactly four elements, no match",
			nums:   []int{1, 2, 3, 4},
			target: 100,
			want:   [][]int{},
		},
		{
			name:   "no combination sums to target",
			nums:   []int{1, 2, 3, 4, 5},
			target: 1000,
			want:   [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.nums, tt.target)
			gotKeys := quadKeys(got)
			wantKeys := quadKeys(tt.want)
			if len(gotKeys) != len(wantKeys) {
				t.Fatalf("fourSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
			for i := range gotKeys {
				if gotKeys[i] != wantKeys[i] {
					t.Errorf("fourSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
					break
				}
			}
		})
	}
}
