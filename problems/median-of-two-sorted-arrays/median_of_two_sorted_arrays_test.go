package medianoftwosortedarrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	cases := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		// {[]int{}, []int{2}, 2},
		{[]int{3}, []int{1, 2, 4}, 2.5},
		// {[]int{1, 3}, []int{2}, 2},
		// {[]int{1, 2, 3}, []int{4, 5, 6, 7}, 4},
		// {[]int{1, 2, 3}, []int{4, 5, 6}, 3.5},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, findMedianSortedArrays(tc.nums1, tc.nums2))
	}
}

func TestFindKth(t *testing.T) {
	cases := []struct {
		nums1    []int
		nums2    []int
		k        int
		expected int
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6, 7}, 7, 7},
		{[]int{1, 2, 3}, []int{4, 5, 6, 7}, 1, 1},
		{[]int{1, 2, 3}, []int{4, 5, 6, 7}, 2, 2},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 1, 1},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 2, 1},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 3, 2},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 4, 2},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 5, 3},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 6, 3},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expected, findKthElements(tc.nums1, tc.nums2, tc.k))
	}
}
