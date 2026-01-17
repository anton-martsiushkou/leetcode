from typing import Optional


class ListNode:
    """Definition for singly-linked list node."""

    def __init__(self, val: int = 0, next: Optional['ListNode'] = None):
        self.val = val
        self.next = next


def has_cycle(head: Optional[ListNode]) -> bool:
    """
    Determines if a linked list has a cycle using Floyd's algorithm.
    Time complexity: O(n), Space complexity: O(1)

    Args:
        head: The head of the linked list

    Returns:
        True if the linked list has a cycle, False otherwise
    """
    if head is None or head.next is None:
        return False

    slow = head
    fast = head

    while fast is not None and fast.next is not None:
        slow = slow.next
        fast = fast.next.next

        if slow == fast:
            return True

    return False


def has_cycle_hash_set(head: Optional[ListNode]) -> bool:
    """
    Determines if a linked list has a cycle using a hash set.
    Time complexity: O(n), Space complexity: O(n)

    Args:
        head: The head of the linked list

    Returns:
        True if the linked list has a cycle, False otherwise
    """
    seen = set()

    curr = head
    while curr is not None:
        if curr in seen:
            return True
        seen.add(curr)
        curr = curr.next

    return False
