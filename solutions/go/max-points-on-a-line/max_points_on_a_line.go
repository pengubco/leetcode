package max_points_on_a_line

// https://leetcode.com/problems/max-points-on-a-line

import (
	"sort"
)

func maxPoints(points [][]int) int {
	n := len(points)
	if n == 1 {
		return 1
	}
	sort.Slice(points, func(i, j int) bool {
		x1, x2 := points[i][0], points[j][0]
		if x1 != x2 {
			return x1 < x2
		}
		return points[i][1] < points[j][1]
	})
	maxPoints := 0

	// find points on vertical lines
	vLines := make(map[int]int)
	for _, p := range points {
		v := vLines[p[0]] + 1
		vLines[p[0]] = v
		if v > maxPoints {
			maxPoints = v
		}
	}

	// non vertical lines
	lines := make(map[[2]float64]int)
	lineStartPoint := make(map[[2]float64]int)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1, p2 := points[i], points[j]
			if p1[0] == p2[0] {
				continue
			}
			a := float64(p2[1]-p1[1]) / float64(p2[0]-p1[0])
			b := float64(p1[1]) - a*float64(p1[0])
			k := [2]float64{a, b}
			// the line has been explored. ignore.
			if idx, ok := lineStartPoint[k]; ok && idx != i {
				continue
			}
			// this is a new line.
			v, ok := lines[k]
			if !ok {
				lineStartPoint[k] = i
				lines[k] = 2
				v = 2
			} else {
				v++
				lines[k] = v
			}
			if v > maxPoints {
				maxPoints = v
			}
		}
	}
	return maxPoints
}
