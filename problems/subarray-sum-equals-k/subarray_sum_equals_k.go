package subarray_sum_equals_k

/*
Sum of subarray a[i...j] = sum[j]-sum[i-1].
*/

func subarraySum(nums []int, k int) int {
	n := len(nums)
	// cnt[x]=y: there are y prefixes, a[0...i], whose sum is x. i < x.
	cnt := make(map[int]int)
	sum := 0
	cnt[0] = 1
	result := 0
	for i := 0; i < n; i++ {
		result += cnt[sum+nums[i]-k]
		sum += nums[i]
		cnt[sum]++
	}
	return result
}
