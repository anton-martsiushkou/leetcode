from typing import List


def exist(board: List[List[str]], word: str) -> bool:
    """
    Returns true if word exists in the grid.
    Uses backtracking with DFS to explore all possible paths.

    Args:
        board: 2D list of characters
        word: Target word to search for

    Returns:
        True if word exists in the grid, False otherwise
    """
    if not board or not board[0] or not word:
        return False

    m, n = len(board), len(board[0])

    def dfs(i: int, j: int, index: int) -> bool:
        # Found the complete word
        if index == len(word):
            return True

        # Out of bounds or character mismatch or already visited
        if i < 0 or i >= m or j < 0 or j >= n or board[i][j] != word[index]:
            return False

        # Mark as visited by temporarily modifying the cell
        temp = board[i][j]
        board[i][j] = '#'

        # Explore all 4 directions
        found = (
            dfs(i + 1, j, index + 1) or
            dfs(i - 1, j, index + 1) or
            dfs(i, j + 1, index + 1) or
            dfs(i, j - 1, index + 1)
        )

        # Backtrack: restore the cell
        board[i][j] = temp

        return found

    # Try each cell as a starting point
    for i in range(m):
        for j in range(n):
            if dfs(i, j, 0):
                return True

    return False
