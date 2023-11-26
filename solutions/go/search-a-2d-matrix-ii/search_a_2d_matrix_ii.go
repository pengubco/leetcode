package searcha2dmatrixii

// https://leetcode.com/problems/search-a-2d-matrix-ii

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	r, c := n-1, 0
	for r >= 0 && c < m {
		if target == matrix[r][c] {
			return true
		}
		if target < matrix[r][c] {
			r--
			continue
		}
		c++
	}
	return false
}
