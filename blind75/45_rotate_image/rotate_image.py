from typing import List


def rotate(matrix: List[List[int]]) -> None:
    """
    Rotates the image by 90 degrees clockwise in-place.
    Uses transpose + reverse approach for O(1) space complexity.
    Modifies the matrix in-place.

    Args:
        matrix: n x n 2D list of integers
    """
    if not matrix or not matrix[0]:
        return

    n = len(matrix)

    # Transpose the matrix
    for i in range(n):
        for j in range(i + 1, n):
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

    # Reverse each row
    for i in range(n):
        matrix[i].reverse()
