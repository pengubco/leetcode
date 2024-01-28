package fyi.peng.maximumProfitInJobScheduling;

import java.util.ArrayList;
import java.util.List;
import java.util.TreeMap;

// https://leetcode.com/problems/maximum-profit-in-job-scheduling/description/ 

public class Solution {
	public int jobScheduling(int[] startTime, int[] endTime, int[] profit) {
		int n = startTime.length;
		List<Job> jobs = new ArrayList<>();
		for (int i=0; i<n; i++){
			jobs.add(new Job(startTime[i], endTime[i], profit[i]));
		}
		jobs.sort((a,b) -> {
			return a.endTime - b.endTime;
		});
		
		int result = 0;
		// profits[k]=v, means, the max profit working on jobs finishes before or at k, is v.
		TreeMap<Integer, Integer> maxProfits = new TreeMap<>();
		for (Job job : jobs) {
			var lastMaxProfit = maxProfits.floorEntry(job.startTime);
			int maxProfit = 0;
			if (lastMaxProfit == null) {
				maxProfit = job.profit;
			} else {
				maxProfit = lastMaxProfit.getValue() + job.profit;
			}
			if (maxProfit > result) {
				result = maxProfit;
				maxProfits.put(job.endTime, maxProfit);
			}
		}

		return result;
	};

	class Job {
		int startTime;
		int endTime;
		int profit;

		public Job(int startTime, int endTime, int profit) {
			this.startTime = startTime;
			this.endTime = endTime;
			this.profit = profit;
		}
	}

};
