package intervallistintersections

// https://leetcode.com/problems/interval-list-intersections/

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var result [][]int
	n, m := len(firstList), len(secondList)
	for i, j := 0, 0; i < n && j < m; {
		bi, ei := firstList[i][0], firstList[i][1]
		bj, ej := secondList[j][0], secondList[j][1]
		if ei < bj {
			i++
			continue
		}
		if ej < bi {
			j++
			continue
		}
		result = append(result, []int{max(bi, bj), min(ei, ej)})
		if ei < ej {
			i++
		} else if ei > ej {
			j++
		} else {
			i++
			j++
		}
	}
	return result
}
