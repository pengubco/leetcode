package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	r, c := len(matrix), len(matrix[0])
	tranlate := func(idx int) int {
		return matrix[idx/c][idx%c]
	}
	return binarySearch(0, r*c-1, tranlate, target)
}

func binarySearch(l, r int, f func(int) int, target int) bool {
	for {
		if l > r {
			return false
		}
		if r-l < 5 {
			for i := l; i <= r; i++ {
				if f(i) == target {
					return true
				}
			}
			return false
		}

		mid := (r + l) / 2
		if f(mid) < target {
			l = mid + 1
		} else if f(mid) == target {
			return true
		} else {
			r = mid - 1
		}
	}
}
