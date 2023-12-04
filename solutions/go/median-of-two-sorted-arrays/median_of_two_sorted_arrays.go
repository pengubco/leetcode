package medianoftwosortedarrays

import "math"

// https://leetcode.com/problems/median-of-two-sorted-arrays
/*
General: find the k-th number of 2 sorted arrays.

Assume we have the first K number. Then there must be x numbers from nums1, and y
numbers from nums2. x>=0 and y>=0.

So the solution is to binary search x.
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if n1 == 0 {
		return medianOfArray(nums2)
	}
	if n2 == 0 {
		return medianOfArray(nums1)
	}

	if (n1+n2)&1 > 0 {
		return float64(findKthElements(nums1, nums2, (n1+n2)/2+1))
	}
	k := (n1 + n2) / 2
	return float64(findKthElements(nums1, nums2, k)+
		findKthElements(nums1, nums2, k+1)) / 2
}

func medianOfArray(nums []int) float64 {
	n := len(nums)
	if n&1 > 0 {
		return float64(nums[n/2])
	}
	return float64((nums[n/2-1] + nums[n/2])) / 2
}

// find the K-th element from two sorted arrays.
func findKthElements(nums1, nums2 []int, k int) int {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		n1, n2 = n2, n1
		nums1, nums2 = nums2, nums1
	}

	// use 0 number from nums1
	if k <= n2 && nums1[0] >= nums2[k-1] {
		return nums2[k-1]
	}
	// use all k numbers from nums1
	if k <= n1 && nums2[0] >= nums1[k-1] {
		return nums1[k-1]
	}

	// binary search how many elements we pick from nums1.
	low, high := 1, min(k-1, n1)
	for low < high {
		mid := (low + high) / 2
		cntInNums2 := k - mid
		if cntInNums2 < 0 {
			high = mid - 1
			continue
		}
		if cntInNums2 > n2 {
			low = mid + 1
			continue
		}
		// a b
		// c d
		// valid rule: a < d && c < b
		a, c := nums1[mid-1], math.MinInt
		if cntInNums2-1 >= 0 {
			c = nums2[cntInNums2-1]
		}
		b, d := math.MaxInt, math.MaxInt
		if mid < n1 {
			b = nums1[mid]
		}
		if cntInNums2 < n2 {
			d = nums2[cntInNums2]
		}
		if a > d {
			high = mid - 1
			continue
		}
		if c > b {
			low = mid + 1
			continue
		}
		low = mid
		break
	}
	ans := nums1[low-1]
	if k-low-1 >= 0 {
		ans = max(nums1[low-1], nums2[k-low-1])
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
