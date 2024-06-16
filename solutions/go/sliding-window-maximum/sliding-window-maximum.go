package sliding_window_maximum

// https://leetcode.com/problems/sliding-window-maximum/

// special case of RMQ.
// https://cp-algorithms.com/sequences/rmq.html
// https://cp-algorithms.com/data_structures/sparse-table.html

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 || len(nums) == 1 {
		return nums
	}
	n := len(nums)
	log2 := calcLog2(n)
	st := calcSparseTable(nums, log2)
	results := make([]int, n-k+1)
	for l := 0; l+k-1 < n; l++ {
		results[l] = rmq(st, l, l+k-1, log2)
	}
	return results
}

func rmq(st [][]int, l, r int, log2 []int) int {
	k := log2[r-l+1]
	return max(st[k][l], st[k][r-(1<<k)+1])
}

func calcLog2(n int) []int {
	// precompute log_2
	log2 := make([]int, n+1)
	log2[1] = 0
	for i := 2; i <= n; i++ {
		log2[i] = log2[i/2] + 1
	}
	return log2
}

func calcSparseTable(a []int, log2 []int) [][]int {
	n := len(a)

	// sparse table, dynamic programming
	// st[i][j]: the minimum in range a[j, j + 2^i - 1]
	k := log2[n]
	st := make([][]int, k+1)
	for i := 0; i <= k; i++ {
		st[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		st[0][j] = a[j]
	}

	for i := 1; i <= k; i++ {
		for j := 0; j+(1<<i)-1 < n; j++ {
			st[i][j] = max(st[i-1][j], st[i-1][j+(1<<(i-1))])
		}
	}

	return st
}
