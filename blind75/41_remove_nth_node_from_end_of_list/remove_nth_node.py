from typing import Optional


class ListNode:
    """Definition for singly-linked list node."""

    def __init__(self, val: int = 0, next: Optional['ListNode'] = None):
        self.val = val
        self.next = next


def remove_nth_from_end(head: Optional[ListNode], n: int) -> Optional[ListNode]:
    """
    Removes the nth node from the end using two pointers (one pass).
    Time complexity: O(n), Space complexity: O(1)

    Args:
        head: The head of the linked list
        n: The position from the end (1-indexed)

    Returns:
        The head of the modified linked list
    """
    dummy = ListNode(0, head)
    fast = dummy
    slow = dummy

    # Move fast n+1 steps ahead
    for _ in range(n + 1):
        fast = fast.next

    # Move both pointers until fast reaches the end
    while fast is not None:
        fast = fast.next
        slow = slow.next

    # Remove the nth node
    slow.next = slow.next.next

    return dummy.next


def remove_nth_from_end_two_pass(head: Optional[ListNode], n: int) -> Optional[ListNode]:
    """
    Removes the nth node from the end using two passes.
    Time complexity: O(n), Space complexity: O(1)

    Args:
        head: The head of the linked list
        n: The position from the end (1-indexed)

    Returns:
        The head of the modified linked list
    """
    # First pass: count the length
    length = 0
    curr = head
    while curr is not None:
        length += 1
        curr = curr.next

    # Edge case: remove the first node
    if n == length:
        return head.next

    # Second pass: move to the node before the one to remove
    dummy = ListNode(0, head)
    curr = dummy
    for _ in range(length - n):
        curr = curr.next

    # Remove the node
    curr.next = curr.next.next

    return dummy.next
