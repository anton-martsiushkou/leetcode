from typing import Optional


class ListNode:
    """Definition for singly-linked list node."""

    def __init__(self, val: int = 0, next: Optional['ListNode'] = None):
        self.val = val
        self.next = next


def reverse_list(head: Optional[ListNode]) -> Optional[ListNode]:
    """
    Reverses a singly linked list using an iterative approach.
    Time complexity: O(n), Space complexity: O(1)

    Args:
        head: The head of the linked list

    Returns:
        The head of the reversed linked list
    """
    prev = None
    curr = head

    while curr is not None:
        next_node = curr.next
        curr.next = prev
        prev = curr
        curr = next_node

    return prev


def reverse_list_recursive(head: Optional[ListNode]) -> Optional[ListNode]:
    """
    Reverses a singly linked list using a recursive approach.
    Time complexity: O(n), Space complexity: O(n) due to call stack

    Args:
        head: The head of the linked list

    Returns:
        The head of the reversed linked list
    """
    if head is None or head.next is None:
        return head

    new_head = reverse_list_recursive(head.next)
    head.next.next = head
    head.next = None

    return new_head
