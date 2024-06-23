package longestincreasingsubsequence

import (
	"math"
	"sort"
)

// f[i] = k. the LIS ends at nums[i] is k. meaning the last element of the LIS is nums[i].
// f[i] = max{f[j]+1, nums[j] < nums[i]}
// The answer is max{f[i]}. Time complexity is O(n^2).
//
// Let g[k] = x denote that, x is the minimum last element of any LIS of length k.
// Then f[i] = max{k+1, g[k] < nums[i]}
// g[k] is strictly increasing. so we can find max{k+1} using binary search.
// Actually, there is no need to keep f[i]. The answer is the largest k s.t., g[k] is set.
//

func lengthOfLIS(nums []int) int {
	n := len(nums)
	g := make([]int, n+1)
	g[0] = math.MinInt
	g[1] = nums[0]
	maxLIS := 1 // g[0..maxLIS] is set.
	for i := 1; i < n; i++ {
		target := nums[i]
		j := sort.Search(maxLIS+1, func(i int) bool {
			return g[i] >= target
		})
		// find a longer LIS ends at nums[i]
		if j == maxLIS+1 && g[maxLIS] < nums[i] {
			maxLIS++
			g[maxLIS] = nums[i]
			continue
		}
		/// find a LIS of the same length but with smaller last elemebnt.
		if g[j] > nums[i] {
			g[j] = nums[i]
		}
	}
	return maxLIS
}
