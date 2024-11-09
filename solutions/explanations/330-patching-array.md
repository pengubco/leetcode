
https://leetcode.com/problems/patching-array/description/

1. If [1, m] can be formed, by adding an number x, [1, m+x] can be formed. 
2. When m+1 is not available in nums, either because it has been used or not exist, 
   then we have to add an x, such that m+x >= m+1. According to 1, the optimal choice
	 of x is m+1 itself. This will effectively double the range of numbers can be formed.  