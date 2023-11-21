package maxconsecutiveonesiii

// 1004 https://leetcode.com/problems/max-consecutive-ones-iii/description/

/*
Number of flips needed for nums[i:j] to all 1s = sum[j]-sum[i-1].
Sliding window. Pick an i, find the largest j. Increase i, find the largest j.
Repeat.
*/
func longestOnes(nums []int, k int) int {
	n := len(nums)
	sum := make([]int, n)
	sum[0] = 1 - nums[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + (1 - nums[i])
	}
	countZero := func(l, r int) int {
		if l == 0 {
			return sum[r]
		}
		return sum[r] - sum[l-1]
	}

	j := 0
	result := 0
	for i := 0; i < n; i++ {
		if j >= n {
			break
		}
		for j < n && countZero(i, j) <= k {
			if j-i+1 > result {
				result = j - i + 1
			}
			j++
		}
	}
	return result
}
