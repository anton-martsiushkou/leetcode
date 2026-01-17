from typing import List


def set_zeroes(matrix: List[List[int]]) -> None:
    """
    Sets entire row and column to 0 if an element is 0.
    Uses first row and column as markers for O(1) space complexity.
    Modifies the matrix in-place.

    Args:
        matrix: 2D list of integers
    """
    if not matrix or not matrix[0]:
        return

    m, n = len(matrix), len(matrix[0])
    first_row_has_zero = False
    first_col_has_zero = False

    # Check if first row has any zeros
    for j in range(n):
        if matrix[0][j] == 0:
            first_row_has_zero = True
            break

    # Check if first column has any zeros
    for i in range(m):
        if matrix[i][0] == 0:
            first_col_has_zero = True
            break

    # Use first row and column as markers
    for i in range(1, m):
        for j in range(1, n):
            if matrix[i][j] == 0:
                matrix[i][0] = 0
                matrix[0][j] = 0

    # Set zeros based on markers
    for i in range(1, m):
        for j in range(1, n):
            if matrix[i][0] == 0 or matrix[0][j] == 0:
                matrix[i][j] = 0

    # Handle first row
    if first_row_has_zero:
        for j in range(n):
            matrix[0][j] = 0

    # Handle first column
    if first_col_has_zero:
        for i in range(m):
            matrix[i][0] = 0
