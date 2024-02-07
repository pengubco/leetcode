package fyi.peng.cheapestFlightsWithinKStops;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class CheapestFlightsWithinKStopsTest {
	
	@Test
	public void testKStops() {
		var sol = new CheapestFlightsWithinKStops();
		var ans = sol.findCheapestPrice(4, new int[][]{
			{0,1,100},{1,2,100},{2,0,100},{1,3,600},{2,3,200}	
		}, 0, 3, 1);
		assertEquals(700, ans);


		ans = sol.findCheapestPrice(3, new int[][]{
			{0,1,100},{1,2,100},{0,2,500}	
		}, 0, 2, 1);
		assertEquals(200, ans);

		ans = sol.findCheapestPrice(3, new int[][]{
			{0,1,100},{1,2,100},{0,2,500}	
		}, 0, 2, 0);
		assertEquals(500, ans);

		ans = sol.findCheapestPrice(3, new int[][]{
			{0,1,2},{1,2,1},{2,0,10}	
		}, 1, 2, 1);
		assertEquals(1, ans);
	}
}
