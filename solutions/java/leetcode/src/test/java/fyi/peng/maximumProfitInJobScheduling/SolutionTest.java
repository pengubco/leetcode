package fyi.peng.maximumProfitInJobScheduling;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class SolutionTest {
	
	@Test
	public void testMaxProfit() {
		Solution s = new Solution();
		assertEquals(120, s.jobScheduling(new int[]{1,2,3,3}, new int[]{3,4,5,6}, new int[]{50,10,40,70}));

		assertEquals(150, s.jobScheduling(new int[]{1,2,3,4,6}, new int[]{3,5,10,6,9}, new int[]{20,20,100,70,60}));

		assertEquals(6, s.jobScheduling(new int[]{1,1,1}, new int[]{2, 3,4}, new int[]{5, 6, 4}));
	}
}
