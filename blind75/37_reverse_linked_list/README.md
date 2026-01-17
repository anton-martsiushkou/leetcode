# Reverse Linked List

## Problem Description

Given the head of a singly linked list, reverse the list, and return the reversed list.

## Examples

**Example 1:**
```
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
```

**Example 2:**
```
Input: head = [1,2]
Output: [2,1]
```

**Example 3:**
```
Input: head = []
Output: []
```

## Constraints

- The number of nodes in the list is in the range `[0, 5000]`
- `-5000 <= Node.val <= 5000`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Iterative (Three Pointers) | Linked List | O(n) | O(1) |
| Recursive | Linked List | O(n) | O(n) |

### Approach 1: Iterative (Recommended)

The iterative solution uses three pointers to reverse the list in place. We traverse the list once, reversing the direction of each pointer as we go.

**Key Insight:** To reverse a linked list, we need to change the `next` pointer of each node to point to its previous node instead of its next node. However, we need to be careful not to lose the reference to the rest of the list.

### Algorithm Steps

1. Initialize three pointers:
   - `prev = nil` (this will be the new tail of the reversed list)
   - `curr = head` (current node being processed)
   - `next = nil` (temporary storage for the next node)
2. While `curr` is not nil:
   - Save the next node: `next = curr.next`
   - Reverse the current node's pointer: `curr.next = prev`
   - Move `prev` and `curr` one step forward:
     - `prev = curr`
     - `curr = next`
3. Return `prev` (which is now the new head of the reversed list)

### Example Walkthrough

For `head = [1, 2, 3, 4, 5]`:

1. **Initial**: prev=nil, curr=1, next=nil
   - List: 1->2->3->4->5

2. **Iteration 1**:
   - next = 2
   - 1.next = nil
   - prev = 1, curr = 2
   - List: nil<-1  2->3->4->5

3. **Iteration 2**:
   - next = 3
   - 2.next = 1
   - prev = 2, curr = 3
   - List: nil<-1<-2  3->4->5

4. **Iteration 3**:
   - next = 4
   - 3.next = 2
   - prev = 3, curr = 4
   - List: nil<-1<-2<-3  4->5

5. **Iteration 4**:
   - next = 5
   - 4.next = 3
   - prev = 4, curr = 5
   - List: nil<-1<-2<-3<-4  5

6. **Iteration 5**:
   - next = nil
   - 5.next = 4
   - prev = 5, curr = nil
   - List: nil<-1<-2<-3<-4<-5

7. **Return**: prev = 5 (new head)

### Why This is Optimal

- **Time Complexity O(n)**: We traverse each node exactly once
- **Space Complexity O(1)**: We only use three pointers regardless of list size
- This is optimal because we must visit every node at least once to reverse the list

### Approach 2: Recursive

The recursive approach reverses the list by recursively reversing the rest of the list and then fixing the pointers. While elegant, it uses O(n) space on the call stack.

**Algorithm:**
1. Base case: If head is nil or head.next is nil, return head
2. Recursively reverse the rest of the list: `newHead = reverseList(head.next)`
3. Fix the pointers: `head.next.next = head` and `head.next = nil`
4. Return the new head

The recursive solution is less preferred for large lists due to stack overflow risk.
