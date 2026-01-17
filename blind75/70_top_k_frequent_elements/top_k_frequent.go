package top_k_frequent

// TopKFrequent returns the k most frequent elements using bucket sort.
// Time: O(n), Space: O(n)
func TopKFrequent(nums []int, k int) []int {
	// Count frequencies
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Create buckets where index is frequency
	buckets := make([][]int, len(nums)+1)
	for num, count := range freq {
		buckets[count] = append(buckets[count], num)
	}

	// Collect k most frequent elements from highest frequency buckets
	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		result = append(result, buckets[i]...)
	}

	return result[:k]
}
