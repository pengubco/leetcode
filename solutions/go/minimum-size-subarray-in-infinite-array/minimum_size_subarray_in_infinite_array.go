package minimumsizesubarrayininfinitearray

import "math"

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
			minSize = min(minSize, i-j)
		}
	}
	if minSize == math.MaxInt {
		return -1
	}
	return minSize + result
}
