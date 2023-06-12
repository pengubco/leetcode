package subarray_sums_divisible_by_k

/*
https://en.wikipedia.org/wiki/Modular_arithmetic
1. If (a-b)%k == 0, then a%k == b%k.
2. The sum of a subarray nums[i...j] = s[j]-s[i-1]. Where s[i] is the sum of nums[0, ...., i].
3. How many valid subarrays end at nums[i]? The answer is the number of s[j] that satisfy (s[i]-s[j])%k == 0, j < i.
*/
func subarraysDivByK(nums []int, k int) int {
	n := len(nums)
	result := 0
	// Running sum
	sum := 0
	// cnt[x] := number of subarrays s[j] with s[j]%k == x
	cnt := make([]int, k)
	cnt[0] = 1
	for i := 0; i < n; i++ {
		sum += nums[i]
		mod := sum % k
		if mod < 0 {
			mod += k
		}
		result += cnt[mod]
		cnt[mod]++
	}
	return result
}
