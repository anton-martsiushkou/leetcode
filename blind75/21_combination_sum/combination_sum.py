from typing import List


def combination_sum(candidates: List[int], target: int) -> List[List[int]]:
    """
    Finds all unique combinations that sum to target.
    Uses backtracking with index tracking to avoid duplicates.

    Args:
        candidates: List of distinct integers
        target: Target sum

    Returns:
        List of all unique combinations that sum to target
    """
    result = []
    path = []

    # Sort for early termination optimization
    candidates.sort()

    def backtrack(start: int, remaining: int) -> None:
        # Base case: found valid combination
        if remaining == 0:
            result.append(path[:])  # Add a copy of path
            return

        # Base case: target is negative
        if remaining < 0:
            return

        # Try each candidate starting from 'start'
        for i in range(start, len(candidates)):
            candidate = candidates[i]

            # Early termination: if candidate > remaining, skip (array is sorted)
            if candidate > remaining:
                break

            # Add candidate to current path
            path.append(candidate)

            # Recurse with same start index (can reuse same number)
            backtrack(i, remaining - candidate)

            # Backtrack: remove last candidate
            path.pop()

    backtrack(0, target)
    return result
