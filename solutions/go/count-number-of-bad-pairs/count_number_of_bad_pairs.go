package count_number_of_bad_pairs

// https://leetcode.com/problems/count-number-of-bad-pairs

func countBadPairs(nums []int) int64 {
	n := len(nums)
	buckets := make(map[int]int64)
	for i, v := range nums {
		buckets[v-i]++
	}
	var validPairs int64
	for _, m := range buckets {
		validPairs += m * (m - 1) / 2
	}
	return int64(n)*int64(n-1)/2 - validPairs
}
