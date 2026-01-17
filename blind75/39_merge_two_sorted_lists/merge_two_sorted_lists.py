from typing import Optional


class ListNode:
    """Definition for singly-linked list node."""

    def __init__(self, val: int = 0, next: Optional['ListNode'] = None):
        self.val = val
        self.next = next


def merge_two_lists(list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
    """
    Merges two sorted linked lists iteratively.
    Time complexity: O(n + m), Space complexity: O(1)

    Args:
        list1: Head of the first sorted linked list
        list2: Head of the second sorted linked list

    Returns:
        Head of the merged sorted linked list
    """
    dummy = ListNode()
    tail = dummy

    while list1 is not None and list2 is not None:
        if list1.val <= list2.val:
            tail.next = list1
            list1 = list1.next
        else:
            tail.next = list2
            list2 = list2.next
        tail = tail.next

    # Attach the remaining nodes
    if list1 is not None:
        tail.next = list1
    else:
        tail.next = list2

    return dummy.next


def merge_two_lists_recursive(list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
    """
    Merges two sorted linked lists recursively.
    Time complexity: O(n + m), Space complexity: O(n + m) due to call stack

    Args:
        list1: Head of the first sorted linked list
        list2: Head of the second sorted linked list

    Returns:
        Head of the merged sorted linked list
    """
    if list1 is None:
        return list2
    if list2 is None:
        return list1

    if list1.val <= list2.val:
        list1.next = merge_two_lists_recursive(list1.next, list2)
        return list1

    list2.next = merge_two_lists_recursive(list1, list2.next)
    return list2
