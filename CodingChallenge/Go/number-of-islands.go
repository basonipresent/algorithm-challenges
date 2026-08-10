package main

import "fmt"

/**
 * Number of Islands
 * Input: grid = [[1,1,1,1,0],[1,1,0,1,0],[1,1,0,0,0],[0,0,0,0,0]]
 * Output: 1
 * Explanation: There are 1 island in the grid.
 */
func numIslands(grid [][]byte) int {
	islands := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '1' {
				dfs(grid, row, col)
				islands++
			}
		}
	}

	return islands
}

func dfs(grid [][]byte, row, col int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) || grid[row][col] != '1' {
		return
	}

	grid[row][col] = '0'
	dfs(grid, row-1, col) // up
	dfs(grid, row+1, col) // down
	dfs(grid, row, col-1) // left
	dfs(grid, row, col+1) // right
}

func numberOfIslandsBFS(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	islands := 0
	rows := len(grid)
	cols := len(grid[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	set := make(map[string]bool)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' && !set[fmt.Sprintf("%d,%d", row, col)] {
				islands++
				bfs(grid, row, col, set, directions, rows, cols)
			}
		}
	}

	return islands
}

func bfs(grid [][]byte, row, col int, set map[string]bool, directions [][]int, rows, cols int) {
	queue := make([][]int, 0)
	queue = append(queue, []int{row, col})
	set[fmt.Sprintf("%d,%d", row, col)] = true

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		for _, d := range directions {
			newRow, newCol := cell[0]+d[0], cell[1]+d[1]
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				continue
			}
			key := fmt.Sprintf("%d,%d", newRow, newCol)
			if grid[newRow][newCol] == '1' && !set[key] {
				set[key] = true
				queue = append(queue, []int{newRow, newCol})
			}
		}
	}
}
