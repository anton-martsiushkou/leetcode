from collections import Counter


def min_window(s: str, t: str) -> str:
    """
    Finds the minimum window substring containing all characters from t.
    Uses sliding window with hash maps for O(m+n) time complexity.

    Args:
        s: Source string to search in
        t: Target string with required characters

    Returns:
        Minimum window substring, or empty string if no valid window exists
    """
    if not s or not t or len(s) < len(t):
        return ""

    need = Counter(t)
    have = {}
    required = len(need)
    formed = 0

    left = 0
    min_len = float('inf')
    min_left = 0

    for right, char in enumerate(s):
        have[char] = have.get(char, 0) + 1

        if char in need and have[char] == need[char]:
            formed += 1

        while formed == required and left <= right:
            if right - left + 1 < min_len:
                min_len = right - left + 1
                min_left = left

            left_char = s[left]
            have[left_char] -= 1
            if left_char in need and have[left_char] < need[left_char]:
                formed -= 1
            left += 1

    if min_len == float('inf'):
        return ""
    return s[min_left:min_left + min_len]
