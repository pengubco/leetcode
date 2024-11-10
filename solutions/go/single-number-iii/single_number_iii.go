package singlenumberiii

func singleNumber(nums []int) []int {
	// find x xor y
	xorResult := 0
	for _, n := range nums {
		xorResult ^= n
	}
	// find one bit where x and y differ, i.e., the bit where x^y is 1.
	rightMostBit := xorResult & -xorResult
	x, y := 0, 0
	for _, n := range nums {
		if n&rightMostBit > 0 {
			x ^= n
		} else {
			y ^= n
		}
	}
	return []int{x, y}
}
