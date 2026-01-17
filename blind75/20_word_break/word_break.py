from typing import List


def word_break(s: str, word_dict: List[str]) -> bool:
    """
    Determines if string s can be segmented into dictionary words.
    Uses dynamic programming with O(n^2) time complexity.

    Args:
        s: String to segment
        word_dict: List of dictionary words

    Returns:
        True if s can be segmented into dictionary words, False otherwise
    """
    # Convert wordDict to set for O(1) lookup
    word_set = set(word_dict)

    n = len(s)
    # dp[i] = True if s[0:i] can be segmented
    dp = [False] * (n + 1)
    dp[0] = True  # Empty string can be segmented

    # Build DP table
    for i in range(1, n + 1):
        for j in range(i):
            # If s[0:j] can be segmented and s[j:i] is in dictionary
            if dp[j] and s[j:i] in word_set:
                dp[i] = True
                break

    return dp[n]
