package maximum_gap

/**
1. Because of the size of input, logN is a constant, log_2(10^5)=17. So technically, using quick sort works.
2. The input is all integers with at most 9 digits, so we can use Radix sort. That brings down the constant to 9.

The two approaches above sort the input. Is there a way that does not sort the input and can get us the answer in two scans
of the data?
Yes. If we can find the following observations.
1. If there an oracle that tells us the max-gap is larger than X, then, we can bucket nums into buckets of range: [min, min+X], [min+X, min+2X],  ....
Because of the max-gap is larger than X, the max-gap cannot come from two numbers from the same bucket. The max-gap must be from two
non-empty adjacent buckets. Assume the two buckets are b[i] and b[j]. The max-gap is min(b[j]) - max(b[i]).

2. So how do we find X? Notice that all we need from X is that just a lower-bound of the max-gap. Assume there are n numbers,
with min and max. One lower-bound is (max-min)/(n-1). We can prove it by contradiction. If the max-gap is smaller than `(max-min)/(n-1)`,
then the total sum of gaps is smaller than (n-1)*(max-min)/(n-1) = max-min, which is the sum of gaps.

So, there we have the X. (max-min)/(n-1). We'll break nums into at most (n-1) buckets. You may think it's weird bucketing n numbers
into (n-1) buckets. What information does the bucketing add? The bucketing let us avoid comparing numbers within the same bucket and
avoid comparing numbers across non-empty non-adjacent buckets.
*/

func maximumGap(nums []int) int {
	// Step 0: special/easy case handling
	n := len(nums)
	if n < 2 {
		return 0
	}
	if n == 2 {
		return nums[1] - nums[0]
	}

	// Step 1: Calculate the lower-bound.
	min, max := nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	if min == max {
		return 0
	}

	// TODO: Explain this "+1".
	x := (max-min)/(n-1) + 1

	// Step 2: Bucketing
	buckets := make([][]int, n+1)
	for _, v := range nums {
		idx := (v - min) / x
		if len(buckets[idx]) == 0 {
			buckets[idx] = []int{v, v}
		} else {
			if v < buckets[idx][0] {
				buckets[idx][0] = v
			} else if v > buckets[idx][1] {
				buckets[idx][1] = v
			}
		}
	}

	// Step 3: Find the max-gap from adjacent buckets
	maxGap := 0
	m := len(buckets)
	for i := 0; i < m-1; {
		j := i + 1
		for j = i + 1; j < m; {
			if len(buckets[j]) == 0 {
				j++
				continue
			}
			v := buckets[j][0] - buckets[i][1]
			if v > maxGap {
				maxGap = v
			}
			i = j
			break
		}
		if j == m {
			break
		}
	}
	return maxGap
}
