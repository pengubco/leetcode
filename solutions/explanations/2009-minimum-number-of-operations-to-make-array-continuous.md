
1. Order of nums does not impact the result. sort it.
2. Duplicated numbers do not help. 10, 10 is the same as 10, 100000. So we can deduplicate numbers.
3. The final continuous array must have at least one number from nums. so the result is at most n-1.
4. Following 3, how many numbers from nums we can reuse?
   Assume nums have been sorted and deduped. We can enumerate the range of numbers we reuse. 
	 `nums[i]` and `nums[j]`. If `nums[j]-nums[i]<=n-1`, then the number of operations needed is 
	 `n-(nums[j] - nums[i]+1)`. The time complexity is O(n^2).
5. Can we do better? Note that, we don't need to enumerate every pair (i,j). For exmaple, 
   if we find (i,j) is a valid pair, there is no need to enumerate (i1,j1) where i1 > i and j1<j. 
	 This is because `nums[j1]-nums[i1] < nums[j]- nums[i] <= n-1`. so choosing `(i1,j-1)` is always 
	 valid. And the operations needed for (i1,j1) is `n-(nums[j1] - nums[i1]+1) > n-(nums[j] - nums[i]+1)`. 
	 Therefore, for each i, we find the largest j that can forms a valid pair. Then we try pairs (x, y)
	 where x=i+1 and y >= j. 
	 It is a sliding window.