package insert_interval

/*
Because intervals are not overlap, we only need to look at the last interval with the interval to insert.
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	n := len(intervals)
	if n == 0 {
		result = append(result, newInterval)
		return result
	}

	addedNewInterval := false
	for i := 0; i < n || !addedNewInterval; {
		// pick which interval to add
		var intervalToAdd []int
		if i >= n {
			intervalToAdd = newInterval
			addedNewInterval = true
		} else if addedNewInterval {
			intervalToAdd = intervals[i]
			i++
		} else {
			if newInterval[0] < intervals[i][0] {
				intervalToAdd = newInterval
				addedNewInterval = true
			} else {
				intervalToAdd = intervals[i]
				i++
			}
		}

		result = append(result, intervalToAdd)
		if len(result) < 2 {
			continue
		}

		// combine last two intervals
		m := len(result)
		a, b := result[m-2], result[m-1]
		if a[0] > b[0] {
			result[m-2], result[m-1] = result[m-1], result[m-2]
			a, b = b, a
		}
		// case 1: non overlap.
		if b[0] > a[1] {
			continue
		}
		// case 2: overlap
		newEnd := a[1]
		if b[1] > newEnd {
			newEnd = b[1]
		}
		result = result[:m-2]
		result = append(result, []int{a[0], newEnd})
	}
	return result
}
