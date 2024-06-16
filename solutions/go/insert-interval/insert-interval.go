package insert_interval

// https://leetcode.com/problems/insert-interval
func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	n := len(intervals)
	newIntervalAdded := false
	i := 0
	for i < n || !newIntervalAdded {
		// should we insert intervals[i] or the newInterval
		var intervalToInsert []int
		if i >= n {
			intervalToInsert = newInterval
			newIntervalAdded = true
		} else if newIntervalAdded {
			intervalToInsert = intervals[i]
			i++
		} else if intervals[i][0] < newInterval[0] {
			intervalToInsert = intervals[i]
			i++
		} else {
			intervalToInsert = newInterval
			newIntervalAdded = true
		}

		// insert the intervalToInsert
		if len(result) == 0 {
			result = append(result, intervalToInsert)
			continue
		}
		m := len(result)
		lastInterval := result[m-1]
		if lastInterval[1] < intervalToInsert[0] {
			result = append(result, intervalToInsert)
			continue
		}
		result[m-1] = []int{
			min(lastInterval[0], intervalToInsert[0]),
			max(lastInterval[1], intervalToInsert[1]),
		}
	}
	return result
}
