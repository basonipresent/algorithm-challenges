package main

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "example from doc comment",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "another classic example",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
		{
			name:   "strictly increasing, no trapped water",
			height: []int{1, 2, 3, 4},
			want:   0,
		},
		{
			name:   "strictly decreasing, no trapped water",
			height: []int{4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "flat terrain, no trapped water",
			height: []int{3, 3, 3},
			want:   0,
		},
		{
			name:   "empty input",
			height: []int{},
			want:   0,
		},
		{
			name:   "single bar",
			height: []int{5},
			want:   0,
		},
		{
			name:   "two bars",
			height: []int{5, 1},
			want:   0,
		},
		{
			name:   "single deep valley",
			height: []int{5, 0, 5},
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.height)
			if got != tt.want {
				t.Errorf("trap(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
