package minimumsizesubarrayininfinitearray

import "math"

/*
Any subarray of the infinite array is comprised of three parts.
1. A suffix of the original array.
2. A prefix of the original array.
3. K copies of the original array.
Each part can be empty.

We can further simplifies by combining part 1 and part 2 to the following.
A subarray of two copies of the original array.
a[i:j] + k * total_sum = target. `a` is the original array concatenated to itself.

a[i:j]%total_sum = target%total_sum
*/
func minSizeSubarray(nums []int, target int) int {
	n := len(nums)
	result := 0
	totalSum := 0
	for i := 0; i < n; i++ {
		totalSum += nums[i]
	}
	result += n * (target / totalSum)
	target = target % totalSum

	if target == 0 {
		return result
	}

	psumToIndex := make(map[int]int)
	sum := 0
	minSize := math.MaxInt
	for i := 0; i < 2*n; i++ {
		sum += nums[i%n]
		psumToIndex[sum] = i
		if sum < target {
			continue
		}
		if sum == target {
			if minSize > i+1 {
				minSize = i + 1
				continue
			}
		}
		if j, ok := psumToIndex[sum-target]; ok {
			if minSize > i-j {
				minSize = i - j
			}
		}
	}
	if minSize == math.MaxInt {
		return -1
	}
	return minSize + result
}
