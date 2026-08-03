package main

import "testing"

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "example from doc comment",
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		{
			name:   "all zeros",
			nums:   []int{0, 0, 0},
			target: 1,
			want:   0,
		},
		{
			name:   "only one distinct combination possible",
			nums:   []int{1, 1, 1, 1},
			target: 0,
			want:   3,
		},
		{
			name:   "exact match available",
			nums:   []int{1, 1, -1, -1, 3},
			target: -1,
			want:   -1,
		},
		{
			name:   "smallest valid input, exactly three elements",
			nums:   []int{0, 1, 2},
			target: 0,
			want:   3,
		},
		{
			name:   "target far above all possible sums",
			nums:   []int{1, 1, 1, 0},
			target: 100,
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSumClosest(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("threeSumClosest(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
