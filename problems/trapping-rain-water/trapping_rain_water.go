package trapping_rain_water

/*
For the i-th location a[i], the water it can trap is max(0, min(l[i-1], r[i+1]) - a[i]).
l[i-1]: maximum value from a[0, ..., i-1].
r[i+1]: maximum value from a[i+1, ...., n-1].
*/
func trap(a []int) int {
	n := len(a)
	l, r := make([]int, n), make([]int, n)
	l[0], r[n-1] = a[0], a[n-1]
	for i := 1; i <= n-3; i++ {
		if a[i] > l[i-1] {
			l[i] = a[i]
		} else {
			l[i] = l[i-1]
		}
	}
	for i := n - 2; i >= 2; i-- {
		if a[i] > r[i+1] {
			r[i] = a[i]
		} else {
			r[i] = r[i+1]
		}
	}
	area := 0
	for i := 1; i <= n-2; i++ {
		lower := l[i-1]
		if r[i+1] < lower {
			lower = r[i+1]
		}
		if lower > a[i] {
			area += lower - a[i]
		}
	}
	return area
}
