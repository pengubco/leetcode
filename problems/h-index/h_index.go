package h_index

import "sort"

/*
1. Given x, how many papers have citation no less than x?
2. Then we can binary search the x.
*/

func hIndex(citations []int) int {
	cnt := make(map[int]int)
	sort.Ints(citations)
	n := len(citations)
	for i := n - 1; i >= 0; i-- {
		if i+1 < n && citations[i] != citations[i+1] {
			cnt[citations[i]] = cnt[citations[i+1]]
		}
		cnt[citations[i]]++
	}
	result := 0
	// The answer range is small enough, no need to binary search.
	for k, v := range cnt {
		if v >= k && k > result {
			result = k
		}
	}
	return result
}
