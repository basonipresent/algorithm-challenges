package main

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "classic example",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "two equal-height lines at the ends",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "strictly increasing heights",
			height: []int{1, 2, 3, 4, 5},
			want:   6,
		},
		{
			name:   "strictly decreasing heights",
			height: []int{5, 4, 3, 2, 1},
			want:   6,
		},
		{
			name:   "tallest lines are adjacent in the middle",
			height: []int{1, 2, 4, 3},
			want:   4,
		},
		{
			name:   "all same height",
			height: []int{4, 4, 4, 4},
			want:   12,
		},
		{
			name:   "some zero-height lines",
			height: []int{0, 2, 0, 4, 0},
			want:   4,
		},
		{
			name:   "two elements only",
			height: []int{3, 7},
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Errorf("maxArea(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
