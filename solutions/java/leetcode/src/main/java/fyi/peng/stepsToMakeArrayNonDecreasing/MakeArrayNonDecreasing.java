package fyi.peng.stepsToMakeArrayNonDecreasing;

import java.util.ArrayDeque;
import java.util.Deque;

public class MakeArrayNonDecreasing {
	public int totalSteps(int[] nums) {
		int maxStep = 0;
		int[] f = new int[nums.length];
		Deque<Integer> stack = new ArrayDeque<Integer>();
		for (int i=0; i<nums.length; i++) {
			int step = 0;
			while (!stack.isEmpty() && nums[stack.peekLast()] <= nums[i] ) {
				step = Math.max(step, f[stack.pollLast()] + 1);
			}
			if (stack.isEmpty()) {
				f[i] = 0;
			} else {
				f[i] = Math.max(1, step);
				maxStep = Math.max(maxStep, f[i]);
			}
			stack.offerLast(i);
		}
		return maxStep;	
	}	
}
