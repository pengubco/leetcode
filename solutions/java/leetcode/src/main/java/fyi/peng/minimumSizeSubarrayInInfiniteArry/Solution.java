package fyi.peng.minimumSizeSubarrayInInfiniteArry;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

// https://leetcode.com/problems/minimum-size-subarray-in-infinite-array
public class Solution {
	public int minSizeSubarray(int[] nums, int target) {
		int totalSum = Arrays.stream(nums).sum();
		int copies = target/totalSum;
		target = target%totalSum;

		if (target == 0) {
			return copies * nums.length;
		}
		
		Map<Integer, Integer> sumToIndex = new HashMap<>();
		sumToIndex.put(0,-1);
		int sum = 0;
		int minSize = Integer.MAX_VALUE;
		for (int i=0; i< 2*nums.length; i++) {
			sum += nums[i%(nums.length)];	
			sumToIndex.put(sum, i);
			Integer idx = sumToIndex.get(sum-target);
			if (idx == null) {
				continue;
			}
			if (i-idx < minSize) {
				minSize = i-idx;
			}
		}
		if (minSize == Integer.MAX_VALUE) {
			return -1;
		}
		return minSize + copies*nums.length;
	}
}
