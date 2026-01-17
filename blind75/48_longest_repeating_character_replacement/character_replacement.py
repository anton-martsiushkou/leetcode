def character_replacement(s: str, k: int) -> int:
    """
    Finds the length of longest substring with same letter after k replacements.
    Uses sliding window with frequency map for O(n) time complexity.

    Args:
        s: Input string of uppercase English letters
        k: Maximum number of character replacements allowed

    Returns:
        Length of the longest substring with same letter after replacements
    """
    freq = {}
    max_count = 0
    max_length = 0
    left = 0

    for right, char in enumerate(s):
        freq[char] = freq.get(char, 0) + 1
        max_count = max(max_count, freq[char])

        window_length = right - left + 1
        if window_length - max_count > k:
            freq[s[left]] -= 1
            left += 1
            window_length -= 1

        max_length = max(max_length, window_length)

    return max_length
