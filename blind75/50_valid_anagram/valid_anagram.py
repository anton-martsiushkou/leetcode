def is_anagram(s: str, t: str) -> bool:
    """
    Checks if t is an anagram of s.
    Uses hash map for O(n) time complexity.

    Args:
        s: First string
        t: Second string

    Returns:
        True if t is an anagram of s, False otherwise
    """
    if len(s) != len(t):
        return False

    freq = {}

    for char in s:
        freq[char] = freq.get(char, 0) + 1

    for char in t:
        if char not in freq:
            return False
        freq[char] -= 1
        if freq[char] < 0:
            return False

    return True
