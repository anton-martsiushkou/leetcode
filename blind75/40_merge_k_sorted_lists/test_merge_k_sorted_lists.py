import pytest
from merge_k_sorted_lists import ListNode, merge_k_lists, merge_k_lists_divide_conquer


def create_list(vals: list[int]) -> ListNode | None:
    """Helper function to create a linked list from a list of values."""
    if not vals:
        return None

    head = ListNode(vals[0])
    curr = head
    for val in vals[1:]:
        curr.next = ListNode(val)
        curr = curr.next
    return head


def list_to_array(head: ListNode | None) -> list[int]:
    """Helper function to convert a linked list to a list."""
    result = []
    while head is not None:
        result.append(head.val)
        head = head.next
    return result


def test_example_1():
    """Test case from example 1."""
    lists = [create_list([1, 4, 5]), create_list([1, 3, 4]), create_list([2, 6])]
    result = merge_k_lists(lists)
    assert list_to_array(result) == [1, 1, 2, 3, 4, 4, 5, 6]


def test_example_2():
    """Test case from example 2 - empty array."""
    lists = []
    result = merge_k_lists(lists)
    assert list_to_array(result) == []


def test_example_3():
    """Test case from example 3 - array with empty list."""
    lists = [create_list([])]
    result = merge_k_lists(lists)
    assert list_to_array(result) == []


def test_single_list():
    """Test with a single list."""
    lists = [create_list([1, 2, 3])]
    result = merge_k_lists(lists)
    assert list_to_array(result) == [1, 2, 3]


def test_two_lists():
    """Test with two lists."""
    lists = [create_list([1, 3]), create_list([2, 4])]
    result = merge_k_lists(lists)
    assert list_to_array(result) == [1, 2, 3, 4]


def test_different_lengths():
    """Test with lists of different lengths."""
    lists = [create_list([1]), create_list([1, 3, 4]), create_list([2, 6])]
    result = merge_k_lists(lists)
    assert list_to_array(result) == [1, 1, 2, 3, 4, 6]


def test_some_empty_lists():
    """Test with some empty lists."""
    lists = [create_list([]), create_list([1]), create_list([])]
    result = merge_k_lists(lists)
    assert list_to_array(result) == [1]


def test_divide_conquer_example_1():
    """Test divide and conquer version with example 1."""
    lists = [create_list([1, 4, 5]), create_list([1, 3, 4]), create_list([2, 6])]
    result = merge_k_lists_divide_conquer(lists)
    assert list_to_array(result) == [1, 1, 2, 3, 4, 4, 5, 6]


def test_divide_conquer_empty_array():
    """Test divide and conquer version with empty array."""
    lists = []
    result = merge_k_lists_divide_conquer(lists)
    assert list_to_array(result) == []


def test_divide_conquer_single_list():
    """Test divide and conquer version with single list."""
    lists = [create_list([1, 2, 3])]
    result = merge_k_lists_divide_conquer(lists)
    assert list_to_array(result) == [1, 2, 3]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
