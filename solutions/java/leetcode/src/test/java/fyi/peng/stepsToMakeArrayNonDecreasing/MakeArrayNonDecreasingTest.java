package fyi.peng.stepsToMakeArrayNonDecreasing;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class MakeArrayNonDecreasingTest {
	
	@Test
	public void testTotalSteps () {
		MakeArrayNonDecreasing s = new MakeArrayNonDecreasing();
		assertEquals(3, s.totalSteps(new int[]{5,3,4,4,7,3,6,11,8,5,11}));
	}
}
