package flipstringtomonotoneincreasing

/*
f[i][j]: min flip so that a[0...i] ends with j. j=0,1

f[i][0] = f[i-1][0] + (a[i]==1)
f[i][1] = min( f[i-1][0] + a[i]==0; f[i-1][1] + a[i]==0)
*/
func minFlipsMonoIncr(s string) int {
	n := len(s)
	f := make([][2]int, n)
	if s[0] == '0' {
		f[0][0] = 0
		f[0][1] = 1
	} else {
		f[0][0] = 1
		f[0][1] = 0
	}
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			f[i][0] = f[i-1][0]
			f[i][1] = min(f[i-1][0], f[i-1][1]) + 1
		} else {
			f[i][0] = f[i-1][0] + 1
			f[i][1] = min(f[i-1][0], f[i-1][1])
		}
	}
	return min(f[n-1][0], f[n-1][1])
}

// only use this if it's before Go 1.21
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
