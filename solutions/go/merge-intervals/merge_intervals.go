package mergeintervals

import "sort"

/*
line scanning from left to right.
*/
func merge(intervals [][]int) [][]int {
	var points []Point
	for _, interval := range intervals {
		a, b := interval[0], interval[1]
		points = append(points, Point{a, true}, Point{b, false})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].x != points[j].x {
			return points[i].x < points[j].x
		}
		return points[i].beginning
	})

	var ans [][]int
	cnt := 0
	start := -1

	n := len(points)
	for i := 0; i < n; i++ {
		if start == -1 {
			start = points[i].x
			cnt++
			continue
		}
		if points[i].beginning {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			ans = append(ans, []int{start, points[i].x})
			start = -1
			continue
		}
	}

	return ans
}

type Point struct {
	x         int
	beginning bool
}
