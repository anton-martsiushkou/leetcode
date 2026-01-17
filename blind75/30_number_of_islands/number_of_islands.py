from typing import List
from collections import deque


def num_islands(grid: List[List[str]]) -> int:
    """
    Counts islands using DFS.
    Time: O(m × n), Space: O(m × n)

    Args:
        grid: 2D grid of '1' (land) and '0' (water)

    Returns:
        Number of islands
    """
    if not grid or not grid[0]:
        return 0

    rows, cols = len(grid), len(grid[0])
    count = 0

    def dfs(r: int, c: int):
        # Check bounds and if it's land
        if r < 0 or r >= rows or c < 0 or c >= cols or grid[r][c] == '0':
            return

        # Mark as visited
        grid[r][c] = '0'

        # Explore 4 directions
        dfs(r + 1, c)
        dfs(r - 1, c)
        dfs(r, c + 1)
        dfs(r, c - 1)

    # Scan the grid
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == '1':
                count += 1
                dfs(r, c)

    return count


def num_islands_bfs(grid: List[List[str]]) -> int:
    """
    Counts islands using BFS.
    Time: O(m × n), Space: O(min(m, n))

    Args:
        grid: 2D grid of '1' (land) and '0' (water)

    Returns:
        Number of islands
    """
    if not grid or not grid[0]:
        return 0

    rows, cols = len(grid), len(grid[0])
    count = 0
    directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

    def bfs(start_r: int, start_c: int):
        queue = deque([(start_r, start_c)])
        grid[start_r][start_c] = '0'

        while queue:
            r, c = queue.popleft()

            for dr, dc in directions:
                nr, nc = r + dr, c + dc
                if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == '1':
                    grid[nr][nc] = '0'
                    queue.append((nr, nc))

    # Scan the grid
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == '1':
                count += 1
                bfs(r, c)

    return count
