from typing import List
from collections import defaultdict


def group_anagrams(strs: List[str]) -> List[List[str]]:
    """
    Groups anagrams together from the input array.
    Uses hash map with sorted string as key for O(n*k*log(k)) time complexity.

    Args:
        strs: List of strings to group

    Returns:
        List of lists, where each inner list contains anagrams
    """
    groups = defaultdict(list)

    for s in strs:
        key = "".join(sorted(s))
        groups[key].append(s)

    return list(groups.values())
