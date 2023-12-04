package minimum_number_of_arrows_to_burst_balloons

import "sort"

// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons

/*
1. If an optimal arrow shoot at x and x is not an end of any interval, then we can always move line left or right to the end
of a point maintaining the optimal solution.
2. For an interval A[x,y], whether we should x or y?
	2.1. If no other interval overlaps with A, it doesn't matter whether we choose x or y.
	2.2. If all intervals is within [x,y], e.g. [x-2, y-1]. Then we should not choose either x or y.
    2.3. If intervals overlap on the right side. e.g. [y-4, y+10], then we should choose y.
	2.4. If intervals overlap on the left side. e.g. [x-5, x+1], then we choose x.

How may we simplify 4 cases?
1. Sort by the right end, y, in ascending order. This remove case 2.4. And for 2.1 and 2.3, we choose y.
2. If there are ties, intervals with the same right end, then we sort based on the length of the interval. That is,
sort in descending order of the left side, this remove case 2.2.
*/

func findMinArrowShots(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] > intervals[j][0]
	})
	n := len(intervals)
	result := 0
	for i := 0; i < n; {
		result++
		line := intervals[i][1]
		j := i + 1
		for j < n {
			if intervals[j][0] <= line && intervals[j][1] >= line {
				j++
			} else {
				break
			}
		}
		i = j
	}
	return result
}
