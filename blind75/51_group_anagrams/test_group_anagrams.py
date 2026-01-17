import pytest
from group_anagrams import group_anagrams


def normalize_result(result):
    """Normalize result for comparison by sorting."""
    return sorted([sorted(group) for group in result])


def test_example_1():
    """Test case from example 1."""
    strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
    result = group_anagrams(strs)
    expected = [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
    assert normalize_result(result) == normalize_result(expected)


def test_example_2():
    """Test case from example 2."""
    strs = [""]
    result = group_anagrams(strs)
    expected = [[""]]
    assert normalize_result(result) == normalize_result(expected)


def test_example_3():
    """Test case from example 3."""
    strs = ["a"]
    result = group_anagrams(strs)
    expected = [["a"]]
    assert normalize_result(result) == normalize_result(expected)


def test_all_same():
    """Test with all anagrams of same word."""
    strs = ["abc", "bca", "cab"]
    result = group_anagrams(strs)
    expected = [["abc", "bca", "cab"]]
    assert normalize_result(result) == normalize_result(expected)


def test_all_different():
    """Test with all different words."""
    strs = ["a", "b", "c"]
    result = group_anagrams(strs)
    expected = [["a"], ["b"], ["c"]]
    assert normalize_result(result) == normalize_result(expected)


def test_with_empty_string():
    """Test with empty strings."""
    strs = ["", "b", ""]
    result = group_anagrams(strs)
    expected = [["", ""], ["b"]]
    assert normalize_result(result) == normalize_result(expected)


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
