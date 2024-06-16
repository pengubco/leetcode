package trapping_rain_water

// https://leetcode.com/problems/trapping-rain-water

/*
For the i-th location a[i], the water it can trap is max(0, min(l[i-1], r[i+1]) - a[i]), which is
max(0, min(l[i], r[i]) - a[i]).

l[i]: maximum value from a[0, ..., i].
r[i]: maximum value from a[i, ...., n-1].
*/
func trap(a []int) int {
	n := len(a)
	l, r := make([]int, n), make([]int, n)
	maxVal := a[0]
	l[0] = a[0]
	for i := 1; i < n; i++ {
		maxVal = max(maxVal, a[i])
		l[i] = maxVal
	}
	maxVal = a[n-1]
	r[n-1] = a[n-1]
	for i := n - 2; i >= 0; i-- {
		maxVal = max(maxVal, a[i])
		r[i] = maxVal
	}

	area := 0
	for i := 1; i <= n-2; i++ {
		area += max(0, min(l[i], r[i])-a[i])
	}

	return area
}
