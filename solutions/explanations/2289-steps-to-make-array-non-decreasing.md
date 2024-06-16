
// https://leetcode.com/problems/steps-to-make-array-non-decreasing

Naive approach.
1. use a double link list to store nodes. in each round, find all nodes that are smaller 
its previous node. then remove them.
2. however, the approach is O(n^2) in the worst case. 
  where the first number is the maximum, and the rest of numbers are increasing. 
	In this case, each round, we can only remove one node.
	For example, [5, 4, 3, 2, 1]

Monotonic Stack.
https://peng.fyi/interview/steps-to-make-array-non-decreasing/ 

We need to look for key observations. First, it is easy to know which integers needs to be removed. All integers that is not in the final non-decreasing array needs to go. Moreover, the remaining integers separate the original array into subarrays and each each subarray can go through the removal steps independently. For example, [5, 4, 3, 2, 11, 10, 8, 9 11] has two subarrays to remove, [4, 3, 2] and [10, 8, 9]. This observation means that we can start refresh whenever we see an integer than cannot be removed.

Let f[i] be the order of step when A[i] is removed. If there is a longest decreasing subarray ends at A[j], say, A[i] > A[i+1] > A[i+2] > ... > A[j] and A[i] <= A[j+1], then f[i+1]=f[i+2]=...=f[j] and f[j+1]=f[i+1]+1. This observation of longest deceasing subarray can be applied recursively, after removal of A[i+1:j]. In order to handle the "removal" recursively, we use a stack. In order to calculate f[i], we store the index of the array in the stack.

