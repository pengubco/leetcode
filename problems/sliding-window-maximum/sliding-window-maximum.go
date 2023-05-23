package sliding_window_maximum

// special case of RMQ.
func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 || len(nums) == 1 {
		return nums
	}
	return nil
}
