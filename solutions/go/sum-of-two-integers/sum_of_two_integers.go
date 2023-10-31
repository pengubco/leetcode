package sumoftwointegers

/*
1. sum the non-carry and the carry.
2. 2's complement works the same for both positive numbers and negative numbers.
*/
func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
