package continuoussubarraysum

// https://leetcode.com/problems/continuous-subarray-sum

func checkSubarraySum(nums []int, k int) bool {
	sumToIndex := make(map[int]int)
	sumToIndex[0] = -1
	sum := 0
	for i, x := range nums {
		sum += x
		if j, ok := sumToIndex[sum%k]; ok {
			if i-j >= 2 {
				return true
			}
		} else {
			sumToIndex[sum%k] = i
		}
	}
	return false
}
