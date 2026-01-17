package meeting_rooms_ii

import (
	"container/heap"
	"sort"
)

// MinHeap implements a min heap for end times
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
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

// MinMeetingRooms returns the minimum number of conference rooms required.
// Time: O(n log n), Space: O(n)
func MinMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Min heap to track end times of ongoing meetings
	h := &MinHeap{}
	heap.Init(h)

	maxRooms := 0

	for _, interval := range intervals {
		// Remove meetings that have ended
		for h.Len() > 0 && (*h)[0] <= interval[0] {
			heap.Pop(h)
		}

		// Add current meeting's end time
		heap.Push(h, interval[1])

		// Track maximum concurrent meetings
		if h.Len() > maxRooms {
			maxRooms = h.Len()
		}
	}

	return maxRooms
}

// MinMeetingRoomsChronological uses chronological ordering approach.
// Time: O(n log n), Space: O(n)
func MinMeetingRoomsChronological(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	n := len(intervals)
	starts := make([]int, n)
	ends := make([]int, n)

	for i, interval := range intervals {
		starts[i] = interval[0]
		ends[i] = interval[1]
	}

	sort.Ints(starts)
	sort.Ints(ends)

	rooms := 0
	maxRooms := 0
	startPtr := 0
	endPtr := 0

	for startPtr < n {
		if starts[startPtr] < ends[endPtr] {
			// Meeting starts, need a room
			rooms++
			startPtr++
			if rooms > maxRooms {
				maxRooms = rooms
			}
		} else {
			// Meeting ends, free a room
			rooms--
			endPtr++
		}
	}

	return maxRooms
}
