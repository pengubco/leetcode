package uglynumberii

// https://leetcode.com/problems/ugly-number-ii

func nthUglyNumber(n int) int {
	f := make([]int, n)
	f[0] = 1
	i, j, k := 0, 0, 0
	for cnt := 1; cnt < n; cnt++ {
		nextMin := min(f[i]*2, f[j]*3, f[k]*5)
		f[cnt] = nextMin
		if nextMin == f[i]*2 {
			i++
		}
		if nextMin == f[j]*3 {
			j++
		}
		if nextMin == f[k]*5 {
			k++
		}
	}
	return f[n-1]
}
