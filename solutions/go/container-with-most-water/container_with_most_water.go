package container_with_most_water

// https://leetcode.com/problems/container-with-most-water

/*
Approach 1: Binary search. O(N*logN)
Observations: Assume the shorter bar is a[i], then the other bar a[j] should satisfy.
1). a[j] >= a[i].
2). |j-i| is max.
Assume j<i, i.e., the higher bar is on the left side, then j is the leftist bar with a[j]>=a[i].
So if j<k<i and a[k]<=a[j], there is no need to consider a[k]. That is, we can "remove" a[k] from the input.
What does this "remove" mean? When considering a[i], we only need to keep an non-decreasing bars on its left.
This "non-decreasing" makes finding of the "leftist bar with a[j]>=a[i]" binary searchable.

We can apply the same algorithm for cases where j > i where we consider rightest instead of leftist.

Approach 2: Two pointers. O(N)
The area is |j-i| * min{a[i], a[j]}. When the |j-i| is fixed, we need to find a pair with the maximum min.
What is the largest |j-i|? i=0, j=n-1.
What is the 2nd largest |j-1|? either (1, n-1) or (0, n-2).
Should we move from (0,n-1) to (1,n-1) or (0,n-2)? We can use the same idea from Approach 1: The further the better, the higher the better.
We should keep the higher bar and replace the shorter bar.
If a[0] < a[n-1], then choose (1, n-1). Otherwise, choose (0,n-2).

Can we apply the same transition to reduce the width (|j-i|) to (n-2)?
Intuitively, Yes. But I think this requires a rigorous proof of the greedy algorithm.
*/

func maxArea(a []int) int {
	result := 0
	i, j := 0, len(a)-1
	for i < j {
		if a[i] < a[j] {
			area := (j - i) * a[i]
			if area > result {
				result = area
			}
			i++
			continue
		}
		area := (j - i) * a[j]
		if area > result {
			result = area
		}
		j--
	}
	return result
}
