package combinationsumiv

import "sort"

// https://leetcode.com/problems/combination-sum-iv

// dp[i][k]: choose i numbers, how many combinations with sum k.
// dp[i][k] = sum { dp[i-1][k-nums[j]] }, 0 <= j < n
//
// simplify
// dp[i][k]: choose at most i numbers, how many combinations with sum k.
// dp[i][k] = dp[i-1]k] + sum { dp[i-1][k-nums[j]] }, 0 <= j < n
// Now, we can ignore the first dimension i.
//
// simplify more
// dp[k]: number of combinations with sum k
// dp[k] = sum { dp[k-nums[i]] }
func combinationSum4(nums []int, target int) int {

	newNums := []int{}
	for _, v := range nums {
		if v > target {
			continue
		}
		newNums = append(newNums, v)
	}
	if len(newNums) == 0 {
		return 0
	}
	sort.Ints(newNums)
	nums = newNums

	dp := make([]int, target+1)
	dp[0] = 1
	for k := 1; k <= target; k++ {
		for _, v := range nums {
			if k < v {
				break
			}
			dp[k] += dp[k-v]
		}
	}
	return dp[target]
}

// This DP algorithm is incorrect because a number can be chosen multiple times.
// dp[i][k]: how many combinations in a[:(i+1)] with sum as k.
// dp[0][*] = 0; dp[0][nums[0]]=1
// dp[i][k] = sum{ dp[j][ k-nums[i] ] }, n < i
func combinationSum4_incrrect(nums []int, target int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}
	if nums[0] <= target {
		dp[0][nums[0]] = 1
	}
	result := dp[0][target]

	for i := 1; i < n; i++ {
		for k := 0; k <= target; k++ {
			for j := 0; j < i; j++ {
				if k < nums[i] {
					continue
				}
				dp[i][k] += dp[j][k-nums[i]]
			}
		}
		result += dp[i][target]
	}

	return result
}
