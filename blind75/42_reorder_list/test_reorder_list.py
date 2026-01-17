import pytest
from reorder_list import ListNode, reorder_list, reorder_list_stack


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
    """Test case from example 1 - even length."""
    head = create_list([1, 2, 3, 4])
    reorder_list(head)
    assert list_to_array(head) == [1, 4, 2, 3]


def test_example_2():
    """Test case from example 2 - odd length."""
    head = create_list([1, 2, 3, 4, 5])
    reorder_list(head)
    assert list_to_array(head) == [1, 5, 2, 4, 3]


def test_single_node():
    """Test with a single node."""
    head = create_list([1])
    reorder_list(head)
    assert list_to_array(head) == [1]


def test_two_nodes():
    """Test with two nodes."""
    head = create_list([1, 2])
    reorder_list(head)
    assert list_to_array(head) == [1, 2]


def test_three_nodes():
    """Test with three nodes."""
    head = create_list([1, 2, 3])
    reorder_list(head)
    assert list_to_array(head) == [1, 3, 2]


def test_six_nodes():
    """Test with six nodes."""
    head = create_list([1, 2, 3, 4, 5, 6])
    reorder_list(head)
    assert list_to_array(head) == [1, 6, 2, 5, 3, 4]


def test_stack_example_1():
    """Test stack version with example 1."""
    head = create_list([1, 2, 3, 4])
    reorder_list_stack(head)
    assert list_to_array(head) == [1, 4, 2, 3]


def test_stack_example_2():
    """Test stack version with example 2."""
    head = create_list([1, 2, 3, 4, 5])
    reorder_list_stack(head)
    assert list_to_array(head) == [1, 5, 2, 4, 3]


def test_stack_single_node():
    """Test stack version with single node."""
    head = create_list([1])
    reorder_list_stack(head)
    assert list_to_array(head) == [1]


def test_stack_two_nodes():
    """Test stack version with two nodes."""
    head = create_list([1, 2])
    reorder_list_stack(head)
    assert list_to_array(head) == [1, 2]


def test_stack_three_nodes():
    """Test stack version with three nodes."""
    head = create_list([1, 2, 3])
    reorder_list_stack(head)
    assert list_to_array(head) == [1, 3, 2]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
