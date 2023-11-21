package minimumnumberofoperationstomakearraycontinuous

import (
	"slices"
)

/*
key observations
 1. order of nums does not impact the result. sort it.
 2. duplicated numbers do not help. 10, 10 is the same as 10, 100000.
 3. The final continuous array must have at least one number from nums. so the result is at most n-1.
 4. the question following 3. is: how many numbers from nums we can reuse?
    Say, the x, y are the minimum and maximum numbers from nums that we reuse.
    Then y-x+1 <= n and number-of-operations = n - count(x,y).
    count(x,y) is the numbers that is >=x and <=y.
    It is easy to calculate `count(x,y)` if have nums without duplication.
    count(x,y) = index(y) - index(x) + 1
    Note that x could be equal to y.
 5. we end up with using sliding window. Fix x, find the largest y. Then increase x, find the largest y.
    repeat.
*/
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
			cnt := j - i + 1
			if n-cnt < result {
				result = n - cnt
			}
			j++
		}
	}
	return result
}
