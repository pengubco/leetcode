package maximum_gap

// https://leetcode.com/problems/maximum-gap/

func maximumGap(nums []int) int {
	// Step 0: special/easy case handling
	n := len(nums)
	if n < 2 {
		return 0
	}
	if n == 2 {
		return nums[1] - nums[0]
	}

	// Step 1: Calculate the lower-bound.
	min, max := nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	if min == max {
		return 0
	}

	bucketSize := 0
	if (max-min)%(n-1) == 0 {
		bucketSize = (max - min) / (n - 1)
	} else {
		bucketSize = (max-min)/(n-1) + 1
	}

	// Step 2: Bucketing
	buckets := make([][]int, n+1)
	for _, v := range nums {
		idx := (v - min) / bucketSize
		if len(buckets[idx]) == 0 {
			buckets[idx] = []int{v, v}
		} else {
			if v < buckets[idx][0] {
				buckets[idx][0] = v
			} else if v > buckets[idx][1] {
				buckets[idx][1] = v
			}
		}
	}

	// Step 3: Find the max-gap from adjacent buckets
	maxGap := 0
	m := len(buckets)
	for i := 0; i < m-1; {
		j := i + 1
		for j < m {
			if len(buckets[j]) == 0 {
				j++
				continue
			}
			v := buckets[j][0] - buckets[i][1]
			if v > maxGap {
				maxGap = v
			}
			i = j
			break
		}
		if j == m {
			break
		}
	}
	return maxGap
}
