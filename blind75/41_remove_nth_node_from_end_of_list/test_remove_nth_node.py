import pytest
from remove_nth_node import ListNode, remove_nth_from_end, remove_nth_from_end_two_pass


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
    result = remove_nth_from_end(head, 2)
    assert list_to_array(result) == [1, 2, 3, 5]


def test_example_2():
    """Test case from example 2 - single node."""
    head = create_list([1])
    result = remove_nth_from_end(head, 1)
    assert list_to_array(result) == []


def test_example_3():
    """Test case from example 3 - remove last."""
    head = create_list([1, 2])
    result = remove_nth_from_end(head, 1)
    assert list_to_array(result) == [1]


def test_remove_first():
    """Test removing the first node."""
    head = create_list([1, 2])
    result = remove_nth_from_end(head, 2)
    assert list_to_array(result) == [2]


def test_remove_middle():
    """Test removing a middle node."""
    head = create_list([1, 2, 3])
    result = remove_nth_from_end(head, 2)
    assert list_to_array(result) == [1, 3]


def test_remove_from_longer_list():
    """Test removing from a longer list."""
    head = create_list([1, 2, 3, 4, 5, 6])
    result = remove_nth_from_end(head, 3)
    assert list_to_array(result) == [1, 2, 3, 5, 6]


def test_two_pass_example_1():
    """Test two-pass version with example 1."""
    head = create_list([1, 2, 3, 4, 5])
    result = remove_nth_from_end_two_pass(head, 2)
    assert list_to_array(result) == [1, 2, 3, 5]


def test_two_pass_single_node():
    """Test two-pass version with single node."""
    head = create_list([1])
    result = remove_nth_from_end_two_pass(head, 1)
    assert list_to_array(result) == []


def test_two_pass_remove_last():
    """Test two-pass version removing last node."""
    head = create_list([1, 2])
    result = remove_nth_from_end_two_pass(head, 1)
    assert list_to_array(result) == [1]


def test_two_pass_remove_first():
    """Test two-pass version removing first node."""
    head = create_list([1, 2])
    result = remove_nth_from_end_two_pass(head, 2)
    assert list_to_array(result) == [2]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
