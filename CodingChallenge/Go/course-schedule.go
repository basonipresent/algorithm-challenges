package main

/**
 * Course Schedule
 * Input: numCourses = 2, prerequisites = [[1,0]]
 * Output: true
 * Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So it is possible.
 */
func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := make([][]int, numCourses)
	for _, pair := range prerequisites {
		adjList[pair[1]] = append(adjList[pair[1]], pair[0])
	}

	visited := make([]bool, numCourses)
	path := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if dfsCourse(i, adjList, visited, path) {
			return false
		}
	}

	return true
}

func dfsCourse(node int, adjList [][]int, visited []bool, path []bool) bool {
	if path[node] {
		return true
	}
	if visited[node] {
		return false
	}

	visited[node] = true
	path[node] = true

	for _, neighbor := range adjList[node] {
		if dfsCourse(neighbor, adjList, visited, path) {
			return true
		}
	}

	path[node] = false
	return false
}
