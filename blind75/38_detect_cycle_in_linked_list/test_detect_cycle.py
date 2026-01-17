import pytest
from detect_cycle import ListNode, has_cycle, has_cycle_hash_set


def create_list_with_cycle(vals: list[int], pos: int) -> ListNode | None:
    """Helper function to create a linked list with a cycle."""
    if not vals:
        return None

    head = ListNode(vals[0])
    curr = head
    cycle_node = None

    if pos == 0:
        cycle_node = head

    for i in range(1, len(vals)):
        curr.next = ListNode(vals[i])
        curr = curr.next
        if i == pos:
            cycle_node = curr

    # Create the cycle
    if pos >= 0:
        curr.next = cycle_node

    return head


def test_example_1():
    """Test case from example 1 - cycle at position 1."""
    head = create_list_with_cycle([3, 2, 0, -4], 1)
    assert has_cycle(head) is True


def test_example_2():
    """Test case from example 2 - cycle at position 0."""
    head = create_list_with_cycle([1, 2], 0)
    assert has_cycle(head) is True


def test_example_3():
    """Test case from example 3 - no cycle."""
    head = create_list_with_cycle([1], -1)
    assert has_cycle(head) is False


def test_empty_list():
    """Test with empty list."""
    head = create_list_with_cycle([], -1)
    assert has_cycle(head) is False


def test_no_cycle_multiple_nodes():
    """Test with multiple nodes and no cycle."""
    head = create_list_with_cycle([1, 2, 3, 4, 5], -1)
    assert has_cycle(head) is False


def test_cycle_at_end():
    """Test with cycle at the end."""
    head = create_list_with_cycle([1, 2, 3], 2)
    assert has_cycle(head) is True


def test_hash_set_example_1():
    """Test hash set version with example 1."""
    head = create_list_with_cycle([3, 2, 0, -4], 1)
    assert has_cycle_hash_set(head) is True


def test_hash_set_example_2():
    """Test hash set version with example 2."""
    head = create_list_with_cycle([1, 2], 0)
    assert has_cycle_hash_set(head) is True


def test_hash_set_no_cycle():
    """Test hash set version with no cycle."""
    head = create_list_with_cycle([1], -1)
    assert has_cycle_hash_set(head) is False


def test_hash_set_empty_list():
    """Test hash set version with empty list."""
    head = create_list_with_cycle([], -1)
    assert has_cycle_hash_set(head) is False


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
