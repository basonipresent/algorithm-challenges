package main

/**
 * Rotting Oranges
 * Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
 * Output: 4
 * Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally.
 */
func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	queue := make([][]int, 0)
	fresh := 0

	for i := range rows {
		for j := range cols {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
			if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	minutes := -1
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		minutes++
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			row, col := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, dir := range directions {
				newRow := row + dir[0]
				newCol := col + dir[1]
				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2
					fresh--
					queue = append(queue, []int{newRow, newCol})
				}
			}
		}
	}

	if fresh == 0 {
		return minutes
	}
	return -1
}
