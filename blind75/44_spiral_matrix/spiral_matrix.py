from typing import List


def spiral_order(matrix: List[List[int]]) -> List[int]:
    """
    Returns all elements of the matrix in spiral order.
    Uses boundary tracking for O(1) space complexity.

    Args:
        matrix: 2D list of integers

    Returns:
        List of integers in spiral order
    """
    if not matrix or not matrix[0]:
        return []

    m, n = len(matrix), len(matrix[0])
    result = []
    top, bottom, left, right = 0, m - 1, 0, n - 1

    while top <= bottom and left <= right:
        # Move right along top row
        for j in range(left, right + 1):
            result.append(matrix[top][j])
        top += 1

        # Move down along right column
        for i in range(top, bottom + 1):
            result.append(matrix[i][right])
        right -= 1

        # Move left along bottom row (if there's a row remaining)
        if top <= bottom:
            for j in range(right, left - 1, -1):
                result.append(matrix[bottom][j])
            bottom -= 1

        # Move up along left column (if there's a column remaining)
        if left <= right:
            for i in range(bottom, top - 1, -1):
                result.append(matrix[i][left])
            left += 1

    return result
