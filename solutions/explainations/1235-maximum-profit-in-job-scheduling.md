
// https://leetcode.com/problems/maximum-profit-in-job-scheduling

Greedy algorithm. 
1. Consider jobs in the non-decreasing order of endTime.
2. Record the max profit we can get at or before endTime. For example, f[k]=v, means, 
   at time k, the max profit we can get is v. 
3. Consider each job j, find the max profit at j.startTime. See whether 
   we can get bigger profit picking up j. 
4. The max profit mapping f is sorted by k. And, by definition of f[k]=v, the 
   value is also sorted. 
5. Implement it in Java since Go doesn't have good support for binary search. 