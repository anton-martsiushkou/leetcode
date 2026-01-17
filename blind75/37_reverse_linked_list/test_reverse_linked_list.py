import pytest
from reverse_linked_list import ListNode, reverse_list, reverse_list_recursive


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
    head = create_list([1, 2, 3, 4, 5])
    result = reverse_list(head)
    assert list_to_array(result) == [5, 4, 3, 2, 1]


def test_example_2():
    """Test case from example 2."""
    head = create_list([1, 2])
    result = reverse_list(head)
    assert list_to_array(result) == [2, 1]


def test_example_3():
    """Test case from example 3 - empty list."""
    head = create_list([])
    result = reverse_list(head)
    assert list_to_array(result) == []


def test_single_node():
    """Test with a single node."""
    head = create_list([1])
    result = reverse_list(head)
    assert list_to_array(result) == [1]


def test_three_nodes():
    """Test with three nodes."""
    head = create_list([1, 2, 3])
    result = reverse_list(head)
    assert list_to_array(result) == [3, 2, 1]


def test_recursive_example_1():
    """Test recursive version with example 1."""
    head = create_list([1, 2, 3, 4, 5])
    result = reverse_list_recursive(head)
    assert list_to_array(result) == [5, 4, 3, 2, 1]


def test_recursive_example_2():
    """Test recursive version with example 2."""
    head = create_list([1, 2])
    result = reverse_list_recursive(head)
    assert list_to_array(result) == [2, 1]


def test_recursive_empty_list():
    """Test recursive version with empty list."""
    head = create_list([])
    result = reverse_list_recursive(head)
    assert list_to_array(result) == []


def test_recursive_single_node():
    """Test recursive version with a single node."""
    head = create_list([1])
    result = reverse_list_recursive(head)
    assert list_to_array(result) == [1]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
