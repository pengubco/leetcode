package superuglynumber

// for each prime number p, keep track of what is the smallest number generated so far that has not been multiplied by p.
// If the number is x, then x*p is the next smallest number if the number is some previous number multiply by p.
// The next smallest number must be a previous number multiplied by some prime number. So it is
// min { x1*p1, x2*p2, ..., }
// Be careful of advancing indices for each prime. 2*3 = 3*2. So we should advance the indices for both 2 and 3 when the
// next smallest numbe is 6.
func nthSuperUglyNumber(n int, primes []int) int {
	m := len(primes)
	// generated numbers
	numbers := make([]int, n)
	numbers[0] = 1
	indices := make([]int, m)
	for i := 1; i < n; i++ {
		next := numbers[indices[0]] * primes[0]
		for j := 1; j < m; j++ {
			temp := numbers[indices[j]] * primes[j]
			if temp < next {
				next = temp
			}
		}
		for j := 0; j < m; j++ {
			if numbers[indices[j]]*primes[j] == next {
				indices[j]++
			}
		}
		numbers[i] = next
	}
	return numbers[n-1]
}
