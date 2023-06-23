package h_index

/*
Observations.
1. The h-index is not necessarily one of the citation number. e.g. [1, 1, 10, 10, 10]. the h-index is 3.
2. If there are x citations >= x, then the h-index >= x. This means, the h-index is monotone. We can binary search.
3. The input size is small, we can iterate it all possible h-index without binary search.
*/

func hIndex(citations []int) int {
	cnt := make(map[int]int)
	n := len(citations)
	min, max := citations[0], citations[0]
	for _, v := range citations {
		cnt[v]++
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	for i := max - 1; i >= 0; i-- {
		cnt[i] += cnt[i+1]
	}
	for i := n; i > 0; i-- {
		if cnt[i] >= i {
			return i
		}
	}
	return 0
}
