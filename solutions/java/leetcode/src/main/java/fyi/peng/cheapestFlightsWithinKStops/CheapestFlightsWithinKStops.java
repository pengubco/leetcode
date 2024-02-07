package fyi.peng.cheapestFlightsWithinKStops;

import java.util.Arrays;

// https://leetcode.com/problems/cheapest-flights-within-k-stops/ 
public class CheapestFlightsWithinKStops {
	public int findCheapestPrice(int n, int[][] flights, int src, int dst, int k) {
		var prevDist = new int[n];
		Arrays.fill(prevDist, Integer.MAX_VALUE);
		prevDist[src] = 0;
		var curDist = new int[n];
		for (int i=0; i<=k; i++) {
			System.arraycopy(prevDist, 0, curDist, 0, n);
			for (var e : flights) {
				if (prevDist[e[0]] == Integer.MAX_VALUE) {
					continue;
				}
				curDist[e[1]] = Math.min(curDist[e[1]], prevDist[e[0]] + e[2]);
			}
			var tmp = prevDist;
			prevDist = curDist;
			curDist = tmp;
		}

		return prevDist[dst] != Integer.MAX_VALUE ? prevDist[dst] : -1;
	}
}
