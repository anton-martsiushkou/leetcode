from typing import List
from collections import defaultdict, deque


def can_finish(num_courses: int, prerequisites: List[List[int]]) -> bool:
    """
    Determines if all courses can be completed using DFS cycle detection.
    Time: O(V + E), Space: O(V + E)

    Args:
        num_courses: Total number of courses
        prerequisites: List of [course, prerequisite] pairs

    Returns:
        True if all courses can be completed, False otherwise
    """
    # Build adjacency list
    graph = defaultdict(list)
    for course, prereq in prerequisites:
        graph[prereq].append(course)

    # 0 = unvisited, 1 = visiting, 2 = visited
    state = [0] * num_courses

    def has_cycle(course: int) -> bool:
        if state[course] == 1:
            # Currently visiting - found a cycle
            return True
        if state[course] == 2:
            # Already visited
            return False

        # Mark as visiting
        state[course] = 1

        # Check all neighbors
        for neighbor in graph[course]:
            if has_cycle(neighbor):
                return True

        # Mark as visited
        state[course] = 2
        return False

    # Check each course
    for i in range(num_courses):
        if has_cycle(i):
            return False

    return True


def can_finish_bfs(num_courses: int, prerequisites: List[List[int]]) -> bool:
    """
    Determines if all courses can be completed using Kahn's algorithm (BFS).
    Time: O(V + E), Space: O(V + E)

    Args:
        num_courses: Total number of courses
        prerequisites: List of [course, prerequisite] pairs

    Returns:
        True if all courses can be completed, False otherwise
    """
    # Build adjacency list and in-degree array
    graph = defaultdict(list)
    in_degree = [0] * num_courses

    for course, prereq in prerequisites:
        graph[prereq].append(course)
        in_degree[course] += 1

    # Find all courses with no prerequisites
    queue = deque([i for i in range(num_courses) if in_degree[i] == 0])

    courses_completed = 0

    # Process courses
    while queue:
        current = queue.popleft()
        courses_completed += 1

        # Reduce in-degree for all dependent courses
        for neighbor in graph[current]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)

    return courses_completed == num_courses
