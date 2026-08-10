package main

import "testing"

func cloneGrid(grid [][]byte) [][]byte {
	out := make([][]byte, len(grid))
	for i, row := range grid {
		out[i] = append([]byte(nil), row...)
	}
	return out
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "example from doc comment, single connected island",
			grid: [][]byte{
				[]byte("11110"),
				[]byte("11010"),
				[]byte("11000"),
				[]byte("00000"),
			},
			want: 1,
		},
		{
			name: "multiple disjoint islands",
			grid: [][]byte{
				[]byte("11000"),
				[]byte("11000"),
				[]byte("00100"),
				[]byte("00011"),
			},
			want: 3,
		},
		{
			name: "all water",
			grid: [][]byte{
				[]byte("0000"),
				[]byte("0000"),
			},
			want: 0,
		},
		{
			name: "all land, single island",
			grid: [][]byte{
				[]byte("111"),
				[]byte("111"),
			},
			want: 1,
		},
		{
			name: "single cell land",
			grid: [][]byte{
				[]byte("1"),
			},
			want: 1,
		},
		{
			name: "single cell water",
			grid: [][]byte{
				[]byte("0"),
			},
			want: 0,
		},
		{
			name: "diagonal cells are not connected",
			grid: [][]byte{
				[]byte("10"),
				[]byte("01"),
			},
			want: 2,
		},
		{
			name: "empty grid",
			grid: [][]byte{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run("dfs/"+tt.name, func(t *testing.T) {
			gridCopy := cloneGrid(tt.grid)
			got := numIslands(gridCopy)
			if got != tt.want {
				t.Errorf("numIslands(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
		t.Run("bfs/"+tt.name, func(t *testing.T) {
			gridCopy := cloneGrid(tt.grid)
			got := numberOfIslandsBFS(gridCopy)
			if got != tt.want {
				t.Errorf("numberOfIslandsBFS(%v) = %d, want %d", tt.grid, got, tt.want)
			}
		})
	}
}
