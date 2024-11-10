package singlenumberii

func singleNumber(nums []int) int {
	result := 0
	for i := 0; i <= 31; i++ {
		bitCount := 0
		for _, x := range nums {
			// handle Two's complement
			if x < 0 {
				x += (1 << 32)
			}
			if x&(1<<i) > 0 {
				bitCount++
			}
		}
		if bitCount%3 > 0 {
			result |= (1 << i)
		}
	}
	// handle Two's complement. the highest bit is sign.
	// say the signed number of 3 digits, [-4, 3].
	// 111 would be 1 * (-4) + 1 * 2 + 1 * 1 = -1.
	if result&(1<<31) > 0 {
		result -= (1 << 32)
	}
	return result
}
