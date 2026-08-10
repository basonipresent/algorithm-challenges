package main

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{
			name:  "classic example",
			piles: []int{3, 6, 7, 11},
			h:     8,
			want:  4,
		},
		{
			name:  "hours exactly equal to number of piles",
			piles: []int{30, 11, 23, 4, 20},
			h:     5,
			want:  30,
		},
		{
			name:  "plenty of extra hours",
			piles: []int{30, 11, 23, 4, 20},
			h:     6,
			want:  23,
		},
		{
			name:  "single pile",
			piles: []int{5},
			h:     10,
			want:  1,
		},
		{
			name:  "single pile, exactly one hour",
			piles: []int{5},
			h:     1,
			want:  5,
		},
		{
			name:  "all piles the same size",
			piles: []int{4, 4, 4},
			h:     3,
			want:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minEatingSpeed(tt.piles, tt.h)
			if got != tt.want {
				t.Errorf("minEatingSpeed(%v, %d) = %d, want %d", tt.piles, tt.h, got, tt.want)
			}
		})
	}
}
