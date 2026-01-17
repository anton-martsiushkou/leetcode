def length_of_longest_substring(s: str) -> int:
    """
    Finds the length of the longest substring without repeating characters.
    Uses sliding window with hash map for O(n) time complexity.

    Args:
        s: Input string

    Returns:
        Length of the longest substring without repeating characters
    """
    char_index = {}
    max_length = 0
    left = 0

    for right, char in enumerate(s):
        if char in char_index and char_index[char] >= left:
            left = char_index[char] + 1
        char_index[char] = right
        current_length = right - left + 1
        max_length = max(max_length, current_length)

    return max_length
