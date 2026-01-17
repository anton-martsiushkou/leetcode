from typing import Optional
import heapq


class ListNode:
    """Definition for singly-linked list node."""

    def __init__(self, val: int = 0, next: Optional['ListNode'] = None):
        self.val = val
        self.next = next


def merge_k_lists(lists: list[Optional[ListNode]]) -> Optional[ListNode]:
    """
    Merges k sorted linked lists using a min heap.
    Time complexity: O(N log k), Space complexity: O(k)
    where N is the total number of nodes and k is the number of lists

    Args:
        lists: List of heads of sorted linked lists

    Returns:
        Head of the merged sorted linked list
    """
    if not lists:
        return None

    # Min heap: (value, index, node)
    # We use index to break ties and ensure stable ordering
    min_heap = []

    # Initialize heap with first node from each list
    for i, node in enumerate(lists):
        if node is not None:
            heapq.heappush(min_heap, (node.val, i, node))

    dummy = ListNode()
    tail = dummy

    # Build merged list
    next_index = len(lists)  # For nodes added later
    while min_heap:
        val, idx, node = heapq.heappop(min_heap)
        tail.next = node
        tail = tail.next

        if node.next is not None:
            heapq.heappush(min_heap, (node.next.val, next_index, node.next))
            next_index += 1

    return dummy.next


def merge_k_lists_divide_conquer(lists: list[Optional[ListNode]]) -> Optional[ListNode]:
    """
    Merges k sorted lists using divide and conquer.
    Time complexity: O(N log k), Space complexity: O(log k) for recursion

    Args:
        lists: List of heads of sorted linked lists

    Returns:
        Head of the merged sorted linked list
    """
    if not lists:
        return None

    def merge_helper(left: int, right: int) -> Optional[ListNode]:
        if left == right:
            return lists[left]
        if left > right:
            return None

        mid = left + (right - left) // 2
        l1 = merge_helper(left, mid)
        l2 = merge_helper(mid + 1, right)

        return merge_two_lists(l1, l2)

    return merge_helper(0, len(lists) - 1)


def merge_two_lists(list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
    """Helper function to merge two sorted linked lists."""
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

    if list1 is not None:
        tail.next = list1
    else:
        tail.next = list2

    return dummy.next
