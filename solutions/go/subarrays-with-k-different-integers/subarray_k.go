package subarray_k

// https://leetcode.com/problems/subarrays-with-k-different-integers/

/*
f[j]: number of subarrays that ends at a[j] with at most k different integers.
f[j] = j-i+1. where a[i...j] is the longest subarray that ends with a[j] has at most k different integers.

define g[j]: number of subarrays that ends at a[j] with at most k-1 different integers.

Then ans = sum{f[j]-g[j]} for 0 <= j <= n-1
*/
func subarraysWithKDistinct(a []int, k int) int {
	n := len(a)
	f := atmostK(a, k)

	ans := 0

	if k == 1 {
		for i := 0; i < n; i++ {
			ans += f[i]
		}
		return ans
	}

	g := atmostK(a, k-1)
	for i := 0; i < n; i++ {
		ans += f[i] - g[i]
	}
	return ans
}

func atmostK(a []int, k int) []int {
	n := len(a)
	cnt := make(map[int]int)
	f := make([]int, n)
	i := 0
	cnt[a[0]] = 1
	f[0] = 1
	for j := 1; j < n; {
		if _, ok := cnt[a[j]]; ok {
			cnt[a[j]]++
			f[j] += j - i + 1
			j++
			continue
		}
		if len(cnt) < k {
			cnt[a[j]] = 1
			f[j] += j - i + 1
			j++
			continue
		}
		cnt[a[i]]--
		if cnt[a[i]] == 0 {
			delete(cnt, a[i])
		}
		i++
	}
	return f
}
