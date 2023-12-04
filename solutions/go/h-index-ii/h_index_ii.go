package h_index_ii

// https://leetcode.com/problems/h-index-ii/

/*
Same idea with the "H Index". However, the size is bigger and we need to use binary search.
*/
func hIndex(citations []int) int {
	cnt := make(map[int]int)
	max := citations[0]
	for _, v := range citations {
		cnt[v]++
		if v > max {
			max = v
		}
	}
	for i := max - 1; i > 0; i-- {
		cnt[i] += cnt[i+1]
	}
	result := binarySearch(0, max, func(i int) bool {
		return cnt[i] >= i
	})
	if result == -1 {
		return 0
	}
	return result
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
