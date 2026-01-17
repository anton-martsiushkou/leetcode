def count_substrings(s: str) -> int:
    """
    Returns the number of palindromic substrings in s.
    Uses expand around center approach for O(n²) time complexity.

    Args:
        s: Input string

    Returns:
        Number of palindromic substrings
    """
    if not s:
        return 0

    count = 0

    for i in range(len(s)):
        # Count odd-length palindromes (center at i)
        count += expand_around_center(s, i, i)

        # Count even-length palindromes (center between i and i+1)
        count += expand_around_center(s, i, i + 1)

    return count


def expand_around_center(s: str, left: int, right: int) -> int:
    """
    Expands around the center and counts palindromes.

    Args:
        s: Input string
        left: Left pointer
        right: Right pointer

    Returns:
        Number of palindromes found during expansion
    """
    count = 0

    while left >= 0 and right < len(s) and s[left] == s[right]:
        count += 1
        left -= 1
        right += 1

    return count
