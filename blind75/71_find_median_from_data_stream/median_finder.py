import heapq


class MedianFinder:
    """
    Maintains a data structure to find median efficiently using two heaps.
    Time: O(log n) for addNum, O(1) for findMedian
    Space: O(n)
    """

    def __init__(self):
        """Initializes the MedianFinder."""
        self.max_heap = []  # Stores smaller half (negated for max heap)
        self.min_heap = []  # Stores larger half

    def addNum(self, num: int) -> None:
        """
        Adds a number to the data structure.
        Time: O(log n)

        Args:
            num: Integer to add
        """
        # Always add to max heap first (negate for max heap behavior)
        heapq.heappush(self.max_heap, -num)

        # Balance: move largest from max heap to min heap
        if self.max_heap and self.min_heap and (-self.max_heap[0] > self.min_heap[0]):
            val = -heapq.heappop(self.max_heap)
            heapq.heappush(self.min_heap, val)

        # Maintain size property: max heap has equal or one more element
        if len(self.min_heap) > len(self.max_heap):
            val = heapq.heappop(self.min_heap)
            heapq.heappush(self.max_heap, -val)

    def findMedian(self) -> float:
        """
        Returns the median of all elements.
        Time: O(1)

        Returns:
            Median as a float
        """
        if len(self.max_heap) > len(self.min_heap):
            return float(-self.max_heap[0])
        return (-self.max_heap[0] + self.min_heap[0]) / 2.0
