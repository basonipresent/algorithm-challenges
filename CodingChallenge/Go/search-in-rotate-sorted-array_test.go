package main

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "target found in rotated right half",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "target found in rotated left half",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 5,
			want:   1,
		},
		{
			name:   "target not present",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "empty array",
			nums:   []int{},
			target: 5,
			want:   -1,
		},
		{
			name:   "single element found",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "single element not found",
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
		{
			name:   "not rotated, target found",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "rotation pivot at the start",
			nums:   []int{1, 2, 3, 4, 5},
			target: 1,
			want:   0,
		},
		{
			name:   "target at the very end",
			nums:   []int{5, 1, 2, 3, 4},
			target: 4,
			want:   4,
		},
		{
			name:   "two elements, rotated",
			nums:   []int{3, 1},
			target: 1,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
