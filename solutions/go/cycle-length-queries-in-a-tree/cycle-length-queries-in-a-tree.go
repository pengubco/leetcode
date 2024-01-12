package cyclelengthqueriesinatree

// https://leetcode.com/problems/cycle-length-queries-in-a-tree/submissions/1143933691/
// the idea is simple. but there are many edge cases to be careful and off-by-one error!
// Let's say the two nodes are A and B. edge case: A is B's ancestor.

func cycleLengthQueries(n int, queries [][]int) []int {
	m := len(queries)
	result := make([]int, m)
	for i, q := range queries {
		a, b := q[0], q[1]
		result[i] = count(a, b) + 1
	}
	return result
}

// count the length of paths from a to b
func count(a, b int) int {
	if a == b {
		return 0
	}
	if a > b {
		a, b = b, a
	}
	pathA, pathB := []int{a}, []int{b}
	for a > 1 {
		a >>= 1
		pathA = append(pathA, a)
	}
	for b > 1 {
		b >>= 1
		pathB = append(pathB, b)
	}
	sameSuffixCnt := 0
	for i, j := len(pathA)-1, len(pathB)-1; i >= 0 && j >= 0 && pathA[i] == pathB[j]; {
		sameSuffixCnt++
		i--
		j--
	}
	return len(pathA) + len(pathB) - 2*sameSuffixCnt
}
