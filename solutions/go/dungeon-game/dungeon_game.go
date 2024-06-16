package dungeon_game

import "math"

// https://leetcode.com/problems/dungeon-game

/*
It may seem like a 'path of maximum sum' at first. However, the problem requires a path must have positive value at any time.

We can binary search the answer. If given an initial value of X, can we find a path that is positive at any time.
Let f(i,j) be the maximum path from top-left to the cell at (i,j).
f(i,j) = max{ f(i-1,j), f(i,j-1) } + a[i,j]. If f(i-1,j), f(i,j-1) and f(i,j) > 0

How to determine lower-bound and upper-bound for the binary search?
lower-bound: -a[0][0]+1, if a[0][0]<0.
upper-bound: pick a path, note the minimum negative value along the path. e.g. path: all the way left -> all the way down;
or all the way down -> all the way left.
*/
func calculateMinimumHP(a [][]int) int {
	v := smallestValueOnPath(a)
	if v > 0 {
		return 1
	}

	lowerBound := 1
	if a[0][0] < 0 {
		lowerBound = -a[0][0] + 1
	}
	upperBound := -v + 1

	n := len(a)
	m := len(a[0])
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, m)
	}

	for lowerBound+1 < upperBound {
		mid := (lowerBound + upperBound) / 2
		if okPath(mid, a, f) {
			upperBound = mid
		} else {
			lowerBound = mid + 1
		}
	}
	if okPath(lowerBound, a, f) {
		return lowerBound
	}
	return upperBound
}

func smallestValueOnPath(a [][]int) int {
	minPathVal, val := a[0][0], a[0][0]
	n, m := len(a), len(a[0])
	for j := 1; j < m; j++ {
		val += a[0][j]
		if val < minPathVal {
			minPathVal = val
		}
	}
	for i := 1; i < n; i++ {
		val += a[i][m-1]
		if val < minPathVal {
			minPathVal = val
		}
	}
	return minPathVal
}

func okPath(initialPoint int, a, f [][]int) bool {
	// initiate f[][]
	n := len(a)
	m := len(a[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			f[i][j] = math.MinInt32
		}
	}
	f[0][0] = a[0][0] + initialPoint
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > 0 && f[i-1][j] > 0 {
				f[i][j] = max(f[i][j], f[i-1][j]+a[i][j])
			}
			if j > 0 && f[i][j-1] > 0 {
				f[i][j] = max(f[i][j], f[i][j-1]+a[i][j])
			}
		}
	}
	return f[n-1][m-1] > 0
}
