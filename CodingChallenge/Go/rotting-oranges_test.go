package main

import "testing"

func cloneIntGrid(grid [][]int) [][]int {
	out := make([][]int, len(grid))
	for i, row := range grid {
		out[i] = append([]int(nil), row...)
	}
	return out
}

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example from doc comment",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			want: 4,
		},
		{
			name: "isolated fresh orange never rots",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			want: -1,
		},
		{
			name: "no fresh oranges at all",
			grid: [][]int{
				{0, 2},
			},
			want: 0,
		},
		{
			name: "no oranges at all",
			grid: [][]int{
				{0, 0, 0},
			},
			want: 0,
		},
		{
			name: "single fresh orange, no rotten source",
			grid: [][]int{
				{1},
			},
			want: -1,
		},
		{
			name: "single rotten orange only",
			grid: [][]int{
				{2},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gridCopy := cloneIntGrid(tt.grid)
			got := orangesRotting(gridCopy)
			if got != tt.want {
				t.Errorf("orangesRotting(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}
