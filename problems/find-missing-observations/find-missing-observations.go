package find_missing_observations

/*
The problem is equivalent to: Find n integers, each from 1 to 6, so that the sum of the n integer is X.

1. First of all, it's easy to know whether there is an answer or not. There is an answer IFF `X >= n && X <=n*6`
2. How to find an answer? Visualize n bins, and X balls drop into bins. No bin should be empty. So we just fill
each bin with one ball at a time, left to right. If there is ball left, repeat ball dropping with the bin from the left.
*/
func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	sum := mean * (m + n)
	for _, v := range rolls {
		sum -= v
	}
	if sum < n || sum > n*6 {
		return []int{}
	}

	result := make([]int, n)
	v := sum / n
	for i := 0; i < n; i++ {
		result[i] = v
	}
	sum -= n * v
	for i := 0; i < n && sum > 0; i++ {
		capacity := 6 - result[i]
		if sum >= capacity {
			sum -= capacity
			result[i] += capacity
		} else {
			result[i] += sum
			break
		}
	}

	return result
}
