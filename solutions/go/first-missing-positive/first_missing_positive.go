package firstmissingpositive

// https://leetcode.com/problems/first-missing-positive/

// Approach 1. Binary search [1, n]. if count(x) < x, then there must be a positive integer in [1:x] missing.
// The time complexity is O(n*logn), which is essentially O(n) because logn < log(10^5)
//
// Approach 2. Place nums[i] to nums[nums[i]-1] if nums[i] is in [1:n]. Then iterate to find the first x s.t.
// nums[x-1] != x. The place would always place one number at its correct place, so in total there are at most n placements.
// Note that, there would't be a cycle here, because if we swap nums[i], nums[j], then nums[j] must not be at its proper
// position because nums[j] is at nums[i]'s position.
// What if nums[i] == nums[j]? In that case, there is a cycle. However, for our case, we can stop placing because
// duplicate won't impact the result for this problem.

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] != i+1 {
			if nums[i] <= 0 || nums[i] > n {
				break
			}
			if nums[nums[i]-1] == nums[i] {
				break
			}
			// swap nums[i] with nums[nums[i]-1] if they are different.
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return 0
}
