from typing import List


def pacific_atlantic(heights: List[List[int]]) -> List[List[int]]:
    """
    Finds cells that can flow to both oceans using DFS.
    Time: O(m × n), Space: O(m × n)

    Args:
        heights: Matrix of cell heights

    Returns:
        List of [row, col] coordinates that can reach both oceans
    """
    if not heights or not heights[0]:
        return []

    rows, cols = len(heights), len(heights[0])
    pacific = [[False] * cols for _ in range(rows)]
    atlantic = [[False] * cols for _ in range(rows)]

    def dfs(r: int, c: int, visited: List[List[bool]], prev_height: int):
        # Check bounds and visited
        if (r < 0 or r >= rows or c < 0 or c >= cols or
                visited[r][c] or heights[r][c] < prev_height):
            return

        visited[r][c] = True

        # Explore 4 directions
        dfs(r + 1, c, visited, heights[r][c])
        dfs(r - 1, c, visited, heights[r][c])
        dfs(r, c + 1, visited, heights[r][c])
        dfs(r, c - 1, visited, heights[r][c])

    # DFS from Pacific borders (top and left)
    for c in range(cols):
        dfs(0, c, pacific, heights[0][c])
        dfs(rows - 1, c, atlantic, heights[rows - 1][c])

    for r in range(rows):
        dfs(r, 0, pacific, heights[r][0])
        dfs(r, cols - 1, atlantic, heights[r][cols - 1])

    # Find intersection
    result = []
    for r in range(rows):
        for c in range(cols):
            if pacific[r][c] and atlantic[r][c]:
                result.append([r, c])

    return result


def pacific_atlantic_bfs(heights: List[List[int]]) -> List[List[int]]:
    """
    Finds cells that can flow to both oceans using BFS.
    Time: O(m × n), Space: O(m × n)

    Args:
        heights: Matrix of cell heights

    Returns:
        List of [row, col] coordinates that can reach both oceans
    """
    if not heights or not heights[0]:
        return []

    rows, cols = len(heights), len(heights[0])
    pacific = [[False] * cols for _ in range(rows)]
    atlantic = [[False] * cols for _ in range(rows)]

    def bfs(queue: List[List[int]], visited: List[List[bool]]):
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        while queue:
            r, c = queue.pop(0)

            for dr, dc in directions:
                nr, nc = r + dr, c + dc
                if (0 <= nr < rows and 0 <= nc < cols and
                        not visited[nr][nc] and
                        heights[nr][nc] >= heights[r][c]):
                    visited[nr][nc] = True
                    queue.append([nr, nc])

    # Initialize queues with border cells
    pacific_queue = []
    atlantic_queue = []

    for c in range(cols):
        pacific[0][c] = True
        pacific_queue.append([0, c])
        atlantic[rows - 1][c] = True
        atlantic_queue.append([rows - 1, c])

    for r in range(rows):
        pacific[r][0] = True
        pacific_queue.append([r, 0])
        atlantic[r][cols - 1] = True
        atlantic_queue.append([r, cols - 1])

    bfs(pacific_queue, pacific)
    bfs(atlantic_queue, atlantic)

    # Find intersection
    result = []
    for r in range(rows):
        for c in range(cols):
            if pacific[r][c] and atlantic[r][c]:
                result.append([r, c])

    return result
