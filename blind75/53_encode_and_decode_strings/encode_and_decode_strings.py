from typing import List


def encode(strs: List[str]) -> str:
    """
    Encodes a list of strings to a single string.
    Uses length-prefix encoding for O(n) time complexity.

    Args:
        strs: List of strings to encode

    Returns:
        Encoded string with format: length#string for each input string
    """
    result = []
    for s in strs:
        result.append(f"{len(s)}#{s}")
    return "".join(result)


def decode(s: str) -> List[str]:
    """
    Decodes a single string back to a list of strings.
    Uses length-prefix decoding for O(n) time complexity.

    Args:
        s: Encoded string

    Returns:
        List of decoded strings
    """
    if not s:
        return []

    result = []
    i = 0

    while i < len(s):
        # Find the delimiter '#'
        j = i
        while j < len(s) and s[j] != '#':
            j += 1

        # Extract the length
        length = int(s[i:j])

        # Move past the '#' delimiter
        j += 1

        # Extract the string of the specified length
        result.append(s[j:j + length])

        # Move index to the next encoded string
        i = j + length

    return result
