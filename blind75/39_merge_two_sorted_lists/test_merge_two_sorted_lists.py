import pytest
from merge_two_sorted_lists import ListNode, merge_two_lists, merge_two_lists_recursive


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
    list1 = create_list([1, 2, 4])
    list2 = create_list([1, 3, 4])
    result = merge_two_lists(list1, list2)
    assert list_to_array(result) == [1, 1, 2, 3, 4, 4]


def test_example_2():
    """Test case from example 2 - both empty."""
    list1 = create_list([])
    list2 = create_list([])
    result = merge_two_lists(list1, list2)
    assert list_to_array(result) == []


def test_example_3():
    """Test case from example 3 - first empty."""
    list1 = create_list([])
    list2 = create_list([0])
    result = merge_two_lists(list1, list2)
    assert list_to_array(result) == [0]


def test_second_empty():
    """Test with second list empty."""
    list1 = create_list([1, 2, 3])
    list2 = create_list([])
    result = merge_two_lists(list1, list2)
    assert list_to_array(result) == [1, 2, 3]


def test_no_overlap():
    """Test with no overlap between lists."""
    list1 = create_list([1, 2, 3])
    list2 = create_list([4, 5, 6])
    result = merge_two_lists(list1, list2)
    assert list_to_array(result) == [1, 2, 3, 4, 5, 6]


def test_interleaved():
    """Test with interleaved values."""
    list1 = create_list([1, 3, 5])
    list2 = create_list([2, 4, 6])
    result = merge_two_lists(list1, list2)
    assert list_to_array(result) == [1, 2, 3, 4, 5, 6]


def test_recursive_example_1():
    """Test recursive version with example 1."""
    list1 = create_list([1, 2, 4])
    list2 = create_list([1, 3, 4])
    result = merge_two_lists_recursive(list1, list2)
    assert list_to_array(result) == [1, 1, 2, 3, 4, 4]


def test_recursive_both_empty():
    """Test recursive version with both lists empty."""
    list1 = create_list([])
    list2 = create_list([])
    result = merge_two_lists_recursive(list1, list2)
    assert list_to_array(result) == []


def test_recursive_first_empty():
    """Test recursive version with first list empty."""
    list1 = create_list([])
    list2 = create_list([0])
    result = merge_two_lists_recursive(list1, list2)
    assert list_to_array(result) == [0]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
