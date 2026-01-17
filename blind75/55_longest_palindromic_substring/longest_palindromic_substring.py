def longest_palindrome(s: str) -> str:
    """
    Returns the longest palindromic substring in s.
    Uses expand around center approach for O(n²) time complexity.

    Args:
        s: Input string

    Returns:
        Longest palindromic substring
    """
    if len(s) < 2:
        return s

    start, max_len = 0, 0

    for i in range(len(s)):
        # Check for odd-length palindromes (center at i)
        len1 = expand_around_center(s, i, i)

        # Check for even-length palindromes (center between i and i+1)
        len2 = expand_around_center(s, i, i + 1)

        # Get the maximum length from both cases
        length = max(len1, len2)

        # Update if we found a longer palindrome
        if length > max_len:
            max_len = length
            start = i - (length - 1) // 2

    return s[start:start + max_len]


def expand_around_center(s: str, left: int, right: int) -> int:
    """
    Expands around the center and returns the length of the palindrome.

    Args:
        s: Input string
        left: Left pointer
        right: Right pointer

    Returns:
        Length of the palindrome found
    """
    while left >= 0 and right < len(s) and s[left] == s[right]:
        left -= 1
        right += 1

    return right - left - 1
