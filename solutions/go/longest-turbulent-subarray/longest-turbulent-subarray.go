package longestturbulentsubarray

// https://leetcode.com/problems/longest-turbulent-subarray
func maxTurbulenceSize(a []int) int {
	n := len(a)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, 2)
	}
	f[0][0], f[0][1] = 1, 1
	ans := 1
	for i := 1; i < n; i++ {
		switch {
		case a[i] > a[i-1]:
			f[i][1] = max(1, f[i-1][0]+1)
			f[i][0] = 1
		case a[i] == a[i-1]:
			f[i][1], f[i][0] = 1, 1
		default:
			f[i][0] = max(1, f[i-1][1]+1)
			f[i][1] = 1
		}
		ans = max(ans, f[i][0])
		ans = max(ans, f[i][1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
