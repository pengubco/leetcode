package h_index_ii

import (
	"slices"
	"sort"
)

// https://leetcode.com/problems/h-index-ii/

/*
Same idea with the "H Index". However, the size is bigger and we need to use binary search.
*/
func hIndex(citations []int) int {
	maxCitation := slices.Max(citations)
	result := binarySearch(0, maxCitation, func(i int) bool {
		return count(citations, i) >= i
	})
	if result == -1 {
		return 0
	}
	return result
}

// a is non-decreasing. return number of elements larger than or equal to the given target.
func count(a []int, target int) int {
	return len(a) - sort.Search(len(a), func(i int) bool {
		return a[i] >= target
	})
}

// With condition "if f(l)=true, then f(l-1)=true" "if f(l) is false, then f(l+1) is false.
// find the largest value x such that  l <= x <=h and f(x) is true
func binarySearch(l, h int, f func(i int) bool) int {
	if l > h || !f(l) {
		return -1
	}
	for l < h-1 {
		mid := (l + h) / 2
		if f(mid) {
			l = mid
			continue
		} else {
			h = mid - 1
		}
	}
	i := h
	for ; i >= l; i-- {
		if f(i) {
			return i
		}
	}
	return i
}
