package three_sum

import (
	"sort"
)

/*
1. Does sorting make the problem easier? Yes.
2. Does assuming no duplicate make the problem easier? Yes.
3. In a sorted, no duplicate array, in worst case, how many triples can have sum(i,j,k) = 0?
I think it's O(n).
4. We'll fix the nums[k] and find pair(i,j) such that nums[i]+nums[j] = -nums[k]. In the worst cae, this is O(n^2).
*/
func threeSum(nums []int) [][]int {
	var results [][]int
	sort.Ints(nums)
	n := len(nums)
	for k := n - 1; k >= 2; {
		if k < n-1 && nums[k] == nums[k+1] {
			k--
			continue
		}
		for i, j := 0, k-1; i < j && j < k; {
			tmp := nums[i] + nums[j] + nums[k]
			if tmp > 0 {
				j--
				continue
			}
			if tmp == 0 {
				m := len(results)
				triple := []int{nums[i], nums[j], nums[k]}
				if m == 0 || (m > 0 && !sliceEqual(results[m-1], triple)) {
					results = append(results, triple)
				}
				j--
				continue
			}
			i++
		}
		k--
	}
	return results
}

func sliceEqual(a, b []int) bool {
	return a[0] == b[0] && a[1] == b[1] && a[2] == b[2]
}
