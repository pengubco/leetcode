package three_sum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var results [][]int
	n := len(nums)
	sort.Ints(nums)

	addTriple := func(t []int) {
		if len(results) == 0 {
			results = append(results, t)
			return
		}
		lastTriple := results[len(results)-1]
		for i := 0; i < 3; i++ {
			if t[i] != lastTriple[i] {
				results = append(results, t)
				break
			}
		}
	}

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, n-1; j < k && k < n; {
			sum := nums[i] + nums[j] + nums[k]
			switch {
			case sum == 0:
				addTriple([]int{nums[i], nums[j], nums[k]})
				k--
			case sum > 0:
				k--
			default:
				j++
			}
		}
	}
	return results
}
