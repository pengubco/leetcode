package pattern

// https://leetcode.com/problems/132-pattern/
// http://localhost:1313/post/monotonicity-subinterval-132/
func find132pattern(a []int) bool {
	n := len(a)
	// f[i] = min{A[0:i]}
	f := make([]int, n)
	f[0] = a[0]
	for i := 1; i < n; i++ {
		f[i] = min(f[i-1], a[i])
	}
	// candidates for silver.
	var s []int
	// enumerate candidates for Gold
	for i := n - 1; i >= 1; i-- {
		// a[i] is not fit for Gold nor Silver
		if a[i] <= f[i] {
			continue
		}
		// check each silver candidate
		for len(s) > 0 && s[len(s)-1] <= f[i] {
			s = s[:len(s)-1]
		}
		// found it!
		if len(s) > 0 && s[len(s)-1] < a[i] {
			return true
		}
		// a[i] cannot be Gold, but it could be Silver
		s = append(s, a[i])
	}
	return false
}
