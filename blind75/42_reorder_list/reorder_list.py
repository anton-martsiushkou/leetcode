from typing import Optional


class ListNode:
    """Definition for singly-linked list node."""

    def __init__(self, val: int = 0, next: Optional['ListNode'] = None):
        self.val = val
        self.next = next


def reorder_list(head: Optional[ListNode]) -> None:
    """
    Reorders the list in-place using find middle + reverse + merge.
    Time complexity: O(n), Space complexity: O(1)

    Args:
        head: The head of the linked list

    Note:
        Modifies the list in-place. Does not return anything.
    """
    if head is None or head.next is None:
        return

    # Step 1: Find the middle of the list
    slow, fast = head, head
    while fast.next is not None and fast.next.next is not None:
        slow = slow.next
        fast = fast.next.next

    # Step 2: Split and reverse the second half
    second = slow.next
    slow.next = None  # Split the list
    second = reverse_list(second)

    # Step 3: Merge the two halves
    first = head
    while second is not None:
        tmp1 = first.next
        tmp2 = second.next

        first.next = second
        second.next = tmp1

        first = tmp1
        second = tmp2


def reverse_list(head: Optional[ListNode]) -> Optional[ListNode]:
    """Helper function to reverse a linked list."""
    prev = None
    curr = head

    while curr is not None:
        next_node = curr.next
        curr.next = prev
        prev = curr
        curr = next_node

    return prev


def reorder_list_stack(head: Optional[ListNode]) -> None:
    """
    Reorders the list using a stack.
    Time complexity: O(n), Space complexity: O(n)

    Args:
        head: The head of the linked list

    Note:
        Modifies the list in-place. Does not return anything.
    """
    if head is None or head.next is None:
        return

    # Push all nodes to stack
    stack = []
    curr = head
    while curr is not None:
        stack.append(curr)
        curr = curr.next

    # Reorder by taking from front and back alternately
    curr = head
    n = len(stack)
    for i in range(n // 2):
        next_node = curr.next
        back = stack.pop()

        curr.next = back
        back.next = next_node

        curr = next_node

    # Set the end of the list
    if curr is not None:
        curr.next = None
