package course_schedule

// CanFinish determines if all courses can be completed using DFS cycle detection
// Time: O(V + E), Space: O(V + E)
func CanFinish(numCourses int, prerequisites [][]int) bool {
	// Build adjacency list
	graph := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		course, prerequisite := prereq[0], prereq[1]
		graph[prerequisite] = append(graph[prerequisite], course)
	}

	// 0 = unvisited, 1 = visiting, 2 = visited
	state := make([]int, numCourses)

	var hasCycle func(course int) bool
	hasCycle = func(course int) bool {
		if state[course] == 1 {
			// Currently visiting - found a cycle
			return true
		}
		if state[course] == 2 {
			// Already visited
			return false
		}

		// Mark as visiting
		state[course] = 1

		// Check all neighbors
		for _, neighbor := range graph[course] {
			if hasCycle(neighbor) {
				return true
			}
		}

		// Mark as visited
		state[course] = 2
		return false
	}

	// Check each course
	for i := 0; i < numCourses; i++ {
		if hasCycle(i) {
			return false
		}
	}

	return true
}

// CanFinishBFS determines if all courses can be completed using Kahn's algorithm (BFS)
// Time: O(V + E), Space: O(V + E)
func CanFinishBFS(numCourses int, prerequisites [][]int) bool {
	// Build adjacency list and in-degree array
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, prereq := range prerequisites {
		course, prerequisite := prereq[0], prereq[1]
		graph[prerequisite] = append(graph[prerequisite], course)
		inDegree[course]++
	}

	// Find all courses with no prerequisites
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	coursesCompleted := 0

	// Process courses
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		coursesCompleted++

		// Reduce in-degree for all dependent courses
		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return coursesCompleted == numCourses
}
