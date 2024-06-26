package minimumnumberofoperationstomakearraycontinuous

import (
	"slices"
)

// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous

func minOperations(nums []int) int {
	n := len(nums)
	slices.Sort(nums)
	// deduped nums
	a := []int{nums[0]}
	for i := 1; i < n; i++ {
		if nums[i] == a[len(a)-1] {
			continue
		}
		a = append(a, nums[i])
	}

	result := n
	j := 0
	m := len(a)
	for i := 0; i < m; i++ {
		if j >= m {
			break
		}
		for j < m && a[j]-a[i]+1 <= n {
			result = min(result, n-(j-i+1))
			j++
		}
	}
	return result
}
