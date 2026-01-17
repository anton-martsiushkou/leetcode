package median_finder

import "container/heap"

// MaxHeap implements a max heap using Go's heap interface
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max heap comparison
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MinHeap implements a min heap using Go's heap interface
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min heap comparison
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MedianFinder maintains a data structure to find median efficiently
type MedianFinder struct {
	maxHeap *MaxHeap // Stores smaller half of numbers
	minHeap *MinHeap // Stores larger half of numbers
}

// Constructor initializes the MedianFinder
func Constructor() MedianFinder {
	maxH := &MaxHeap{}
	minH := &MinHeap{}
	heap.Init(maxH)
	heap.Init(minH)
	return MedianFinder{
		maxHeap: maxH,
		minHeap: minH,
	}
}

// AddNum adds a number to the data structure
// Time: O(log n)
func (mf *MedianFinder) AddNum(num int) {
	// Always add to max heap first
	heap.Push(mf.maxHeap, num)

	// Balance: move largest from max heap to min heap
	if mf.maxHeap.Len() > 0 && mf.minHeap.Len() > 0 && (*mf.maxHeap)[0] > (*mf.minHeap)[0] {
		val := heap.Pop(mf.maxHeap).(int)
		heap.Push(mf.minHeap, val)
	}

	// Maintain size property: max heap has equal or one more element
	if mf.minHeap.Len() > mf.maxHeap.Len() {
		val := heap.Pop(mf.minHeap).(int)
		heap.Push(mf.maxHeap, val)
	}
}

// FindMedian returns the median of all elements
// Time: O(1)
func (mf *MedianFinder) FindMedian() float64 {
	if mf.maxHeap.Len() > mf.minHeap.Len() {
		return float64((*mf.maxHeap)[0])
	}
	return float64((*mf.maxHeap)[0]+(*mf.minHeap)[0]) / 2.0
}
